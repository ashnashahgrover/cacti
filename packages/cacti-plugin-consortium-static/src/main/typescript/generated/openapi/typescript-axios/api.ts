/* tslint:disable */
/* eslint-disable */
/**
 * Hyperledger Cactus Plugin - Consortium Web Service
 * Manage a Cactus consortium through the APIs. Needs administrative privileges.
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
 * @interface BroadcastRequestV1
 */
export interface BroadcastRequestV1 {
    /**
     * 
     * @type {BroadcastRequestV1Message}
     * @memberof BroadcastRequestV1
     */
    'message': BroadcastRequestV1Message;
    /**
     * 
     * @type {string}
     * @memberof BroadcastRequestV1
     */
    'signature': string;
}
/**
 * 
 * @export
 * @interface BroadcastRequestV1Message
 */
export interface BroadcastRequestV1Message {
    /**
     * 
     * @type {NewNodeRequestV1}
     * @memberof BroadcastRequestV1Message
     */
    'message': NewNodeRequestV1;
    /**
     * 
     * @type {string}
     * @memberof BroadcastRequestV1Message
     */
    'pubKey': string;
}
/**
 * 
 * @export
 * @interface BroadcastResponseV1
 */
export interface BroadcastResponseV1 {
    /**
     * 
     * @type {boolean}
     * @memberof BroadcastResponseV1
     */
    'result': boolean;
}
/**
 * A Cactus node can be a single server, or a set of servers behind a load balancer acting as one.
 * @export
 * @interface CactusNode
 */
export interface CactusNode {
    /**
     * 
     * @type {string}
     * @memberof CactusNode
     */
    'nodeApiHost': string;
    /**
     * The PEM encoded public key that was used to generate the JWS included in the response (the jws property)
     * @type {string}
     * @memberof CactusNode
     */
    'publicKeyPem': string;
    /**
     * 
     * @type {string}
     * @memberof CactusNode
     */
    'id': string;
    /**
     * 
     * @type {string}
     * @memberof CactusNode
     */
    'consortiumId': string;
    /**
     * 
     * @type {string}
     * @memberof CactusNode
     */
    'memberId': string;
    /**
     * Stores an array of Ledger entity IDs that are reachable (routable) via this Cactus Node. This information is used by the client side SDK API client to figure out at runtime where to send API requests that are specific to a certain ledger such as requests to execute transactions.
     * @type {Array<string>}
     * @memberof CactusNode
     */
    'ledgerIds': Array<string>;
    /**
     * 
     * @type {Array<string>}
     * @memberof CactusNode
     */
    'pluginInstanceIds': Array<string>;
}
/**
 * 
 * @export
 * @interface CactusNodeAllOf
 */
export interface CactusNodeAllOf {
    /**
     * 
     * @type {string}
     * @memberof CactusNodeAllOf
     */
    'id': string;
    /**
     * 
     * @type {string}
     * @memberof CactusNodeAllOf
     */
    'consortiumId': string;
    /**
     * 
     * @type {string}
     * @memberof CactusNodeAllOf
     */
    'memberId': string;
    /**
     * Stores an array of Ledger entity IDs that are reachable (routable) via this Cactus Node. This information is used by the client side SDK API client to figure out at runtime where to send API requests that are specific to a certain ledger such as requests to execute transactions.
     * @type {Array<string>}
     * @memberof CactusNodeAllOf
     */
    'ledgerIds': Array<string>;
    /**
     * 
     * @type {Array<string>}
     * @memberof CactusNodeAllOf
     */
    'pluginInstanceIds': Array<string>;
}
/**
 * A Cactus node meta information
 * @export
 * @interface CactusNodeMeta
 */
export interface CactusNodeMeta {
    /**
     * 
     * @type {string}
     * @memberof CactusNodeMeta
     */
    'nodeApiHost': string;
    /**
     * The PEM encoded public key that was used to generate the JWS included in the response (the jws property)
     * @type {string}
     * @memberof CactusNodeMeta
     */
    'publicKeyPem': string;
}
/**
 * 
 * @export
 * @interface GetConsortiumJwsResponse
 */
export interface GetConsortiumJwsResponse {
    /**
     * 
     * @type {JWSGeneral}
     * @memberof GetConsortiumJwsResponse
     */
    'jws': JWSGeneral;
}
/**
 * 
 * @export
 * @interface GetNodeJwsResponse
 */
export interface GetNodeJwsResponse {
    /**
     * 
     * @type {JWSGeneral}
     * @memberof GetNodeJwsResponse
     */
    'jws': JWSGeneral;
}
/**
 * Identity object of a Cacti Node
 * @export
 * @interface Identity
 */
export interface Identity {
    /**
     * Public key of the Node
     * @type {string}
     * @memberof Identity
     */
    'pubKey': string;
    /**
     * memberId of the organization the node belongs to
     * @type {string}
     * @memberof Identity
     */
    'memberId': string;
    /**
     * JWT generated by the organization the node belongs to
     * @type {string}
     * @memberof Identity
     */
    'proof': string;
}
/**
 * 
 * @export
 * @interface JWSGeneral
 */
export interface JWSGeneral {
    /**
     * 
     * @type {string}
     * @memberof JWSGeneral
     */
    'payload': string;
    /**
     * 
     * @type {Array<JWSRecipient>}
     * @memberof JWSGeneral
     */
    'signatures': Array<JWSRecipient>;
}
/**
 * A JSON Web Signature. See: https://tools.ietf.org/html/rfc7515 for info about standard.
 * @export
 * @interface JWSRecipient
 */
export interface JWSRecipient {
    /**
     * 
     * @type {string}
     * @memberof JWSRecipient
     */
    'signature': string;
    /**
     * 
     * @type {string}
     * @memberof JWSRecipient
     */
    'protected'?: string;
    /**
     * 
     * @type {{ [key: string]: any; }}
     * @memberof JWSRecipient
     */
    'header'?: { [key: string]: any; };
}
/**
 * 
 * @export
 * @interface Ledger
 */
export interface Ledger {
    /**
     * 
     * @type {string}
     * @memberof Ledger
     */
    'id': string;
    /**
     * 
     * @type {LedgerType}
     * @memberof Ledger
     */
    'ledgerType': LedgerType;
    /**
     * 
     * @type {string}
     * @memberof Ledger
     */
    'consortiumMemberId'?: string;
}


/**
 * Enumerates the different ledger vendors and their major versions encoded within the name of the LedgerType. For example \"BESU_1X\" involves all of the [1.0.0;2.0.0) where 1.0.0 is included and anything up until, but not 2.0.0. See: https://stackoverflow.com/a/4396303/698470 for further explanation.
 * @export
 * @enum {string}
 */

export const LedgerType = {
    Besu1X: 'BESU_1X',
    Besu2X: 'BESU_2X',
    Burrow0X: 'BURROW_0X',
    Corda4X: 'CORDA_4X',
    Fabric2: 'FABRIC_2',
    Sawtooth1X: 'SAWTOOTH_1X'
} as const;

export type LedgerType = typeof LedgerType[keyof typeof LedgerType];


/**
 * 
 * @export
 * @interface NewNodeRequestV1
 */
export interface NewNodeRequestV1 {
    /**
     * 
     * @type {Identity}
     * @memberof NewNodeRequestV1
     */
    'identity': Identity;
    /**
     * proof of common configs and policies defined by the consortium
     * @type {string}
     * @memberof NewNodeRequestV1
     */
    'treeHash': string;
    /**
     * 
     * @type {CactusNode}
     * @memberof NewNodeRequestV1
     */
    'node': CactusNode;
    /**
     * Ledger metadata
     * @type {Array<Ledger>}
     * @memberof NewNodeRequestV1
     */
    'ledger': Array<Ledger>;
    /**
     * Plugin Instance metadata
     * @type {Array<PluginInstance>}
     * @memberof NewNodeRequestV1
     */
    'pluginInstance': Array<PluginInstance>;
    /**
     * signature of the message
     * @type {string}
     * @memberof NewNodeRequestV1
     */
    'signature': string;
}
/**
 * 
 * @export
 * @interface PluginInstance
 */
export interface PluginInstance {
    /**
     * 
     * @type {string}
     * @memberof PluginInstance
     */
    'id': string;
    /**
     * 
     * @type {string}
     * @memberof PluginInstance
     */
    'packageName': string;
}

/**
 * DefaultApi - axios parameter creator
 * @export
 */
export const DefaultApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Adds a node to consortium JWS
         * @param {NewNodeRequestV1} [newNodeRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        addNodeToConsortiumV1: async (newNodeRequestV1?: NewNodeRequestV1, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cacti-plugin-consortium-static/add-node`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(newNodeRequestV1, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * The JWS asserting the consortium metadata (pub keys and hosts of nodes)
         * @summary Retrieves a consortium JWS
         * @param {object} [body] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getConsortiumJwsV1: async (body?: object, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cacti-plugin-consortium-static/consortium/jws`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(body, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Retrieves the JWT of a Cactus Node
         * @param {object} [body] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getNodeJwsV1: async (body?: object, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cacti-plugin-consortium-static/node/jws`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(body, localVarRequestOptions, configuration)

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
            const localVarPath = `/api/v1/plugins/@hyperledger/cacti-plugin-consortium-static/get-prometheus-exporter-metrics`;
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
         * @summary Adds a node to consortium JWS
         * @param {BroadcastRequestV1} [broadcastRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        receiveBroadcastV1: async (broadcastRequestV1?: BroadcastRequestV1, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cacti-plugin-consortium-static/receive-broadcast`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(broadcastRequestV1, localVarRequestOptions, configuration)

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
         * @summary Adds a node to consortium JWS
         * @param {NewNodeRequestV1} [newNodeRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async addNodeToConsortiumV1(newNodeRequestV1?: NewNodeRequestV1, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<GetConsortiumJwsResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.addNodeToConsortiumV1(newNodeRequestV1, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * The JWS asserting the consortium metadata (pub keys and hosts of nodes)
         * @summary Retrieves a consortium JWS
         * @param {object} [body] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getConsortiumJwsV1(body?: object, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<GetConsortiumJwsResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getConsortiumJwsV1(body, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Retrieves the JWT of a Cactus Node
         * @param {object} [body] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getNodeJwsV1(body?: object, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<GetNodeJwsResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getNodeJwsV1(body, options);
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
         * @summary Adds a node to consortium JWS
         * @param {BroadcastRequestV1} [broadcastRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async receiveBroadcastV1(broadcastRequestV1?: BroadcastRequestV1, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<BroadcastResponseV1>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.receiveBroadcastV1(broadcastRequestV1, options);
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
         * @summary Adds a node to consortium JWS
         * @param {NewNodeRequestV1} [newNodeRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        addNodeToConsortiumV1(newNodeRequestV1?: NewNodeRequestV1, options?: any): AxiosPromise<GetConsortiumJwsResponse> {
            return localVarFp.addNodeToConsortiumV1(newNodeRequestV1, options).then((request) => request(axios, basePath));
        },
        /**
         * The JWS asserting the consortium metadata (pub keys and hosts of nodes)
         * @summary Retrieves a consortium JWS
         * @param {object} [body] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getConsortiumJwsV1(body?: object, options?: any): AxiosPromise<GetConsortiumJwsResponse> {
            return localVarFp.getConsortiumJwsV1(body, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Retrieves the JWT of a Cactus Node
         * @param {object} [body] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getNodeJwsV1(body?: object, options?: any): AxiosPromise<GetNodeJwsResponse> {
            return localVarFp.getNodeJwsV1(body, options).then((request) => request(axios, basePath));
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
         * @summary Adds a node to consortium JWS
         * @param {BroadcastRequestV1} [broadcastRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        receiveBroadcastV1(broadcastRequestV1?: BroadcastRequestV1, options?: any): AxiosPromise<BroadcastResponseV1> {
            return localVarFp.receiveBroadcastV1(broadcastRequestV1, options).then((request) => request(axios, basePath));
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
     * @summary Adds a node to consortium JWS
     * @param {NewNodeRequestV1} [newNodeRequestV1] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public addNodeToConsortiumV1(newNodeRequestV1?: NewNodeRequestV1, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).addNodeToConsortiumV1(newNodeRequestV1, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * The JWS asserting the consortium metadata (pub keys and hosts of nodes)
     * @summary Retrieves a consortium JWS
     * @param {object} [body] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getConsortiumJwsV1(body?: object, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).getConsortiumJwsV1(body, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Retrieves the JWT of a Cactus Node
     * @param {object} [body] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getNodeJwsV1(body?: object, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).getNodeJwsV1(body, options).then((request) => request(this.axios, this.basePath));
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
     * @summary Adds a node to consortium JWS
     * @param {BroadcastRequestV1} [broadcastRequestV1] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public receiveBroadcastV1(broadcastRequestV1?: BroadcastRequestV1, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).receiveBroadcastV1(broadcastRequestV1, options).then((request) => request(this.axios, this.basePath));
    }
}


