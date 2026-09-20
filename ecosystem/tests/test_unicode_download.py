import hashlib
import importlib.util
import io
from pathlib import Path
import unittest
from unittest.mock import patch


SPEC = importlib.util.spec_from_file_location("unicode_generator", Path(__file__).resolve().parents[1] / "unicode_text/generate.py")
GENERATOR = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(GENERATOR)


class UnicodeDownloads(unittest.TestCase):
    def setUp(self):
        self.content = b"pinned data"
        self.manifest = {"sources": {"sample.txt": {
            "url": "https://www.unicode.org/test/sample.txt",
            "sha256": hashlib.sha256(self.content).hexdigest(),
        }}}
        self.patch = patch.object(GENERATOR, "MANIFEST", self.manifest)
        self.patch.start()
        self.addCleanup(self.patch.stop)
        GENERATOR.SOURCES.clear()
        self.addCleanup(GENERATOR.SOURCES.clear)

    def test_each_run_downloads_again_without_reusing_previous_data(self):
        with patch.object(GENERATOR.urllib.request, "urlopen", side_effect=lambda *args, **kwargs: io.BytesIO(self.content)) as request:
            GENERATOR.download_sources()
            self.assertEqual(GENERATOR.source("sample.txt"), "pinned data")
            GENERATOR.download_sources()
            self.assertEqual(request.call_count, 2)
            for call in request.call_args_list:
                self.assertEqual(call.args[0].get_header("Cache-control"), "no-cache")

    def test_checksum_failure_does_not_fall_back_to_previous_sources(self):
        GENERATOR.SOURCES["sample.txt"] = self.content
        with patch.object(GENERATOR.urllib.request, "urlopen", return_value=io.BytesIO(b"changed")):
            with self.assertRaisesRegex(RuntimeError, "checksum mismatch"):
                GENERATOR.download_sources()
        with self.assertRaisesRegex(RuntimeError, "has not been downloaded"):
            GENERATOR.source("sample.txt")

    def test_oversized_responses_fail_before_checksum_acceptance(self):
        with patch.object(GENERATOR, "MAX_SOURCE_BYTES", 3):
            with patch.object(GENERATOR.urllib.request, "urlopen", return_value=io.BytesIO(self.content)):
                with self.assertRaisesRegex(RuntimeError, "download limit"):
                    GENERATOR.download_sources()
        self.assertEqual(GENERATOR.SOURCES, {})


if __name__ == "__main__":
    unittest.main()
