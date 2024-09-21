window.BENCHMARK_DATA = {
  "lastUpdate": 1726881367500,
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
          "id": "299af74b0e74d5bdb0224fb5a80303ef75024fdb",
          "message": "docs(api-server): explain local plugin import through packageSrc config\n\nThis came up during a discussion here and I thought it best to document\nit a little more thoroughly so that later it can be referenced for others\nas well:\nhttps://github.com/hyperledger/cacti/issues/3406#issuecomment-2299654552\n\nSigned-off-by: Peter Somogyvari <peter.somogyvari@accenture.com>",
          "timestamp": "2024-09-17T19:15:15-07:00",
          "tree_id": "a6698ef08ae1c594584960b0297cf549de9a8ec7",
          "url": "https://github.com/ashnashahgrover/cacti/commit/299af74b0e74d5bdb0224fb5a80303ef75024fdb"
        },
        "date": 1726880439569,
        "tool": "benchmarkjs",
        "benches": [
          {
            "name": "cmd-api-server_HTTP_GET_getOpenApiSpecV1",
            "value": 573,
            "range": "±1.74%",
            "unit": "ops/sec",
            "extra": "178 samples"
          },
          {
            "name": "cmd-api-server_gRPC_GetOpenApiSpecV1",
            "value": 699,
            "range": "±3.06%",
            "unit": "ops/sec",
            "extra": "181 samples"
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
            "email": "petermetz@users.noreply.github.com",
            "name": "Peter Somogyvari",
            "username": "petermetz"
          },
          "distinct": true,
          "id": "299af74b0e74d5bdb0224fb5a80303ef75024fdb",
          "message": "docs(api-server): explain local plugin import through packageSrc config\n\nThis came up during a discussion here and I thought it best to document\nit a little more thoroughly so that later it can be referenced for others\nas well:\nhttps://github.com/hyperledger/cacti/issues/3406#issuecomment-2299654552\n\nSigned-off-by: Peter Somogyvari <peter.somogyvari@accenture.com>",
          "timestamp": "2024-09-17T19:15:15-07:00",
          "tree_id": "a6698ef08ae1c594584960b0297cf549de9a8ec7",
          "url": "https://github.com/ashnashahgrover/cacti/commit/299af74b0e74d5bdb0224fb5a80303ef75024fdb"
        },
        "date": 1726881365067,
        "tool": "benchmarkjs",
        "benches": [
          {
            "name": "plugin-ledger-connector-besu_HTTP_GET_getOpenApiSpecV1",
            "value": 719,
            "range": "±2.82%",
            "unit": "ops/sec",
            "extra": "180 samples"
          }
        ]
      }
    ]
  }
}