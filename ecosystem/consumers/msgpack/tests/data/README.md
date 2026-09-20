# Independent reference fixtures

Reference: [msgpack-python 1.1.2](https://files.pythonhosted.org/packages/4d/f2/bfb55a6236ed8725a96b0aa3acbd0ec17588e6a2c3b62a93eb513ed8783f/msgpack-1.1.2.tar.gz).

2,490 cases covering all tags, truncations, binary lengths, float bit patterns, recursive values, timestamps, typed Serde layouts, options and streams. The input generator used seed 20260920. Expected bytes come from the reference packer, with specification-derived float32 bit preservation and malformed-input rejection expectations.

Reference archive SHA-256: `3b60763c1373dd60f398488069bcdc703cd08a711477b5d480eecc9f9626f47e`.

The fixture was captured once during migration of the verification harness. Normal tests read it directly with GoML; no Python interpreter, package download or reference runtime is required. Inputs and expected values are independent of the GoML implementation.

Fixture SHA-256: `8c0c0b7e18a062fcfa09ccaed4da0fb7b44d8d17b54ad93ac9a1b628cea3159f`.
