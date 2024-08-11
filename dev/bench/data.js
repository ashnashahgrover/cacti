window.BENCHMARK_DATA = {
  "lastUpdate": 1723411809579,
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
      },
      {
        "commit": {
          "author": {
            "email": "peter.somogyvari@accenture.com",
            "name": "Peter Somogyvari",
            "username": "petermetz"
          },
          "committer": {
            "email": "sandeepn.official@gmail.com",
            "name": "Sandeep Nishad",
            "username": "sandeepnRES"
          },
          "distinct": true,
          "id": "d0e4539a9b106fa684cd34a6cdb1ff835b870ce4",
          "message": "ci(github): upgrade actions/github-script to 7.0.1 project-wide\n\nFixes #3458\n\nSigned-off-by: Peter Somogyvari <peter.somogyvari@accenture.com>",
          "timestamp": "2024-08-09T05:57:27+05:30",
          "tree_id": "1b12638b0ee30d8845ca2446fe5a82b172922a85",
          "url": "https://github.com/ashnashahgrover/cacti/commit/d0e4539a9b106fa684cd34a6cdb1ff835b870ce4"
        },
        "date": 1723410802278,
        "tool": "benchmarkjs",
        "benches": [
          {
            "name": "cmd-api-server_HTTP_GET_getOpenApiSpecV1",
            "value": 594,
            "range": "±1.57%",
            "unit": "ops/sec",
            "extra": "178 samples"
          },
          {
            "name": "cmd-api-server_gRPC_GetOpenApiSpecV1",
            "value": 348,
            "range": "±2.04%",
            "unit": "ops/sec",
            "extra": "182 samples"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "peter.somogyvari@accenture.com",
            "name": "Peter Somogyvari",
            "username": "petermetz"
          },
          "committer": {
            "email": "sandeepn.official@gmail.com",
            "name": "Sandeep Nishad",
            "username": "sandeepnRES"
          },
          "distinct": true,
          "id": "d0e4539a9b106fa684cd34a6cdb1ff835b870ce4",
          "message": "ci(github): upgrade actions/github-script to 7.0.1 project-wide\n\nFixes #3458\n\nSigned-off-by: Peter Somogyvari <peter.somogyvari@accenture.com>",
          "timestamp": "2024-08-09T05:57:27+05:30",
          "tree_id": "1b12638b0ee30d8845ca2446fe5a82b172922a85",
          "url": "https://github.com/ashnashahgrover/cacti/commit/d0e4539a9b106fa684cd34a6cdb1ff835b870ce4"
        },
        "date": 1723411808151,
        "tool": "benchmarkjs",
        "benches": [
          {
            "name": "plugin-ledger-connector-besu_HTTP_GET_getOpenApiSpecV1",
            "value": 756,
            "range": "±2.97%",
            "unit": "ops/sec",
            "extra": "180 samples"
          }
        ]
      }
    ]
  }
}