/* tslint:disable */
/* eslint-disable */
/**
 * Hyperledger Cactus Plugin - Connector Polkadot
 * Can perform basic tasks on a Polkadot parachain
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
 * @interface DeployContractInkRequest
 */
export interface DeployContractInkRequest {
    /**
     * 
     * @type {Web3SigningCredential}
     * @memberof DeployContractInkRequest
     */
    'web3SigningCredential': Web3SigningCredential;
    /**
     * raw wasm for the compiled contract in base64 format
     * @type {string}
     * @memberof DeployContractInkRequest
     */
    'wasm': string;
    /**
     * 
     * @type {PolkadotTransactionConfigTransferSubmittable}
     * @memberof DeployContractInkRequest
     */
    'constructorMethod'?: PolkadotTransactionConfigTransferSubmittable;
    /**
     * 
     * @type {PolkadotTransactionConfigTransferSubmittable}
     * @memberof DeployContractInkRequest
     */
    'metadata': PolkadotTransactionConfigTransferSubmittable;
    /**
     * 
     * @type {DeployContractInkRequestGasLimit}
     * @memberof DeployContractInkRequest
     */
    'gasLimit': DeployContractInkRequestGasLimit;
    /**
     * 
     * @type {DeployContractInkRequestStorageDepositLimit}
     * @memberof DeployContractInkRequest
     */
    'storageDepositLimit'?: DeployContractInkRequestStorageDepositLimit | null;
    /**
     * The list of arguments to pass in to the contract method being deployed
     * @type {Array<any>}
     * @memberof DeployContractInkRequest
     */
    'params'?: Array<any>;
    /**
     * 
     * @type {DeployContractInkRequestBalance}
     * @memberof DeployContractInkRequest
     */
    'balance'?: DeployContractInkRequestBalance;
    /**
     * 
     * @type {DeployContractInkRequestSalt}
     * @memberof DeployContractInkRequest
     */
    'salt'?: DeployContractInkRequestSalt | null;
}
/**
 * @type DeployContractInkRequestBalance
 * @export
 */
export type DeployContractInkRequestBalance = number | string;

/**
 * 
 * @export
 * @interface DeployContractInkRequestGasLimit
 */
export interface DeployContractInkRequestGasLimit {
    /**
     * 
     * @type {number}
     * @memberof DeployContractInkRequestGasLimit
     */
    'refTime': number;
    /**
     * 
     * @type {number}
     * @memberof DeployContractInkRequestGasLimit
     */
    'proofSize': number;
}
/**
 * @type DeployContractInkRequestSalt
 * @export
 */
export type DeployContractInkRequestSalt = string;

/**
 * @type DeployContractInkRequestStorageDepositLimit
 * @export
 */
export type DeployContractInkRequestStorageDepositLimit = number | string;

/**
 * 
 * @export
 * @interface DeployContractInkResponse
 */
export interface DeployContractInkResponse {
    /**
     * 
     * @type {boolean}
     * @memberof DeployContractInkResponse
     */
    'success': boolean;
    /**
     * 
     * @type {string}
     * @memberof DeployContractInkResponse
     */
    'contractAddress'?: string;
}
/**
 * Error response from the connector.
 * @export
 * @interface ErrorExceptionResponse
 */
export interface ErrorExceptionResponse {
    /**
     * Short error description message.
     * @type {string}
     * @memberof ErrorExceptionResponse
     */
    'message': string;
    /**
     * Detailed error information.
     * @type {string}
     * @memberof ErrorExceptionResponse
     */
    'error': string;
}
/**
 * 
 * @export
 * @interface InvokeContractRequest
 */
export interface InvokeContractRequest {
    /**
     * 
     * @type {PolkadotContractInvocationType}
     * @memberof InvokeContractRequest
     */
    'invocationType': PolkadotContractInvocationType;
    /**
     * 
     * @type {string}
     * @memberof InvokeContractRequest
     */
    'accountAddress': string;
    /**
     * 
     * @type {Web3SigningCredential}
     * @memberof InvokeContractRequest
     */
    'web3SigningCredential': Web3SigningCredential;
    /**
     * 
     * @type {PolkadotTransactionConfigTransferSubmittable}
     * @memberof InvokeContractRequest
     */
    'metadata': PolkadotTransactionConfigTransferSubmittable;
    /**
     * 
     * @type {string}
     * @memberof InvokeContractRequest
     */
    'contractAddress': string;
    /**
     * The name of the contract method to invoke.
     * @type {string}
     * @memberof InvokeContractRequest
     */
    'methodName': string;
    /**
     * 
     * @type {DeployContractInkRequestGasLimit}
     * @memberof InvokeContractRequest
     */
    'gasLimit': DeployContractInkRequestGasLimit;
    /**
     * 
     * @type {DeployContractInkRequestStorageDepositLimit}
     * @memberof InvokeContractRequest
     */
    'storageDepositLimit'?: DeployContractInkRequestStorageDepositLimit | null;
    /**
     * 
     * @type {DeployContractInkRequestBalance}
     * @memberof InvokeContractRequest
     */
    'balance'?: DeployContractInkRequestBalance;
    /**
     * The list of arguments to pass in to the contract method being invoked
     * @type {Array<any>}
     * @memberof InvokeContractRequest
     */
    'params'?: Array<any>;
}


/**
 * 
 * @export
 * @interface InvokeContractResponse
 */
export interface InvokeContractResponse {
    /**
     * 
     * @type {any}
     * @memberof InvokeContractResponse
     */
    'callOutput'?: any;
    /**
     * 
     * @type {boolean}
     * @memberof InvokeContractResponse
     */
    'success': boolean;
    /**
     * 
     * @type {string}
     * @memberof InvokeContractResponse
     */
    'txHash'?: string;
    /**
     * 
     * @type {string}
     * @memberof InvokeContractResponse
     */
    'blockHash'?: string;
}
/**
 * 
 * @export
 * @enum {string}
 */

export const PolkadotContractInvocationType = {
    Send: 'SEND',
    Query: 'QUERY'
} as const;

export type PolkadotContractInvocationType = typeof PolkadotContractInvocationType[keyof typeof PolkadotContractInvocationType];


/**
 * 
 * @export
 * @interface PolkadotTransactionConfig
 */
export interface PolkadotTransactionConfig {
    [key: string]: any;

    /**
     * 
     * @type {PolkadotTransactionConfigTransferSubmittable}
     * @memberof PolkadotTransactionConfig
     */
    'transferSubmittable'?: PolkadotTransactionConfigTransferSubmittable;
    /**
     * 
     * @type {PolkadotTransactionConfigTransferSubmittable}
     * @memberof PolkadotTransactionConfig
     */
    'to'?: PolkadotTransactionConfigTransferSubmittable;
    /**
     * 
     * @type {PolkadotTransactionConfigValue}
     * @memberof PolkadotTransactionConfig
     */
    'value'?: PolkadotTransactionConfigValue;
}
/**
 * @type PolkadotTransactionConfigTransferSubmittable
 * @export
 */
export type PolkadotTransactionConfigTransferSubmittable = string;

/**
 * @type PolkadotTransactionConfigValue
 * @export
 */
export type PolkadotTransactionConfigValue = number;

/**
 * 
 * @export
 * @interface RawTransactionRequest
 */
export interface RawTransactionRequest {
    /**
     * 
     * @type {string}
     * @memberof RawTransactionRequest
     */
    'to': string;
    /**
     * 
     * @type {number}
     * @memberof RawTransactionRequest
     */
    'value': number;
}
/**
 * 
 * @export
 * @interface RawTransactionResponse
 */
export interface RawTransactionResponse {
    /**
     * 
     * @type {RawTransactionResponseResponseContainer}
     * @memberof RawTransactionResponse
     */
    'responseContainer': RawTransactionResponseResponseContainer;
}
/**
 * 
 * @export
 * @interface RawTransactionResponseData
 */
export interface RawTransactionResponseData {
    /**
     * 
     * @type {string}
     * @memberof RawTransactionResponseData
     */
    'rawTransaction': string;
}
/**
 * 
 * @export
 * @interface RawTransactionResponseResponseContainer
 */
export interface RawTransactionResponseResponseContainer {
    /**
     * 
     * @type {RawTransactionResponseData}
     * @memberof RawTransactionResponseResponseContainer
     */
    'response_data': RawTransactionResponseData;
    /**
     * 
     * @type {boolean}
     * @memberof RawTransactionResponseResponseContainer
     */
    'succeeded': boolean;
    /**
     * 
     * @type {string}
     * @memberof RawTransactionResponseResponseContainer
     */
    'message': string;
    /**
     * 
     * @type {string}
     * @memberof RawTransactionResponseResponseContainer
     */
    'error': string | null;
}
/**
 * 
 * @export
 * @interface RunTransactionRequest
 */
export interface RunTransactionRequest {
    /**
     * 
     * @type {Web3SigningCredential}
     * @memberof RunTransactionRequest
     */
    'web3SigningCredential': Web3SigningCredential;
    /**
     * 
     * @type {PolkadotTransactionConfig}
     * @memberof RunTransactionRequest
     */
    'transactionConfig': PolkadotTransactionConfig;
}
/**
 * 
 * @export
 * @interface RunTransactionResponse
 */
export interface RunTransactionResponse {
    /**
     * 
     * @type {boolean}
     * @memberof RunTransactionResponse
     */
    'success': boolean;
    /**
     * 
     * @type {string}
     * @memberof RunTransactionResponse
     */
    'txHash'?: string;
    /**
     * 
     * @type {string}
     * @memberof RunTransactionResponse
     */
    'blockHash'?: string;
}
/**
 * 
 * @export
 * @interface SignRawTransactionRequest
 */
export interface SignRawTransactionRequest {
    /**
     * 
     * @type {string}
     * @memberof SignRawTransactionRequest
     */
    'rawTransaction': string;
    /**
     * 
     * @type {string}
     * @memberof SignRawTransactionRequest
     */
    'mnemonic': string;
    /**
     * 
     * @type {object}
     * @memberof SignRawTransactionRequest
     */
    'signingOptions'?: object;
}
/**
 * 
 * @export
 * @interface SignRawTransactionResponse
 */
export interface SignRawTransactionResponse {
    /**
     * 
     * @type {boolean}
     * @memberof SignRawTransactionResponse
     */
    'success': boolean;
    /**
     * 
     * @type {string}
     * @memberof SignRawTransactionResponse
     */
    'signedTransaction': string;
}
/**
 * 
 * @export
 * @interface TransactionInfoRequest
 */
export interface TransactionInfoRequest {
    /**
     * 
     * @type {string}
     * @memberof TransactionInfoRequest
     */
    'accountAddress': string;
    /**
     * 
     * @type {number}
     * @memberof TransactionInfoRequest
     */
    'transactionExpiration': number | null;
}
/**
 * 
 * @export
 * @interface TransactionInfoResponse
 */
export interface TransactionInfoResponse {
    /**
     * 
     * @type {TransactionInfoResponseResponseContainer}
     * @memberof TransactionInfoResponse
     */
    'responseContainer': TransactionInfoResponseResponseContainer;
}
/**
 * 
 * @export
 * @interface TransactionInfoResponseData
 */
export interface TransactionInfoResponseData {
    /**
     * 
     * @type {object}
     * @memberof TransactionInfoResponseData
     */
    'nonce': object;
    /**
     * 
     * @type {object}
     * @memberof TransactionInfoResponseData
     */
    'blockHash': object;
    /**
     * 
     * @type {object}
     * @memberof TransactionInfoResponseData
     */
    'era': object | null;
}
/**
 * 
 * @export
 * @interface TransactionInfoResponseResponseContainer
 */
export interface TransactionInfoResponseResponseContainer {
    /**
     * 
     * @type {TransactionInfoResponseData}
     * @memberof TransactionInfoResponseResponseContainer
     */
    'response_data': TransactionInfoResponseData;
    /**
     * 
     * @type {boolean}
     * @memberof TransactionInfoResponseResponseContainer
     */
    'succeeded': boolean;
    /**
     * 
     * @type {string}
     * @memberof TransactionInfoResponseResponseContainer
     */
    'message': string;
    /**
     * 
     * @type {string}
     * @memberof TransactionInfoResponseResponseContainer
     */
    'error': string | null;
}
/**
 * @type Web3SigningCredential
 * @export
 */
export type Web3SigningCredential = Web3SigningCredentialCactusKeychainRef | Web3SigningCredentialMnemonicString | Web3SigningCredentialNone;

/**
 * 
 * @export
 * @interface Web3SigningCredentialCactusKeychainRef
 */
export interface Web3SigningCredentialCactusKeychainRef {
    /**
     * 
     * @type {Web3SigningCredentialType}
     * @memberof Web3SigningCredentialCactusKeychainRef
     */
    'type': Web3SigningCredentialType;
    /**
     * The key to use when looking up the the keychain entry holding the secret pointed to by the  keychainEntryKey parameter.
     * @type {string}
     * @memberof Web3SigningCredentialCactusKeychainRef
     */
    'keychainEntryKey': string;
    /**
     * The keychain ID to use when looking up the the keychain plugin instance that will be used to retrieve the secret pointed to by the keychainEntryKey parameter.
     * @type {string}
     * @memberof Web3SigningCredentialCactusKeychainRef
     */
    'keychainId': string;
}


/**
 * 
 * @export
 * @interface Web3SigningCredentialMnemonicString
 */
export interface Web3SigningCredentialMnemonicString {
    /**
     * 
     * @type {Web3SigningCredentialType}
     * @memberof Web3SigningCredentialMnemonicString
     */
    'type': Web3SigningCredentialType;
    /**
     * The Polkadot account\'s seed phrase for signing transaction
     * @type {string}
     * @memberof Web3SigningCredentialMnemonicString
     */
    'mnemonic': string;
}


/**
 * Using this denotes that there is no signing required because the transaction is pre-signed.
 * @export
 * @interface Web3SigningCredentialNone
 */
export interface Web3SigningCredentialNone {
    /**
     * 
     * @type {Web3SigningCredentialType}
     * @memberof Web3SigningCredentialNone
     */
    'type': Web3SigningCredentialType;
}


/**
 * 
 * @export
 * @enum {string}
 */

export const Web3SigningCredentialType = {
    CactusKeychainRef: 'CACTUS_KEYCHAIN_REF',
    MnemonicString: 'MNEMONIC_STRING',
    None: 'NONE'
} as const;

export type Web3SigningCredentialType = typeof Web3SigningCredentialType[keyof typeof Web3SigningCredentialType];



/**
 * DefaultApi - axios parameter creator
 * @export
 */
export const DefaultApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Deploys the ink! contract
         * @param {DeployContractInkRequest} [deployContractInkRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deployContractInk: async (deployContractInkRequest?: DeployContractInkRequest, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-polkadot/deploy-contract-ink`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(deployContractInkRequest, localVarRequestOptions, configuration)

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
        getPrometheusMetrics: async (options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-polkadot/get-prometheus-exporter-metrics`;
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
         * @summary Get raw unsigned transaction
         * @param {RawTransactionRequest} [rawTransactionRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getRawTransaction: async (rawTransactionRequest?: RawTransactionRequest, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-polkadot/get-raw-transaction`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(rawTransactionRequest, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Get the necessary Transaction Info for a account
         * @param {TransactionInfoRequest} [transactionInfoRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getTransactionInfo: async (transactionInfoRequest?: TransactionInfoRequest, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-polkadot/get-transaction-info`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(transactionInfoRequest, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Invokes a contract on a polkadot ledger
         * @param {InvokeContractRequest} [invokeContractRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        invokeContract: async (invokeContractRequest?: InvokeContractRequest, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-polkadot/invoke-contract`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(invokeContractRequest, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Executes a transaction on a Polkadot ledger
         * @param {RunTransactionRequest} [runTransactionRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        runTransaction: async (runTransactionRequest?: RunTransactionRequest, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-polkadot/run-transaction`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(runTransactionRequest, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary sign the raw transaction
         * @param {SignRawTransactionRequest} [signRawTransactionRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        signRawTransaction: async (signRawTransactionRequest?: SignRawTransactionRequest, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-polkadot/sign-raw-transaction`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(signRawTransactionRequest, localVarRequestOptions, configuration)

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
         * @summary Deploys the ink! contract
         * @param {DeployContractInkRequest} [deployContractInkRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async deployContractInk(deployContractInkRequest?: DeployContractInkRequest, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<DeployContractInkResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.deployContractInk(deployContractInkRequest, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Get the Prometheus Metrics
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getPrometheusMetrics(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getPrometheusMetrics(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Get raw unsigned transaction
         * @param {RawTransactionRequest} [rawTransactionRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getRawTransaction(rawTransactionRequest?: RawTransactionRequest, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<RawTransactionResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getRawTransaction(rawTransactionRequest, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Get the necessary Transaction Info for a account
         * @param {TransactionInfoRequest} [transactionInfoRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getTransactionInfo(transactionInfoRequest?: TransactionInfoRequest, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<TransactionInfoResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getTransactionInfo(transactionInfoRequest, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Invokes a contract on a polkadot ledger
         * @param {InvokeContractRequest} [invokeContractRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async invokeContract(invokeContractRequest?: InvokeContractRequest, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<InvokeContractResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.invokeContract(invokeContractRequest, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Executes a transaction on a Polkadot ledger
         * @param {RunTransactionRequest} [runTransactionRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async runTransaction(runTransactionRequest?: RunTransactionRequest, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<RunTransactionResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.runTransaction(runTransactionRequest, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary sign the raw transaction
         * @param {SignRawTransactionRequest} [signRawTransactionRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async signRawTransaction(signRawTransactionRequest?: SignRawTransactionRequest, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<SignRawTransactionResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.signRawTransaction(signRawTransactionRequest, options);
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
         * @summary Deploys the ink! contract
         * @param {DeployContractInkRequest} [deployContractInkRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deployContractInk(deployContractInkRequest?: DeployContractInkRequest, options?: any): AxiosPromise<DeployContractInkResponse> {
            return localVarFp.deployContractInk(deployContractInkRequest, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Get the Prometheus Metrics
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getPrometheusMetrics(options?: any): AxiosPromise<string> {
            return localVarFp.getPrometheusMetrics(options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Get raw unsigned transaction
         * @param {RawTransactionRequest} [rawTransactionRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getRawTransaction(rawTransactionRequest?: RawTransactionRequest, options?: any): AxiosPromise<RawTransactionResponse> {
            return localVarFp.getRawTransaction(rawTransactionRequest, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Get the necessary Transaction Info for a account
         * @param {TransactionInfoRequest} [transactionInfoRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getTransactionInfo(transactionInfoRequest?: TransactionInfoRequest, options?: any): AxiosPromise<TransactionInfoResponse> {
            return localVarFp.getTransactionInfo(transactionInfoRequest, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Invokes a contract on a polkadot ledger
         * @param {InvokeContractRequest} [invokeContractRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        invokeContract(invokeContractRequest?: InvokeContractRequest, options?: any): AxiosPromise<InvokeContractResponse> {
            return localVarFp.invokeContract(invokeContractRequest, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Executes a transaction on a Polkadot ledger
         * @param {RunTransactionRequest} [runTransactionRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        runTransaction(runTransactionRequest?: RunTransactionRequest, options?: any): AxiosPromise<RunTransactionResponse> {
            return localVarFp.runTransaction(runTransactionRequest, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary sign the raw transaction
         * @param {SignRawTransactionRequest} [signRawTransactionRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        signRawTransaction(signRawTransactionRequest?: SignRawTransactionRequest, options?: any): AxiosPromise<SignRawTransactionResponse> {
            return localVarFp.signRawTransaction(signRawTransactionRequest, options).then((request) => request(axios, basePath));
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
     * @summary Deploys the ink! contract
     * @param {DeployContractInkRequest} [deployContractInkRequest] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public deployContractInk(deployContractInkRequest?: DeployContractInkRequest, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).deployContractInk(deployContractInkRequest, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Get the Prometheus Metrics
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getPrometheusMetrics(options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).getPrometheusMetrics(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Get raw unsigned transaction
     * @param {RawTransactionRequest} [rawTransactionRequest] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getRawTransaction(rawTransactionRequest?: RawTransactionRequest, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).getRawTransaction(rawTransactionRequest, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Get the necessary Transaction Info for a account
     * @param {TransactionInfoRequest} [transactionInfoRequest] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getTransactionInfo(transactionInfoRequest?: TransactionInfoRequest, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).getTransactionInfo(transactionInfoRequest, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Invokes a contract on a polkadot ledger
     * @param {InvokeContractRequest} [invokeContractRequest] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public invokeContract(invokeContractRequest?: InvokeContractRequest, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).invokeContract(invokeContractRequest, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Executes a transaction on a Polkadot ledger
     * @param {RunTransactionRequest} [runTransactionRequest] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public runTransaction(runTransactionRequest?: RunTransactionRequest, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).runTransaction(runTransactionRequest, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary sign the raw transaction
     * @param {SignRawTransactionRequest} [signRawTransactionRequest] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public signRawTransaction(signRawTransactionRequest?: SignRawTransactionRequest, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).signRawTransaction(signRawTransactionRequest, options).then((request) => request(this.axios, this.basePath));
    }
}


