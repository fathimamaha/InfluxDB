Welcome to the InfluxDB wiki!

# Introduction to Time Series
***

### Time Series Data
Time series data is a collection of data from the same source collected over the some interval of time, with x-axis always being the time variable.

E.g., A graph measuring temperature in time intervals, a heat map showing its findings over time.

### Types of Time-Series Data
1. Regular time intervals (can convert by summarizing irregular findings over a time interval)
2. Irregular time intervals

### Time-Series Database
Huge volume, time-stamped data, real-time, sensitive, growing market

Use cases like IoT, Monitoring solutions and Real-time Analytics use time-series data extensively.

### InfluxData
It can quickly ingest data from everywhere and efficiently store it.
It can automate what to do with that data and control the functions that are carried out on it.
You can observe and visualise the data, can see the change over time.
You can also facilitate Machine learning algorithms to this data over time or detect anomaly in the data.

1. Instrument
2. Automate
3. Observe
4. Learn

### The TICK Stack
InfluxDB- OpenSource
>* Telegraf: Agent for collecting and reporting all the Metrics (more than 200 plugins)
>* Chronograf: the UI for InfluxData
>* Kapacitor: Processing Engine

Note the architecture diagram*

### InfluxDB Data Model
1. Measurement- what you want to measure
2. Tickers- Metadata for that measurement
3. Field- The value we are collecting
4. Time Stamps

In a text format
> Measurement(Bucket),ticker=<name>,field=<value> time stamp(can be even nanoseconds)
> measurement,tagset fielset timestamp

### Specific UseCases
>BBOXX | AWS |
>WAYFAIR |
>PLAYTECH

***

# Introduction to InfluxDB & Telegraf
***

### Telegraf
It is a plugin driven agent for collecting and reporting various metrics.

Different plugins (easy to write):
* Input plugins
* Output plugins
* Aggregator plugins
* Processor plugins

### Installation and Data Ingestion
Installing telegraf, chronograf and influxdb
Go to http://localhost:8086 (or you can check)

>Measurement,Tagset Field set timestamp
>cpu_load,hostname=server02,az=ind temp=24.5,volts=7 123234200000

### Considerations

1. Tags are indexed
2. Fields are not indexed (integers)
3. Field types immutable
4. GROUP BY only tags
5. Downsampling: high precision for a limited time and lower precision summarized data for a long time. To mitigate storage concerns.

### TELEGRAF

Can act as agent, collector, ingest pipeline
Minimal memory footprint
Tagging of metrics

Some plugins:
> * histogram
> * nginx
> * kubernetes

***
# Introduction to Telegraf and Plugins Ecosystem
***

### Line protocol
temperature,pin=620015,country=India value=45,humidity=10 1388432000000

Schema:
Mathematical operations only on fields.

> * SELECT * FROM _
> * SELECT x+y FROM India

Telegraf- Written in GO for collecting metrics
Multiple telegraf agents can push to the same bucket in influxdb


### Plugins
plugins very easy to support, to expand, etc.

Input Plugins: cpu,mem,diskio
actively searching for data

Output Plugins: Influxdv,graphite
flush data into influx

Service Plugins: TCP, UDP
extension of input plugins

### Command

config (configuration file)
check out a sample configuration
[agent]
[input]
[output]

example:
using Grahite input stream
Graphite => Line Protocol => Reconstructed Graphite

***

***
# Chronograf and Dashboarding
***

Go to local host
Boards>Dashboards>Buckets>View dashboards

Visualise max, min, average, histogram

***
# Introduction to flux
***
Flux is an alternative to InfluxQueryLanguage: to filter and play wiht the time series data on InfluxDB
Flux is a functional data scripting and query language to interface with data
Compose
Read 
Share

example:
from(bucket:"telegraf/autogen")
range(start:-1h)
r.measurement=="CPU" AND
r.field=="usage_system" AND
r.cpu=="cpu_total"

looking for a tag named all these conditions

flux functions:

Input: from()
Output:yield()
Transformation: mean(), count()
systemtime()
Or custom defined or more


***
# Advanced Telegraf topics
***

flush: the interval where you have to output data

Queuing system:
1. RabbitMQ
1. Kafka
1. MQTT
1. NATS
1. NSQ

when agents is true:
> * internal_agents: basic stats across all plugins
> * Internal_write: output plugins aggregated by plugin type
> * Internal_gather: stats for input plugins aggregated by plugin type
> * Internal_memstats: measurement by memory usage

challenges
> * assign a host name
> * environment variables


***
# Chronograf and Dashboarding
***
Objective of Chronograf:
1. Unified Experience
2. Deeply integrated Data Exploration
3. Rapid Time-to-Value

All 4 TICKstack components can be downloaded or individually

Playing around with the dashboards, status

We can add Query like sql or flux.

Also build alert rules or write TICKscripts
will be detailed in next section

***
# Intro to Kapacitor for alerting and Anomaly detection
***

Real-time streaming Data Processing Engine

It can do:
> * Alerting
> * Processing ETL jobs

Processes both strea and batch data from influxDB
It can plug in custom logic or userdefined functions to
-process alerts
-match metric patterns
-compute statistics
-perform specific actions

We can do a batch task or stream task:
Batch: periodically, does not buffer, places additional query load on influxdb
Streaming: writes mirrored to influxDB instance, Bufers data in RAM, additional query load on Kapacitor

example tick script

> > var measurement='requests'

> > > var data=stream
> > > |from () 
> > > .measurement(measurement)
> > > |where(lambda: "is_up" == TRUE)
> > > |where(lambda: "my_field" >10)
> > > |window()
> > > .period(5m)
> > > .every(5m)

> > > //count numberif points in window data
> > > |count('value')
> > > .as('the_average')

### Using Kapacitor for Downsampling

> typically writes back into InfluxDB
process of reducing a sequence of points in a series to a single data point
we can compute the average, min, max, etc of a window of data for particular series.

for faster queries and storing less data


***

***
# Telegraf Input, Output Data Formats
***

### Installing telegraf

* installing through the website
wget https://dl.influxdata.com/telegraf/releases/telegraf_1.17.0-1_amd64.deb
sudo dpkg -i telegraf_1.17.0-1_amd64.deb
* create a config file
* generate token
* call config file

### telegraf keywords

* agent: gathers metrics from the declared input plugins and sends metrics to the declared output plugins, based on the plugins enabled by the given configuration.

* aggregator plugin: receive raw metrics from input plugins and create aggregate metrics from them

* batch size: size of each write batch that Telegraf sends to the output plugins

* collection interval: interval can be overridden by each individual input plugin’s configuration

* collection jitter: each collection interval every input plugin will sleep for a random time between zero and the collection jitter before collecting the metrics

* flush interval

* flush jitter

* input plugin: gather metrics and deliver them to the core agent, where aggregator, processor, and output plugins can operate on the metrics

* metric buffer

* output plugin: deliver metrics to their configured destination

* precision: the precision configuration setting determines how much timestamp precision is retained in the points received from input plugins. All incoming timestamps are truncated to the given precision (ns, us or µs, ms, and s)

precision-ms, the nanosecond timestamp 1480000000123456789 would be truncated to 1480000000123 in millisecond precision

* processor plugin: processor plugins transform, decorate, and/or filter metrics collected by input plugins, passing the transformed metrics to the output plugins

* service input plugin: service input plugins are input plugins that run in a passive collection mode while the Telegraf agent is running

### Telegraf metrics
Measurement name: Description and namespace for the metric.
Tags: Key/Value string pairs and usually used to identify the metric.
Fields: Key/Value pairs that are typed and usually contain the metric data.
Timestamp: Date and time associated with the fields.


***
# Write Data in Influx DB
***

What is needed to write data
* organization
* bucket
* authentication token
* InfluxDB URL

-Use telegraf agent to collect and write data to influxdb
-scrape data using third party softwares

# Query Data using Flux
Flux is used to query with real time data in influxdb, it is used to analyse, processs and filter data in influxdb.

pipe-forward operators (|>): After each function or operation, Flux returns a table or collection of tables containing data, the pipe-forward operator pipes those tables into the next function or operation where they are further processed or manipulated

table properties
* group key

How to query with influxdb?
1. from(bucket:"bucketname")
2. time range specification
  |> range(start: -1h, stop: -10m)
3. filter the data
  |> filter(fn: (r) =>
and add conditions
    r._measurement == "cpu" and
    r._field == "usage_system"
4. yield data using
  |> yield()

flux functions
5. window the data
  |> window(every: 5m)
6. aggregate the data
  |> mean()

covered all the conditional logics and other functions
check all the sample queries

sample flux script to load dat afrom one bucket an dload to another

// Task options
option task = {
    name: "cqinterval15m",
    every: 1h,
    offset: 0m,
    concurrency: 1,
    retry: 5
}

// Data source
data = from(bucket: "telegraf/default")
  |> range(start: -task.every)
  |> filter(fn: (r) =>
    r._measurement == "mem" and
    r.host == "myHost"
  )

data
  // Data transformation
  |> aggregateWindow(
    every: 5m,
    fn: mean
  )
  // Data destination
  |> to(bucket: "telegraf_downsampled")


***
# Backup and Restoring data
***
Backing up: to save it on local file system
#Syntax
influx backup <backup-path> -t <root-token>

Restore: 
#Syntax
influxd restore <path-to-backup-directory>


***

***
# Manage InfluxDB users
***

Steps
1. enable authentication in conf file
2. restart influx  service
3. create admin user
~# curl -XPOST "http://localhost:8086/query" --data-urlencode "q=CREATE USER chronothan WITH PASSWORD 'supersecret' WITH ALL PRIVILEGES"

view users on chronograf
1. read access
2. write data
3. grant all access

edit different accesses with different


***
# Restore a Chronograf database
***
steps 
1. locate chronograf file
2. stop chronograf process
3. replace current database with new server
Remove the current database
rm chronograf-v1.db

Replace it with the desired backup file
cp backup/chronograf-v1.db.1.4.4.2 chronograf-v1.db
4. start chronograf


***
# Import and export Chronograf dashboards
***

Export dashboard on chronograf 
Import dashboard on chronograf

#keep in mind the permissions

***
# Visualization types in Chronograf
***

Line Graph
Stacked Graph
Step-Plot Graph
Single Stat
Line Graph + Single Stat
Bar Graph
Gauge
Table
Note


***


***
# Create Chronograf alert rules
***

1. create a kapacitor connection
2. click the configuration (wrench) icon in the sidebar menu, then select Add Config in the Active Kapacitator column
3. Kapacitor URL field, enter the hostname or IP of the machine that Kapacitor is running on. Be sure to include Kapacitor’s default port: 9092
4. click connect

5. Alerting>manage tasks
* name the alert rule
* alert type
Threshold
Alert if data crosses a boundary.

Relative
Alert if data changes relative to data in a different time range.

Deadman
Alert if InfluxDB receives no relevant data for a specified time duration.

* select time-series data, measurement, fields, tags

6. put the condition
7. slack alert handler, push the alert destination to a slack channel
8. make some alerts, view it and view it in alert history

***
# Use dashboard template variable
***

interact with dashboard and make queries

* SELECT :variable_name: FROM "telegraf"."autogen".:measurement: WHERE time < :dashboardTime:
* SELECT "usage_system" AS "System CPU Usage"
* FROM "telegraf".."cpu"
* WHERE time > :dashboardTime:

Template variables:
Databases
Measurements
Field Keys
Tag Keys
Tag Values
CSV
Map
Custom Meta Query
Text

***







