/**
 *
 * Please note:
 * This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * Do not edit this file manually.
 *
 */

@file:Suppress(
    "ArrayInDataClass",
    "EnumEntryName",
    "RemoveRedundantQualifierName",
    "UnusedImport"
)

package org.openapitools.client.models

import org.openapitools.client.models.GatewayOptions
import org.openapitools.client.models.GetBlockRequestV1Query
import org.openapitools.client.models.GetBlockResponseTypeV1

import com.squareup.moshi.Json
import com.squareup.moshi.JsonClass

/**
 * Request for GetBlock endpoint.
 *
 * @param channelName Fabric channel which we want to query.
 * @param gatewayOptions 
 * @param query 
 * @param connectionChannelName Fabric channel we want to connect to. If not provided, then one from channelName parameter will be used
 * @param responseType 
 */


data class GetBlockRequestV1 (

    /* Fabric channel which we want to query. */
    @Json(name = "channelName")
    val channelName: kotlin.String,

    @Json(name = "gatewayOptions")
    val gatewayOptions: GatewayOptions,

    @Json(name = "query")
    val query: GetBlockRequestV1Query,

    /* Fabric channel we want to connect to. If not provided, then one from channelName parameter will be used */
    @Json(name = "connectionChannelName")
    val connectionChannelName: kotlin.String? = null,

    @Json(name = "responseType")
    val responseType: GetBlockResponseTypeV1? = GetBlockResponseTypeV1.Full

)
