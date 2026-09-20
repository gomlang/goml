import os
from pathlib import Path
import subprocess
import sys
import tempfile


ROOT = Path(__file__).resolve().parents[1]
sys.path.insert(0, str(ROOT))
import verify


def main():
    goml = str(ROOT.parent / "stage2" / "bin" / "goml")
    environment = os.environ.copy()
    environment["GOML_HOME"] = str(verify.registry_snapshot())
    cases = [
        ("#[derive(cli::Args)]\nenum Invalid { Item }", "Args derive requires a struct"),
        ("#[derive(cli::ArgValue)]\nenum Invalid { Item(string) }", "ArgValue variants must have no payload"),
        ("#[derive(cli::Args)]\nstruct Invalid { #[arg(typo)] value: string }", "unknown CLI argument switch"),
        ("#[derive(cli::Args)]\nstruct Invalid { #[arg(short = \"abc\")] value: string }", "CLI short options require one ASCII character"),
        ("#[derive(cli::Args)]\nstruct Invalid { #[arg(count)] value: bool }", "CLI count fields must have type isize"),
        ("#[derive(cli::ArgValue)]\nenum Invalid { #[value(name = \"same\")] A, #[value(alias = \"same\")] B }", "duplicate enum argument name or alias"),
        ("#[derive(cli::Args)]\n#[command(unknown = \"x\")]\nstruct Invalid { value: string }", "unknown command attribute"),
    ]
    artifacts = ROOT / "_artifact" / "diagnostics"
    artifacts.mkdir(parents=True, exist_ok=True)
    with tempfile.TemporaryDirectory(prefix="cli-", dir=artifacts) as temporary:
        for index, (declaration, expected) in enumerate(cases):
            directory = Path(temporary) / str(index)
            directory.mkdir()
            (directory / "goml.toml").write_text('[module]\npath = "diagnostics::cli"\n[dependencies]\n"ecosystem::cli" = "0.1.0"\n')
            (directory / "main.gom").write_text("package main;\nuse ecosystem::cli;\n" + declaration + "\nfn main() -> () {}\n")
            subprocess.run([goml, "fmt"], cwd=directory, env=environment, check=True)
            result = subprocess.run([goml, "check"], cwd=directory, env=environment, capture_output=True, text=True)
            output = result.stdout + result.stderr
            if result.returncode == 0 or expected not in output:
                raise AssertionError(f"case {index} expected {expected!r}:\n{output}")
            (artifacts / f"cli-{index}.log").write_text(output)
    print(f"CLI derive diagnostics: {len(cases)} cases passed")


if __name__ == "__main__":
    main()
