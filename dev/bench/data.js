window.BENCHMARK_DATA = {
  "lastUpdate": 1725323480871,
  "repoUrl": "https://github.com/ashnashahgrover/cacti",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "adrian.batuto@accenture.com",
            "name": "adrianbatuto",
            "username": "adrianbatuto"
          },
          "committer": {
            "email": "petermetz@users.noreply.github.com",
            "name": "Peter Somogyvari",
            "username": "petermetz"
          },
          "distinct": true,
          "id": "ec9683d38670fe5d657b602db8215e602fd4209d",
          "message": "feat(corda): support 5.1 via TS/HTTP (no JVM)\nFixes #2978\nFixes #3293\n\nSigned-off-by: adrianbatuto <adrian.batuto@accenture.com>",
          "timestamp": "2024-08-19T14:43:10-07:00",
          "tree_id": "530c66f1928ba9481fcc2d1d760582bf58be6677",
          "url": "https://github.com/ashnashahgrover/cacti/commit/ec9683d38670fe5d657b602db8215e602fd4209d"
        },
        "date": 1724279181985,
        "tool": "benchmarkjs",
        "benches": [
          {
            "name": "cmd-api-server_HTTP_GET_getOpenApiSpecV1",
            "value": 588,
            "range": "±1.67%",
            "unit": "ops/sec",
            "extra": "177 samples"
          },
          {
            "name": "cmd-api-server_gRPC_GetOpenApiSpecV1",
            "value": 342,
            "range": "±1.95%",
            "unit": "ops/sec",
            "extra": "181 samples"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "adrian.batuto@accenture.com",
            "name": "adrianbatuto",
            "username": "adrianbatuto"
          },
          "committer": {
            "email": "petermetz@users.noreply.github.com",
            "name": "Peter Somogyvari",
            "username": "petermetz"
          },
          "distinct": true,
          "id": "ec9683d38670fe5d657b602db8215e602fd4209d",
          "message": "feat(corda): support 5.1 via TS/HTTP (no JVM)\nFixes #2978\nFixes #3293\n\nSigned-off-by: adrianbatuto <adrian.batuto@accenture.com>",
          "timestamp": "2024-08-19T14:43:10-07:00",
          "tree_id": "530c66f1928ba9481fcc2d1d760582bf58be6677",
          "url": "https://github.com/ashnashahgrover/cacti/commit/ec9683d38670fe5d657b602db8215e602fd4209d"
        },
        "date": 1724279960184,
        "tool": "benchmarkjs",
        "benches": [
          {
            "name": "plugin-ledger-connector-besu_HTTP_GET_getOpenApiSpecV1",
            "value": 724,
            "range": "±2.71%",
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
            "email": "petermetz@users.noreply.github.com",
            "name": "Peter Somogyvari",
            "username": "petermetz"
          },
          "distinct": true,
          "id": "6872ec53485b3b960fc3f88d162510646a2272ab",
          "message": "test(test-tooling): jest migrate postgres container constructor options\n\nSigned-off-by: Peter Somogyvari <peter.somogyvari@accenture.com>",
          "timestamp": "2024-08-29T09:58:36-07:00",
          "tree_id": "9e125a93f6cef4fcdec1584946020ae65b86746f",
          "url": "https://github.com/ashnashahgrover/cacti/commit/6872ec53485b3b960fc3f88d162510646a2272ab"
        },
        "date": 1724972101759,
        "tool": "benchmarkjs",
        "benches": [
          {
            "name": "cmd-api-server_HTTP_GET_getOpenApiSpecV1",
            "value": 585,
            "range": "±1.72%",
            "unit": "ops/sec",
            "extra": "176 samples"
          },
          {
            "name": "cmd-api-server_gRPC_GetOpenApiSpecV1",
            "value": 344,
            "range": "±2.03%",
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
          "id": "6872ec53485b3b960fc3f88d162510646a2272ab",
          "message": "test(test-tooling): jest migrate postgres container constructor options\n\nSigned-off-by: Peter Somogyvari <peter.somogyvari@accenture.com>",
          "timestamp": "2024-08-29T09:58:36-07:00",
          "tree_id": "9e125a93f6cef4fcdec1584946020ae65b86746f",
          "url": "https://github.com/ashnashahgrover/cacti/commit/6872ec53485b3b960fc3f88d162510646a2272ab"
        },
        "date": 1724972961686,
        "tool": "benchmarkjs",
        "benches": [
          {
            "name": "plugin-ledger-connector-besu_HTTP_GET_getOpenApiSpecV1",
            "value": 709,
            "range": "±2.95%",
            "unit": "ops/sec",
            "extra": "179 samples"
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
          "id": "957da7c3e1d80068391485a825ba6bb1e68333ac",
          "message": "feat(core-api): add createIsJwsGeneralTypeGuard, createAjvTypeGuard<T>\n\n1. createAjvTypeGuard<T>() is the lower level utility which can be used\nto construct the more convenient, higher level type predicates/type guards\nsuch as createIsJwsGeneralTypeGuard() which uses createAjvTypeGuard<JwsGeneral>\nunder the hood.\n2. This commit is also meant to be establishing a larger, more generic\npattern of us being able to create type guards out of the Open API specs\nin a convenient way instead of having to write the validation code by hand.\n\nAn example usage of the new createAjvTypeGuard<T>() utility is the\ncreateIsJwsGeneralTypeGuard() function itself.\n\nAn example usage of the new createIsJwsGeneralTypeGuard() can be found\nin packages/cactus-plugin-consortium-manual/src/main/typescript/plugin-consortium-manual.ts\n\nThe code documentation contains examples as well for maximum discoverabilty\nand I'll also include it here:\n\n```typescript\nimport { JWSGeneral } from \"@hyperledger/cactus-core-api\";\nimport { createIsJwsGeneralTypeGuard } from \"@hyperledger/cactus-core-api\";\n\nexport class PluginConsortiumManual {\n  private readonly isJwsGeneral: (x: unknown) => x is JWSGeneral;\n\n  constructor() {\n    // Creating the type-guard function is relatively costly due to the Ajv schema\n    // compilation that needs to happen as part of it so it is good practice to\n    // cache the type-guard function as much as possible, for examle by adding it\n    // as a class member on a long-lived object such as a plugin instance which is\n    // expected to match the life-cycle of the API server NodeJS process itself.\n    // The specific anti-pattern would be to create a new type-guard function\n    // for each request received by a plugin as this would affect performance\n    // negatively.\n    this.isJwsGeneral = createIsJwsGeneralTypeGuard();\n  }\n\n  public async getNodeJws(): Promise<JWSGeneral> {\n    // rest of the implementation that produces a JWS ...\n    const jws = await joseGeneralSign.sign();\n\n    if (!this.isJwsGeneral(jws)) {\n      throw new TypeError(\"Jose GeneralSign.sign() gave non-JWSGeneral type\");\n    }\n    return jws;\n  }\n}\n```\n\nRelevant discussion took place here:\nhttps://github.com/hyperledger/cacti/pull/3471#discussion_r1731894747\n\nSigned-off-by: Peter Somogyvari <peter.somogyvari@accenture.com>",
          "timestamp": "2024-09-02T08:23:56-07:00",
          "tree_id": "f0a96975cecbe783530ab2031beadee53efbd0b2",
          "url": "https://github.com/ashnashahgrover/cacti/commit/957da7c3e1d80068391485a825ba6bb1e68333ac"
        },
        "date": 1725323479595,
        "tool": "benchmarkjs",
        "benches": [
          {
            "name": "cmd-api-server_HTTP_GET_getOpenApiSpecV1",
            "value": 594,
            "range": "±1.64%",
            "unit": "ops/sec",
            "extra": "178 samples"
          },
          {
            "name": "cmd-api-server_gRPC_GetOpenApiSpecV1",
            "value": 348,
            "range": "±1.91%",
            "unit": "ops/sec",
            "extra": "179 samples"
          }
        ]
      }
    ]
  }
}