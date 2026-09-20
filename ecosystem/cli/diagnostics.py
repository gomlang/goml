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
        ("#[derive(cli::Subcommands)]\nstruct Invalid { value: string }", "Subcommands derive requires a nonempty enum"),
        ("#[derive(cli::Subcommands)]\nenum Invalid { Item(string, string) }", "Subcommands variants must be unit variants or carry exactly one Args value"),
        ("#[derive(cli::Subcommands)]\nenum Invalid { Item { value: string } }", "Subcommands variants must be unit variants or carry exactly one Args value"),
        ("#[derive(cli::Subcommands)]\nenum Invalid { #[command(name = \"same\")] A, #[command(alias = \"same\")] B }", "duplicate subcommand name or alias"),
        ("#[derive(cli::Subcommands)]\nenum Invalid { #[command(name = \"--bad\")] A }", "invalid subcommand name or alias"),
        ("#[derive(cli::Subcommands)]\nenum Invalid { #[command(alias = \"\")] A }", "invalid subcommand name or alias"),
        ("#[derive(cli::Subcommands)]\nenum Invalid { #[command(unknown = \"x\")] A }", "unknown subcommand attribute"),
        ("#[derive(cli::Subcommands)]\nenum Invalid { #[command(name = 1)] A }", "CLI attributes require identifiers and named string values"),
        ("#[derive(cli::Subcommands)]\n#[command(name = \"invalid\")]\nenum Invalid { A }", "Subcommands command attributes belong on variants"),
        ("#[derive(cli::Args)]\nstruct Invalid { #[arg(flatten, long = \"x\")] value: string }", "flatten and subcommand fields cannot use ordinary argument attributes"),
        ("#[derive(cli::Args)]\nstruct Invalid { #[arg(flatten, subcommand)] value: string }", "flatten and subcommand fields cannot use ordinary argument attributes"),
        ("#[derive(cli::Args)]\nstruct Invalid { #[arg(flatten)] value: Option[string] }", "flatten requires an Args type, not Option or Vec"),
        ("#[derive(cli::Args)]\nstruct Invalid { #[arg(subcommand)] value: Vec[string] }", "subcommand fields cannot have type Vec[T]"),
        ("#[derive(cli::Args)]\nstruct Invalid { #[arg(subcommand)] first: string, #[arg(subcommand)] second: string }", "Args permits only one subcommand field at each level"),
        ("#[derive(cli::Args)]\nstruct Invalid { #[arg(subcommand, optional)] value: string }", "optional subcommand fields must have type Option[T]"),
        ("#[derive(cli::Args)]\nstruct Invalid { #[arg(subcommand, required, optional)] value: Option[string] }", "CLI required and optional are mutually exclusive"),
        ("#[derive(cli::Args)]\nstruct Invalid { #[arg(optional, multiple)] value: Option[string] }", "CLI optional, multiple and count modes are mutually exclusive"),
        ("#[derive(cli::Args)]\nstruct Invalid { #[arg(optional)] value: string }", "CLI optional fields must have type Option[T]"),
        ("#[derive(cli::Args)]\nstruct Invalid { #[arg(multiple)] value: string }", "CLI multiple fields must have type Vec[T]"),
        ('#[derive(cli::Args)]\nstruct Invalid { #[arg(42)] value: string }', "CLI attributes require identifiers and named string values"),
        ('#[derive(cli::Args)]\n#[command(name = 42)]\nstruct Invalid { value: string }', "CLI attributes require identifiers and named string values"),
        ('#[derive(cli::Args)]\nstruct Invalid { #[arg(default = 42)] value: string }', "CLI attributes require identifiers and named string values"),
        ('#[derive(cli::ArgValue)]\nenum Invalid { #[value(name = 42)] Item }', "CLI attributes require identifiers and named string values"),
        ('#[derive(cli::Args)]\nstruct Invalid { #[arg(unknown = r"value")] value: string }', "unknown CLI argument attribute"),
        ('#[derive(cli::Subcommands)]\nenum Invalid { #[command(unknown = r"value")] Item }', "unknown subcommand attribute"),
        ('#[derive(cli::Args)]\nstruct Invalid { #[arg(long = "first", long = r"second")] value: string }', "duplicate CLI attribute"),
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
