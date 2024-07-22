window.BENCHMARK_DATA = {
  "lastUpdate": 1721655587792,
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
          "id": "497ea3226631fdcad763e6281ee058d91ca01988",
          "message": "test(test-tooling): add container image builder utilities\n\n1. Currently our integration tests depend on pre-published container\nimages to be on the official registry (ghcr.io). This has pros and cons.\nThe pro is that we can pin the tests to a specific ledger version and\nthen have confidence that the test code works with that specific image.\nOn the other hand if the image itself has problems we won't know it until\nafter it was published and then tests were executed with it (unless we\nperform manual testing which is a lot of effrot as it requires the\nmanual modification of the test cases).\n2. In order to gives us the ability to test against the container image\ndefinitions as they are in the current revision of the source code,\nwe are adding here a couple of utility functions to streamline writing\ntest cases that build the container images for themselves as part of the\ntest case.\n\nAn example of how to use it in a test case:\n\n```typescript\nconst imgConnectorJvm = await buildImageConnectorCordaServer({\n    logLevel,\n});\n\n// ...\n\nconnector = new CordaConnectorContainer({\n    logLevel,\n    imageName: imgConnectorJvm.imageName,\n    imageVersion: imgConnectorJvm.imageVersion,\n    envVars: [envVarSpringAppJson],\n});\n\n```\n\nSigned-off-by: Peter Somogyvari <peter.somogyvari@accenture.com>",
          "timestamp": "2024-07-17T17:31:23-07:00",
          "tree_id": "59d3d8a612cce5ee7e4b23eb014491baec319e68",
          "url": "https://github.com/ashnashahgrover/cacti/commit/497ea3226631fdcad763e6281ee058d91ca01988"
        },
        "date": 1721284236220,
        "tool": "benchmarkjs",
        "benches": [
          {
            "name": "cmd-api-server_HTTP_GET_getOpenApiSpecV1",
            "value": 586,
            "range": "±1.66%",
            "unit": "ops/sec",
            "extra": "177 samples"
          },
          {
            "name": "cmd-api-server_gRPC_GetOpenApiSpecV1",
            "value": 355,
            "range": "±1.39%",
            "unit": "ops/sec",
            "extra": "181 samples"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "committer": {
            "email": "petermetz@users.noreply.github.com",
            "name": "Peter Somogyvari",
            "username": "petermetz"
          },
          "distinct": true,
          "id": "603ff0ea5b6243d8f1adf43e264d0524fa31c454",
          "message": "build: bump curve25519-dalek\n\nBumps the cargo group with 1 update in the\n/packages/cacti-plugin-ledger-connector-stellar/src/test/rust/demo-contract\ndirectory: [curve25519-dalek](https://github.com/dalek-cryptography/curve25519-dalek).\n\nUpdates `curve25519-dalek` from 4.1.1 to 4.1.2\n- [Release notes](https://github.com/dalek-cryptography/curve25519-dalek/releases)\n- Commits:\nhttps://github.com/dalek-cryptography/curve25519-dalek/compare/curve25519-4.1.1...curve25519-4.1.2\n\n---\nupdated-dependencies:\n- dependency-name: curve25519-dalek\n  dependency-type: indirect\n  dependency-group: cargo\n...\n\nCo-authored-by: Peter Somogyvari <peter.somogyvari@accenture.com>\n\nSigned-off-by: dependabot[bot] <support@github.com>\nSigned-off-by: Peter Somogyvari <peter.somogyvari@accenture.com>",
          "timestamp": "2024-07-19T14:21:03-07:00",
          "tree_id": "3b2280f1255fb5e886c8672b7dc9d151ae16d54e",
          "url": "https://github.com/ashnashahgrover/cacti/commit/603ff0ea5b6243d8f1adf43e264d0524fa31c454"
        },
        "date": 1721654694715,
        "tool": "benchmarkjs",
        "benches": [
          {
            "name": "cmd-api-server_HTTP_GET_getOpenApiSpecV1",
            "value": 552,
            "range": "±1.79%",
            "unit": "ops/sec",
            "extra": "177 samples"
          },
          {
            "name": "cmd-api-server_gRPC_GetOpenApiSpecV1",
            "value": 344,
            "range": "±1.26%",
            "unit": "ops/sec",
            "extra": "180 samples"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "committer": {
            "email": "petermetz@users.noreply.github.com",
            "name": "Peter Somogyvari",
            "username": "petermetz"
          },
          "distinct": true,
          "id": "603ff0ea5b6243d8f1adf43e264d0524fa31c454",
          "message": "build: bump curve25519-dalek\n\nBumps the cargo group with 1 update in the\n/packages/cacti-plugin-ledger-connector-stellar/src/test/rust/demo-contract\ndirectory: [curve25519-dalek](https://github.com/dalek-cryptography/curve25519-dalek).\n\nUpdates `curve25519-dalek` from 4.1.1 to 4.1.2\n- [Release notes](https://github.com/dalek-cryptography/curve25519-dalek/releases)\n- Commits:\nhttps://github.com/dalek-cryptography/curve25519-dalek/compare/curve25519-4.1.1...curve25519-4.1.2\n\n---\nupdated-dependencies:\n- dependency-name: curve25519-dalek\n  dependency-type: indirect\n  dependency-group: cargo\n...\n\nCo-authored-by: Peter Somogyvari <peter.somogyvari@accenture.com>\n\nSigned-off-by: dependabot[bot] <support@github.com>\nSigned-off-by: Peter Somogyvari <peter.somogyvari@accenture.com>",
          "timestamp": "2024-07-19T14:21:03-07:00",
          "tree_id": "3b2280f1255fb5e886c8672b7dc9d151ae16d54e",
          "url": "https://github.com/ashnashahgrover/cacti/commit/603ff0ea5b6243d8f1adf43e264d0524fa31c454"
        },
        "date": 1721655584803,
        "tool": "benchmarkjs",
        "benches": [
          {
            "name": "plugin-ledger-connector-besu_HTTP_GET_getOpenApiSpecV1",
            "value": 652,
            "range": "±3.11%",
            "unit": "ops/sec",
            "extra": "178 samples"
          }
        ]
      }
    ]
  }
}