window.BENCHMARK_DATA = {
  "lastUpdate": 1722819640496,
  "repoUrl": "https://github.com/ashnashahgrover/cacti",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "aldousss.alvarez@gmail.com",
            "name": "aldousalvarez",
            "username": "aldousalvarez"
          },
          "committer": {
            "email": "petermetz@users.noreply.github.com",
            "name": "Peter Somogyvari",
            "username": "petermetz"
          },
          "distinct": true,
          "id": "fdce6b3222fbec1c0f05db41dd5b93fbc8a8939d",
          "message": "ci(connector-xdai): fix docker rate limit issues with openethereum image pull\n\nPrimary Changes\n----------------\n1. Migrated all the xdai connector tests to besu ledger\n   images that is being pulled from ghcr\n\nFixes #3413\n\nSigned-off-by: aldousalvarez <aldousss.alvarez@gmail.com>",
          "timestamp": "2024-08-01T10:57:28-07:00",
          "tree_id": "2fecc0a53a69a02fe88066a2239c35ccea8728a9",
          "url": "https://github.com/ashnashahgrover/cacti/commit/fdce6b3222fbec1c0f05db41dd5b93fbc8a8939d"
        },
        "date": 1722818754162,
        "tool": "benchmarkjs",
        "benches": [
          {
            "name": "cmd-api-server_HTTP_GET_getOpenApiSpecV1",
            "value": 603,
            "range": "±1.78%",
            "unit": "ops/sec",
            "extra": "176 samples"
          },
          {
            "name": "cmd-api-server_gRPC_GetOpenApiSpecV1",
            "value": 375,
            "range": "±1.53%",
            "unit": "ops/sec",
            "extra": "181 samples"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "aldousss.alvarez@gmail.com",
            "name": "aldousalvarez",
            "username": "aldousalvarez"
          },
          "committer": {
            "email": "petermetz@users.noreply.github.com",
            "name": "Peter Somogyvari",
            "username": "petermetz"
          },
          "distinct": true,
          "id": "fdce6b3222fbec1c0f05db41dd5b93fbc8a8939d",
          "message": "ci(connector-xdai): fix docker rate limit issues with openethereum image pull\n\nPrimary Changes\n----------------\n1. Migrated all the xdai connector tests to besu ledger\n   images that is being pulled from ghcr\n\nFixes #3413\n\nSigned-off-by: aldousalvarez <aldousss.alvarez@gmail.com>",
          "timestamp": "2024-08-01T10:57:28-07:00",
          "tree_id": "2fecc0a53a69a02fe88066a2239c35ccea8728a9",
          "url": "https://github.com/ashnashahgrover/cacti/commit/fdce6b3222fbec1c0f05db41dd5b93fbc8a8939d"
        },
        "date": 1722819636985,
        "tool": "benchmarkjs",
        "benches": [
          {
            "name": "plugin-ledger-connector-besu_HTTP_GET_getOpenApiSpecV1",
            "value": 684,
            "range": "±3.48%",
            "unit": "ops/sec",
            "extra": "178 samples"
          }
        ]
      }
    ]
  }
}