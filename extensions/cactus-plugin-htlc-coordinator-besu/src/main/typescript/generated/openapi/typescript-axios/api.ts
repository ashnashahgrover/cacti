/* tslint:disable */
/* eslint-disable */
/**
 * Hyperledger Cactus Plugin - HTLC Coordinator
 * Can exchange assets between networks
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
 * @interface CounterpartyHTLCRequest
 */
export interface CounterpartyHTLCRequest {
    /**
     * 
     * @type {HtlcPackage}
     * @memberof CounterpartyHTLCRequest
     */
    'htlcPackage': HtlcPackage;
    /**
     * connector Instance Id for the connector plugin
     * @type {string}
     * @memberof CounterpartyHTLCRequest
     */
    'connectorInstanceId': string;
    /**
     * keychainId for the keychain plugin
     * @type {string}
     * @memberof CounterpartyHTLCRequest
     */
    'keychainId': string;
    /**
     * Id for the HTLC
     * @type {string}
     * @memberof CounterpartyHTLCRequest
     */
    'htlcId': string;
    /**
     * 
     * @type {Web3SigningCredential}
     * @memberof CounterpartyHTLCRequest
     */
    'web3SigningCredential': Web3SigningCredential;
    /**
     * 
     * @type {number}
     * @memberof CounterpartyHTLCRequest
     */
    'gas'?: number;
}


/**
 * 
 * @export
 * @enum {string}
 */

export const HtlcPackage = {
    Besu: 'BESU',
    BesuErc20: 'BESU_ERC20'
} as const;

export type HtlcPackage = typeof HtlcPackage[keyof typeof HtlcPackage];


/**
 * 
 * @export
 * @interface OwnHTLCRequest
 */
export interface OwnHTLCRequest {
    /**
     * 
     * @type {HtlcPackage}
     * @memberof OwnHTLCRequest
     */
    'htlcPackage': HtlcPackage;
    /**
     * connector Instance Id for the connector plugin
     * @type {string}
     * @memberof OwnHTLCRequest
     */
    'connectorInstanceId': string;
    /**
     * keychainId for the keychain plugin
     * @type {string}
     * @memberof OwnHTLCRequest
     */
    'keychainId': string;
    /**
     * 
     * @type {Array<any>}
     * @memberof OwnHTLCRequest
     */
    'constructorArgs': Array<any>;
    /**
     * 
     * @type {Web3SigningCredential}
     * @memberof OwnHTLCRequest
     */
    'web3SigningCredential': Web3SigningCredential;
    /**
     * Input amount to lock
     * @type {number}
     * @memberof OwnHTLCRequest
     */
    'inputAmount': number;
    /**
     * Output amount to lock
     * @type {number}
     * @memberof OwnHTLCRequest
     */
    'outputAmount': number;
    /**
     * Timestamp to expire the contract
     * @type {number}
     * @memberof OwnHTLCRequest
     */
    'expiration': number;
    /**
     * Hashlock needed to refund the amount
     * @type {string}
     * @memberof OwnHTLCRequest
     */
    'hashLock': string;
    /**
     * The token address
     * @type {string}
     * @memberof OwnHTLCRequest
     */
    'tokenAddress': string;
    /**
     * The receiver address
     * @type {string}
     * @memberof OwnHTLCRequest
     */
    'receiver': string;
    /**
     * The output network id
     * @type {string}
     * @memberof OwnHTLCRequest
     */
    'outputNetwork': string;
    /**
     * The output addreess to receive the tokens
     * @type {string}
     * @memberof OwnHTLCRequest
     */
    'outputAddress': string;
    /**
     * 
     * @type {number}
     * @memberof OwnHTLCRequest
     */
    'gas'?: number;
}


/**
 * @type Web3SigningCredential
 * @export
 */
export type Web3SigningCredential = Web3SigningCredentialCactusKeychainRef | Web3SigningCredentialNone | Web3SigningCredentialPrivateKeyHex;

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
     * The ethereum account (public key) that the credential  belongs to. Basically the username in the traditional  terminology of authentication.
     * @type {string}
     * @memberof Web3SigningCredentialCactusKeychainRef
     */
    'ethAccount': string;
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
 * @interface Web3SigningCredentialPrivateKeyHex
 */
export interface Web3SigningCredentialPrivateKeyHex {
    /**
     * 
     * @type {Web3SigningCredentialType}
     * @memberof Web3SigningCredentialPrivateKeyHex
     */
    'type': Web3SigningCredentialType;
    /**
     * The ethereum account (public key) that the credential belongs to. Basically the username in the traditional terminology of authentication.
     * @type {string}
     * @memberof Web3SigningCredentialPrivateKeyHex
     */
    'ethAccount': string;
    /**
     * The HEX encoded private key of an eth account.
     * @type {string}
     * @memberof Web3SigningCredentialPrivateKeyHex
     */
    'secret': string;
}


/**
 * 
 * @export
 * @enum {string}
 */

export const Web3SigningCredentialType = {
    CactusKeychainRef: 'CACTUS_KEYCHAIN_REF',
    GethKeychainPassword: 'GETH_KEYCHAIN_PASSWORD',
    PrivateKeyHex: 'PRIVATE_KEY_HEX',
    None: 'NONE'
} as const;

export type Web3SigningCredentialType = typeof Web3SigningCredentialType[keyof typeof Web3SigningCredentialType];


/**
 * 
 * @export
 * @interface WithdrawCounterpartyRequest
 */
export interface WithdrawCounterpartyRequest {
    /**
     * 
     * @type {HtlcPackage}
     * @memberof WithdrawCounterpartyRequest
     */
    'htlcPackage': HtlcPackage;
    /**
     * connector Instance Id for the connector plugin
     * @type {string}
     * @memberof WithdrawCounterpartyRequest
     */
    'connectorInstanceId': string;
    /**
     * keychainId for the keychain plugin
     * @type {string}
     * @memberof WithdrawCounterpartyRequest
     */
    'keychainId': string;
    /**
     * contractId for the contract
     * @type {string}
     * @memberof WithdrawCounterpartyRequest
     */
    'contractId'?: string;
    /**
     * 
     * @type {Web3SigningCredential}
     * @memberof WithdrawCounterpartyRequest
     */
    'web3SigningCredential': Web3SigningCredential;
    /**
     * Id for the HTLC
     * @type {string}
     * @memberof WithdrawCounterpartyRequest
     */
    'htlcId': string;
    /**
     * Counterparty HTLC secret
     * @type {string}
     * @memberof WithdrawCounterpartyRequest
     */
    'secret': string;
    /**
     * 
     * @type {number}
     * @memberof WithdrawCounterpartyRequest
     */
    'gas'?: number;
}



/**
 * DefaultApi - axios parameter creator
 * @export
 */
export const DefaultApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Create an instance to interact with the counterparty HTLC
         * @param {CounterpartyHTLCRequest} [counterpartyHTLCRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        counterpartyHtlcV1: async (counterpartyHTLCRequest?: CounterpartyHTLCRequest, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-htlc-coordinator-besu/counterparty-htlc`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(counterpartyHTLCRequest, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Create an instance to interact with the own HTLC.
         * @param {OwnHTLCRequest} [ownHTLCRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        ownHtlcV1: async (ownHTLCRequest?: OwnHTLCRequest, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-htlc-coordinator-besu/own-htlc`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(ownHTLCRequest, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Withdraw funds of the counterparty HTLC
         * @param {WithdrawCounterpartyRequest} [withdrawCounterpartyRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        withdrawCounterpartyV1: async (withdrawCounterpartyRequest?: WithdrawCounterpartyRequest, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-htlc-coordinator-besu/withdraw-counterparty`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(withdrawCounterpartyRequest, localVarRequestOptions, configuration)

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
         * @summary Create an instance to interact with the counterparty HTLC
         * @param {CounterpartyHTLCRequest} [counterpartyHTLCRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async counterpartyHtlcV1(counterpartyHTLCRequest?: CounterpartyHTLCRequest, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<any>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.counterpartyHtlcV1(counterpartyHTLCRequest, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Create an instance to interact with the own HTLC.
         * @param {OwnHTLCRequest} [ownHTLCRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async ownHtlcV1(ownHTLCRequest?: OwnHTLCRequest, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<any>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.ownHtlcV1(ownHTLCRequest, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Withdraw funds of the counterparty HTLC
         * @param {WithdrawCounterpartyRequest} [withdrawCounterpartyRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async withdrawCounterpartyV1(withdrawCounterpartyRequest?: WithdrawCounterpartyRequest, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<any>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.withdrawCounterpartyV1(withdrawCounterpartyRequest, options);
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
         * @summary Create an instance to interact with the counterparty HTLC
         * @param {CounterpartyHTLCRequest} [counterpartyHTLCRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        counterpartyHtlcV1(counterpartyHTLCRequest?: CounterpartyHTLCRequest, options?: any): AxiosPromise<any> {
            return localVarFp.counterpartyHtlcV1(counterpartyHTLCRequest, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Create an instance to interact with the own HTLC.
         * @param {OwnHTLCRequest} [ownHTLCRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        ownHtlcV1(ownHTLCRequest?: OwnHTLCRequest, options?: any): AxiosPromise<any> {
            return localVarFp.ownHtlcV1(ownHTLCRequest, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Withdraw funds of the counterparty HTLC
         * @param {WithdrawCounterpartyRequest} [withdrawCounterpartyRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        withdrawCounterpartyV1(withdrawCounterpartyRequest?: WithdrawCounterpartyRequest, options?: any): AxiosPromise<any> {
            return localVarFp.withdrawCounterpartyV1(withdrawCounterpartyRequest, options).then((request) => request(axios, basePath));
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
     * @summary Create an instance to interact with the counterparty HTLC
     * @param {CounterpartyHTLCRequest} [counterpartyHTLCRequest] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public counterpartyHtlcV1(counterpartyHTLCRequest?: CounterpartyHTLCRequest, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).counterpartyHtlcV1(counterpartyHTLCRequest, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Create an instance to interact with the own HTLC.
     * @param {OwnHTLCRequest} [ownHTLCRequest] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public ownHtlcV1(ownHTLCRequest?: OwnHTLCRequest, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).ownHtlcV1(ownHTLCRequest, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Withdraw funds of the counterparty HTLC
     * @param {WithdrawCounterpartyRequest} [withdrawCounterpartyRequest] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public withdrawCounterpartyV1(withdrawCounterpartyRequest?: WithdrawCounterpartyRequest, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).withdrawCounterpartyV1(withdrawCounterpartyRequest, options).then((request) => request(this.axios, this.basePath));
    }
}


