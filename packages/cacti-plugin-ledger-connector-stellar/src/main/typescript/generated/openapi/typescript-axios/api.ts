/* tslint:disable */
/* eslint-disable */
/**
 * Hyperledger Cacti Plugin - Connector Stellar
 * Can perform basic smart contract tasks on a Stellar ledger
 *
 * The version of the OpenAPI document: 2.0.0-rc.6
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import type { Configuration } from './configuration';
import type { AxiosPromise, AxiosInstance, AxiosRequestConfig } from 'axios';
import globalAxios from 'axios';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from './common';
import type { RequestArgs } from './base';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, BaseAPI, RequiredError } from './base';

/**
 * 
 * @export
 * @interface DeployContractV1Request
 */
export interface DeployContractV1Request {
    /**
     * The hash of the wasm code installed in the ledger to be deployed into a new instance.
     * @type {string}
     * @memberof DeployContractV1Request
     */
    'wasmHash'?: string | null;
    /**
     * A Base64-encoded string that contains the binary data of the WASM code. This buffer is used to deploy WebAssembly code to the ledger.
     * @type {string}
     * @memberof DeployContractV1Request
     */
    'wasmBuffer'?: string | null;
    /**
     * 
     * @type {TransactionInvocation}
     * @memberof DeployContractV1Request
     */
    'transactionInvocation': TransactionInvocation;
}
/**
 * 
 * @export
 * @interface DeployContractV1Response
 */
export interface DeployContractV1Response {
    /**
     * The ID of the contract that was deployed.
     * @type {string}
     * @memberof DeployContractV1Response
     */
    'contractId': string | null;
    /**
     * The hash of the wasm code installed in the ledger.
     * @type {string}
     * @memberof DeployContractV1Response
     */
    'wasmHash': string | null;
}
/**
 * 
 * @export
 * @interface RunSorobanTransactionRequest
 */
export interface RunSorobanTransactionRequest {
    /**
     * The ID of the contract that was deployed.
     * @type {string}
     * @memberof RunSorobanTransactionRequest
     */
    'contractId': string;
    /**
     * Array of strings containing the XDR of the contract specification.
     * @type {Array<string>}
     * @memberof RunSorobanTransactionRequest
     */
    'specXdr': Array<string>;
    /**
     * The method to be called on the contract.
     * @type {string}
     * @memberof RunSorobanTransactionRequest
     */
    'method': string;
    /**
     * The arguments to pass to the method.
     * @type {object}
     * @memberof RunSorobanTransactionRequest
     */
    'methodArgs'?: object | null;
    /**
     * 
     * @type {TransactionInvocation}
     * @memberof RunSorobanTransactionRequest
     */
    'transactionInvocation': TransactionInvocation;
    /**
     * Flag indicating if the transaction should be read-only.
     * @type {boolean}
     * @memberof RunSorobanTransactionRequest
     */
    'readOnly': boolean;
}
/**
 * Response object containing the result of a contract invocation or error information if it failed.
 * @export
 * @interface RunSorobanTransactionResponse
 */
export interface RunSorobanTransactionResponse {
    /**
     * The result of the invoked contract method.
     * @type {object}
     * @memberof RunSorobanTransactionResponse
     */
    'result'?: object;
}
/**
 * 
 * @export
 * @interface TransactionHeader
 */
export interface TransactionHeader {
    /**
     * The public key of the source account of the transaction.
     * @type {string}
     * @memberof TransactionHeader
     */
    'source': string;
    /**
     * The maximum base fee in stroops to be paid for the transaction.
     * @type {number}
     * @memberof TransactionHeader
     */
    'fee': number;
    /**
     * The maximum number of ledger close time in seconds that the transaction is valid for. \'0\' equals to no timeout.
     * @type {number}
     * @memberof TransactionHeader
     */
    'timeout': number;
}
/**
 * 
 * @export
 * @interface TransactionInvocation
 */
export interface TransactionInvocation {
    /**
     * 
     * @type {TransactionHeader}
     * @memberof TransactionInvocation
     */
    'header': TransactionHeader;
    /**
     * 
     * @type {Array<string>}
     * @memberof TransactionInvocation
     */
    'signers': Array<string>;
}
/**
 * 
 * @export
 * @enum {string}
 */

export const WatchBlocksV1 = {
    Subscribe: 'org.hyperledger.cacti.api.async.stellar.WatchBlocksV1.Subscribe',
    Next: 'org.hyperledger.cacti.api.async.stellar.WatchBlocksV1.Next',
    Unsubscribe: 'org.hyperledger.cacti.api.async.stellar.WatchBlocksV1.Unsubscribe',
    Error: 'org.hyperledger.cacti.api.async.stellar.WatchBlocksV1.Error',
    Complete: 'org.hyperledger.cacti.api.async.stellar.WatchBlocksV1.Complete'
} as const;

export type WatchBlocksV1 = typeof WatchBlocksV1[keyof typeof WatchBlocksV1];


/**
 * 
 * @export
 * @interface WatchBlocksV1Progress
 */
export interface WatchBlocksV1Progress {
    /**
     * 
     * @type {{ [key: string]: any; }}
     * @memberof WatchBlocksV1Progress
     */
    'blockHeader': { [key: string]: any; };
}

/**
 * DefaultApi - axios parameter creator
 * @export
 */
export const DefaultApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Deploys a smart contract to the Stellar ledger associated with the connector.
         * @param {DeployContractV1Request} [deployContractV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deployContractV1: async (deployContractV1Request?: DeployContractV1Request, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cacti-plugin-ledger-connector-stellar/deploy-contract`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(deployContractV1Request, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Retrieves the .json file that contains the OpenAPI specification for the plugin.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getOpenApiSpecV1: async (options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cacti-plugin-ledger-connector-stellar/get-open-api-spec`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Get the Prometheus Metrics
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getPrometheusMetricsV1: async (options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cacti-plugin-ledger-connector-stellar/get-prometheus-exporter-metrics`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Executes a Soroban transaction on a stellar ledger
         * @param {RunSorobanTransactionRequest} [runSorobanTransactionRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        runSorobanTransactionV1: async (runSorobanTransactionRequest?: RunSorobanTransactionRequest, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cacti-plugin-ledger-connector-stellar/run-soroban-transaction`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(runSorobanTransactionRequest, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * DefaultApi - functional programming interface
 * @export
 */
export const DefaultApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = DefaultApiAxiosParamCreator(configuration)
    return {
        /**
         * 
         * @summary Deploys a smart contract to the Stellar ledger associated with the connector.
         * @param {DeployContractV1Request} [deployContractV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async deployContractV1(deployContractV1Request?: DeployContractV1Request, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<DeployContractV1Response>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.deployContractV1(deployContractV1Request, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Retrieves the .json file that contains the OpenAPI specification for the plugin.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getOpenApiSpecV1(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getOpenApiSpecV1(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Get the Prometheus Metrics
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getPrometheusMetricsV1(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getPrometheusMetricsV1(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Executes a Soroban transaction on a stellar ledger
         * @param {RunSorobanTransactionRequest} [runSorobanTransactionRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async runSorobanTransactionV1(runSorobanTransactionRequest?: RunSorobanTransactionRequest, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<RunSorobanTransactionResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.runSorobanTransactionV1(runSorobanTransactionRequest, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
    }
};

/**
 * DefaultApi - factory interface
 * @export
 */
export const DefaultApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = DefaultApiFp(configuration)
    return {
        /**
         * 
         * @summary Deploys a smart contract to the Stellar ledger associated with the connector.
         * @param {DeployContractV1Request} [deployContractV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deployContractV1(deployContractV1Request?: DeployContractV1Request, options?: any): AxiosPromise<DeployContractV1Response> {
            return localVarFp.deployContractV1(deployContractV1Request, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Retrieves the .json file that contains the OpenAPI specification for the plugin.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getOpenApiSpecV1(options?: any): AxiosPromise<string> {
            return localVarFp.getOpenApiSpecV1(options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Get the Prometheus Metrics
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getPrometheusMetricsV1(options?: any): AxiosPromise<string> {
            return localVarFp.getPrometheusMetricsV1(options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Executes a Soroban transaction on a stellar ledger
         * @param {RunSorobanTransactionRequest} [runSorobanTransactionRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        runSorobanTransactionV1(runSorobanTransactionRequest?: RunSorobanTransactionRequest, options?: any): AxiosPromise<RunSorobanTransactionResponse> {
            return localVarFp.runSorobanTransactionV1(runSorobanTransactionRequest, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * DefaultApi - object-oriented interface
 * @export
 * @class DefaultApi
 * @extends {BaseAPI}
 */
export class DefaultApi extends BaseAPI {
    /**
     * 
     * @summary Deploys a smart contract to the Stellar ledger associated with the connector.
     * @param {DeployContractV1Request} [deployContractV1Request] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public deployContractV1(deployContractV1Request?: DeployContractV1Request, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).deployContractV1(deployContractV1Request, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Retrieves the .json file that contains the OpenAPI specification for the plugin.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getOpenApiSpecV1(options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).getOpenApiSpecV1(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Get the Prometheus Metrics
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getPrometheusMetricsV1(options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).getPrometheusMetricsV1(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Executes a Soroban transaction on a stellar ledger
     * @param {RunSorobanTransactionRequest} [runSorobanTransactionRequest] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public runSorobanTransactionV1(runSorobanTransactionRequest?: RunSorobanTransactionRequest, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).runSorobanTransactionV1(runSorobanTransactionRequest, options).then((request) => request(this.axios, this.basePath));
    }
}


