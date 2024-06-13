window.BENCHMARK_DATA = {
  "lastUpdate": 1718275841826,
  "repoUrl": "https://github.com/ashnashahgrover/cacti",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "vramakr2@in.ibm.com",
            "name": "VRamakrishna",
            "username": "VRamakrishna"
          },
          "committer": {
            "email": "sandeepn.official@gmail.com",
            "name": "Sandeep Nishad",
            "username": "sandeepnRES"
          },
          "distinct": true,
          "id": "cc2f9c58975f72907576c529c2094b330d8b2b92",
          "message": "ci(ossf): configured GitHub action and badge for OpenSSF scorecard\n\nSigned-off-by: VRamakrishna <vramakr2@in.ibm.com>",
          "timestamp": "2024-06-10T01:11:43+05:30",
          "tree_id": "8f80768252cab6792094c2872aebb5438315caff",
          "url": "https://github.com/ashnashahgrover/cacti/commit/cc2f9c58975f72907576c529c2094b330d8b2b92"
        },
        "date": 1718095028565,
        "tool": "benchmarkjs",
        "benches": [
          {
            "name": "cmd-api-server_HTTP_GET_getOpenApiSpecV1",
            "value": 573,
            "range": "±1.75%",
            "unit": "ops/sec",
            "extra": "177 samples"
          },
          {
            "name": "cmd-api-server_gRPC_GetOpenApiSpecV1",
            "value": 353,
            "range": "±1.24%",
            "unit": "ops/sec",
            "extra": "181 samples"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "vramakr2@in.ibm.com",
            "name": "VRamakrishna",
            "username": "VRamakrishna"
          },
          "committer": {
            "email": "sandeepn.official@gmail.com",
            "name": "Sandeep Nishad",
            "username": "sandeepnRES"
          },
          "distinct": true,
          "id": "cc2f9c58975f72907576c529c2094b330d8b2b92",
          "message": "ci(ossf): configured GitHub action and badge for OpenSSF scorecard\n\nSigned-off-by: VRamakrishna <vramakr2@in.ibm.com>",
          "timestamp": "2024-06-10T01:11:43+05:30",
          "tree_id": "8f80768252cab6792094c2872aebb5438315caff",
          "url": "https://github.com/ashnashahgrover/cacti/commit/cc2f9c58975f72907576c529c2094b330d8b2b92"
        },
        "date": 1718097815139,
        "tool": "benchmarkjs",
        "benches": [
          {
            "name": "plugin-ledger-connector-besu_HTTP_GET_getOpenApiSpecV1",
            "value": 715,
            "range": "±3.16%",
            "unit": "ops/sec",
            "extra": "176 samples"
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
          "id": "ec8123cf954b09ba8cb213c7332dfe82224c351f",
          "message": "feat(connector-fabric): drop support for Fabric v1.x\n\nBREAKING CHANGE: The Open API specification that has the enums for\nledger versions will no longer have an option for Fabric v1.x\nThis means that in the core-api package the LedgerType enum has changes\nwhich means that code that depends on that enum value will need to be\nupdated.\n\nFabric v1.x has had unmaintained dependencies associated with it such as\nthe native grpc package that stopped receiving security updates years ago\nand therefore it's dangerous to have around.\n\nThere are also some issues with Fabric v1.x that make the AIO image flaky\nwhich also makes the relevant tests flaky due to which we couldn't run\nthe v1.x Fabric tests on the CI for a while now anyway.\n\nIn order to reduce the CI resource usage and our own maintenance burden\nI suggest that we get rid of the Fabric v1.x support meaning that we can\neliminate the AIO image build and some code complexity from the test ledger\ncode as well.\n\nIn addition some old fixtures can be removed that the tests were using.\nOverall a net-positive as deleting code without losing functionality (that\nwe care about) is always a plus.\n\nSigned-off-by: Peter Somogyvari <peter.somogyvari@accenture.com>",
          "timestamp": "2024-06-12T11:12:42-07:00",
          "tree_id": "cfaf922e1ca373842ba7d01f923a3ad75d880f8a",
          "url": "https://github.com/ashnashahgrover/cacti/commit/ec8123cf954b09ba8cb213c7332dfe82224c351f"
        },
        "date": 1718274867040,
        "tool": "benchmarkjs",
        "benches": [
          {
            "name": "cmd-api-server_HTTP_GET_getOpenApiSpecV1",
            "value": 579,
            "range": "±1.64%",
            "unit": "ops/sec",
            "extra": "179 samples"
          },
          {
            "name": "cmd-api-server_gRPC_GetOpenApiSpecV1",
            "value": 362,
            "range": "±1.42%",
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
          "id": "ec8123cf954b09ba8cb213c7332dfe82224c351f",
          "message": "feat(connector-fabric): drop support for Fabric v1.x\n\nBREAKING CHANGE: The Open API specification that has the enums for\nledger versions will no longer have an option for Fabric v1.x\nThis means that in the core-api package the LedgerType enum has changes\nwhich means that code that depends on that enum value will need to be\nupdated.\n\nFabric v1.x has had unmaintained dependencies associated with it such as\nthe native grpc package that stopped receiving security updates years ago\nand therefore it's dangerous to have around.\n\nThere are also some issues with Fabric v1.x that make the AIO image flaky\nwhich also makes the relevant tests flaky due to which we couldn't run\nthe v1.x Fabric tests on the CI for a while now anyway.\n\nIn order to reduce the CI resource usage and our own maintenance burden\nI suggest that we get rid of the Fabric v1.x support meaning that we can\neliminate the AIO image build and some code complexity from the test ledger\ncode as well.\n\nIn addition some old fixtures can be removed that the tests were using.\nOverall a net-positive as deleting code without losing functionality (that\nwe care about) is always a plus.\n\nSigned-off-by: Peter Somogyvari <peter.somogyvari@accenture.com>",
          "timestamp": "2024-06-12T11:12:42-07:00",
          "tree_id": "cfaf922e1ca373842ba7d01f923a3ad75d880f8a",
          "url": "https://github.com/ashnashahgrover/cacti/commit/ec8123cf954b09ba8cb213c7332dfe82224c351f"
        },
        "date": 1718275838906,
        "tool": "benchmarkjs",
        "benches": [
          {
            "name": "plugin-ledger-connector-besu_HTTP_GET_getOpenApiSpecV1",
            "value": 750,
            "range": "±2.97%",
            "unit": "ops/sec",
            "extra": "181 samples"
          }
        ]
      }
    ]
  }
}