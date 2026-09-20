# Compile-fail cases

The 33 declarations and expected diagnostic substrings are maintained specification examples, migrated without changing their expected results. `diagnostics_test.gom` creates an isolated downstream project under `_artifact`, invokes the repository GoML driver with `std::process`, checks the nonzero exit and diagnostic, and removes the project. `GOML_VERIFY_DRIVER` can select an alternative driver.
