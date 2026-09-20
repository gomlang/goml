from concurrent.futures import ThreadPoolExecutor
import importlib.util
from pathlib import Path
import tempfile
import threading
import tomllib
import unittest
from unittest.mock import patch


SPEC = importlib.util.spec_from_file_location("ecosystem_verify", Path(__file__).resolve().parents[1] / "verify.py")
VERIFY = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(VERIFY)


class RegistrySnapshots(unittest.TestCase):
    def setUp(self):
        self.temporary = tempfile.TemporaryDirectory()
        self.addCleanup(self.temporary.cleanup)
        self.root = Path(self.temporary.name)
        self.addCleanup(patch.stopall)
        patch.object(VERIFY, "ROOT", self.root).start()
        patch.object(VERIFY, "MODULES", ("one", "two")).start()
        for name in VERIFY.MODULES:
            module = self.root / name
            module.mkdir()
            (module / "goml.toml").write_text(f'[module]\npath = "ecosystem::{name}"\n')
            (module / "main.gom").write_text(f"package {name};\n")
            (module / "_artifact").mkdir()
            (module / "_artifact" / "ignored").write_text("build output")

    def verify_complete(self, home):
        registry = home / "cache" / "registry"
        index = tomllib.loads((registry / "index.toml").read_text())
        self.assertEqual(set(index["modules"]), {"ecosystem::one", "ecosystem::two"})
        for name in VERIFY.MODULES:
            package = registry / "ecosystem" / name / "0.1.0"
            self.assertEqual((package / "main.gom").read_text(), f"package {name};\n")
            self.assertEqual(tomllib.loads((package / "goml.toml").read_text())["module"]["path"], f"ecosystem::{name}")
            self.assertFalse((package / "_artifact").exists())

    def test_concurrent_publication_exposes_complete_identical_snapshots(self):
        barrier = threading.Barrier(8)

        def publish(_):
            barrier.wait(timeout=10)
            home = VERIFY.registry_snapshot()
            self.verify_complete(home)
            return home

        with ThreadPoolExecutor(max_workers=8) as executor:
            homes = list(executor.map(publish, range(8)))
        self.assertEqual(len(set(homes)), 1)
        self.assertEqual(list((self.root / "_artifact" / "registry-snapshots").glob(".registry-*")), [])

    def test_snapshot_uses_captured_bytes_when_source_changes_after_read(self):
        source = self.root / "one" / "main.gom"
        read = Path.read_bytes

        def changing_read(path):
            value = read(path)
            if path == source:
                source.write_text("package changed;\n")
            return value

        with patch.object(Path, "read_bytes", changing_read):
            first = VERIFY.registry_snapshot()
        self.verify_complete(first)
        second = VERIFY.registry_snapshot()
        self.assertNotEqual(first, second)
        self.assertEqual((second / "cache/registry/ecosystem/one/0.1.0/main.gom").read_text(), "package changed;\n")

    def test_invalid_coordinate_never_publishes_a_partial_registry(self):
        (self.root / "two" / "goml.toml").write_text('[module]\npath = "unexpected::two"\n')
        with self.assertRaisesRegex(RuntimeError, "unexpected module coordinate"):
            VERIFY.registry_snapshot()
        self.assertEqual(list((self.root / "_artifact" / "registry-snapshots").iterdir()), [])


if __name__ == "__main__":
    unittest.main()
