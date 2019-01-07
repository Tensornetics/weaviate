/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN: Bob van Luijt (bob@k10y.co)
 */

// Package descriptions provides the descriptions as used by the graphql endpoint for Weaviate
package descriptions

import ()

// AGGREGATE
const LocalAggregateThingsDesc string = "Aggregate Things on a local Weaviate"
const LocalAggregateActionsDesc string = "Aggregate Things on a local Weaviate"

const GroupByDesc string = "Specify which properties to group by"

const LocalAggregateObjDesc string = "An object allowing Aggregation of Things and Actions"

const LocalAggregateThingsActionsObjDesc string = "An object allowing Aggregation of %ss on a local Weaviate"

const LocalAggregateMeanDesc string = "Aggregate on the mean of numeric property values"
const LocalAggregateSumDesc string = "Aggregate on the sum of numeric property values"
const LocalAggregateMedianDesc string = "Aggregate on the median of numeric property values"
const LocalAggregateModeDesc string = "Aggregate on the mode of numeric property values"
const LocalAggregateMinDesc string = "Aggregate on the minimum of numeric property values"
const LocalAggregateMaxDesc string = "Aggregate on the maximum of numeric property values"
const LocalAggregateCountDesc string = "Aggregate on the total amount of found property values"
const LocalAggregateGroupedByDesc string = "Indicates the group of returned data"

const LocalAggregateNumericObj string = "An object containing the %s of numeric properties"

const LocalAggregateCountObj string = "An object containing countable properties"

const LocalAggregateGroupedByObjDesc string = "An object containing the path and value of the grouped property"

const LocalAggregateGroupedByGroupedByPathDesc string = "The path of the grouped property"
const LocalAggregateGroupedByGroupedByValueDesc string = "The value of the grouped property"

// NETWORK
const NetworkAggregateWeaviateObjDesc string = "An object containing Get Things and Actions fields for network Weaviate instance: "

const NetworkAggregateDesc string = "Perform Aggregation of Things and Actions"

const NetworkAggregateThingsDesc string = "Aggregate Things on a network Weaviate"
const NetworkAggregateActionsDesc string = "Aggregate Things on a network Weaviate"

const NetworkAggregateObjDesc string = "An object allowing Aggregation of Things and Actions"

const NetworkAggregateThingsActionsObjDesc string = "An object allowing Aggregation of %ss on a network Weaviate"

const NetworkAggregateMeanDesc string = "Aggregate on the mean of numeric property values"
const NetworkAggregateSumDesc string = "Aggregate on the sum of numeric property values"
const NetworkAggregateMedianDesc string = "Aggregate on the median of numeric property values"
const NetworkAggregateModeDesc string = "Aggregate on the mode of numeric property values"
const NetworkAggregateMinDesc string = "Aggregate on the minimum of numeric property values"
const NetworkAggregateMaxDesc string = "Aggregate on the maximum of numeric property values"
const NetworkAggregateCountDesc string = "Aggregate on the total amount of found property values"
const NetworkAggregateGroupedByDesc string = "Indicates the group of returned data"

const NetworkAggregateNumericObj string = "An object containing the %s of numeric properties"

const NetworkAggregateCountObj string = "An object containing countable properties"

const NetworkAggregateGroupedByObjDesc string = "An object containing the path and value of the grouped property"

const NetworkAggregateGroupedByGroupedByPathDesc string = "The path of the grouped property"
const NetworkAggregateGroupedByGroupedByValueDesc string = "The value of the grouped property"