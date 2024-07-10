window.BENCHMARK_DATA = {
  "lastUpdate": 1720585795937,
  "repoUrl": "https://github.com/ashnashahgrover/cacti",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "peter.somogyvari@accenture.com",
            "name": "Peter Somogyvari",
            "username": "petermetz"
          },
          "committer": {
            "email": "petermetz@users.noreply.github.com",
            "name": "Peter Somogyvari",
            "username": "petermetz"
          },
          "distinct": true,
          "id": "5f644a6081d11108fc354f040d2c682074f393c8",
          "message": "docs(connector-fabric): remove references to old Fabric v1.4.x tests\n\n1. We've removed the Fabric v1.4.x test cases a while back but I forgot\nto delete them from the .taprc and jest.config.js files for example.\n2. Also in a couple of places the documentation was referring to the\ntest cases as well so that had to be updated also.\n\nSigned-off-by: Peter Somogyvari <peter.somogyvari@accenture.com>",
          "timestamp": "2024-07-09T21:03:35-07:00",
          "tree_id": "cae98449766f283967d505d31e8299c99f808d31",
          "url": "https://github.com/ashnashahgrover/cacti/commit/5f644a6081d11108fc354f040d2c682074f393c8"
        },
        "date": 1720585793956,
        "tool": "benchmarkjs",
        "benches": [
          {
            "name": "cmd-api-server_HTTP_GET_getOpenApiSpecV1",
            "value": 586,
            "range": "±1.77%",
            "unit": "ops/sec",
            "extra": "176 samples"
          },
          {
            "name": "cmd-api-server_gRPC_GetOpenApiSpecV1",
            "value": 357,
            "range": "±1.38%",
            "unit": "ops/sec",
            "extra": "181 samples"
          }
        ]
      }
    ]
  }
}