window.BENCHMARK_DATA = {
  "lastUpdate": 1721129907675,
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
          "id": "40c391fc6626fb1461bf012cf83145933647c9d6",
          "message": "docs(examples): upgrade to Angular v17.x.x project-wide\nPrimary Changes\n----------------\n1. Upgraded Angular to use version 17.x.x and also its dependencies.\n\nFixes #3063\n\nSigned-off-by: aldousalvarez <aldousss.alvarez@gmail.com>",
          "timestamp": "2024-07-15T11:14:19-07:00",
          "tree_id": "cabe8ea26980044b641d12c6b77f46f489e05323",
          "url": "https://github.com/ashnashahgrover/cacti/commit/40c391fc6626fb1461bf012cf83145933647c9d6"
        },
        "date": 1721129904749,
        "tool": "benchmarkjs",
        "benches": [
          {
            "name": "cmd-api-server_HTTP_GET_getOpenApiSpecV1",
            "value": 551,
            "range": "±1.72%",
            "unit": "ops/sec",
            "extra": "175 samples"
          },
          {
            "name": "cmd-api-server_gRPC_GetOpenApiSpecV1",
            "value": 341,
            "range": "±1.56%",
            "unit": "ops/sec",
            "extra": "180 samples"
          }
        ]
      }
    ]
  }
}