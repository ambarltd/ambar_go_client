# CreateDataSourceRequestDataSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | A case sensitive string for the user Ambar should use to connect to your Postgres database. | 
**Password** | **string** | A case sensitive string for the user Ambar should use to connect to your Postgres database. | 
**Hostname** | **string** | A case insensitive string for the host on which your Postgres database is running and which Ambar can use to reach the database. | 
**HostPort** | **string** | The port number passed as a string which Ambar can use to connect to your Postgres database instance. | 
**DatabaseName** | **string** | The case sensitive string name of the database on your database host. | 
**TableName** | **string** | The case sensitive string name of the table the DataSource should read. | 
**GloballyUniqueColumnName** | **string** | The name of a column where the value for any record is globally unique. | 
**IncrementingColumnName** | **string** | The name of a column which monotonically increases on database writes. | 
**PartitioningColumnName** | **string** | A case sensitive string for the name of the column in your table Ambar can partition on.  Note that partition keys must be unique to a given sequence of records. | 
**AdditionalColumns** | **string** | A case sensitive, comma separated list string of additional columns to be read from  the table into Ambar. | 
**BinLogReplicationServerId** | **string** | The server_id value used when starting the MySQL server. See MySQL docs for more information about this value and its defaults when not configured. | 
**TlsTerminationOverrideHost** | **string** | The hostname of the server responsible for terminating TLS connections for your server,  for example if your server is behind a load balancer. | 
**PublicationName** | **string** | The named publication Ambar should use to indicate to Postgres what tables Ambar will be replicating from.  The publication name must be distinct from the name of any existing publication in the current database. See Postgres documentation for your specific version of Postgres for more information on Postgres publications. | 
**ColumnNames** | **string** | A case sensitive, comma separated list string of columns which Ambar should read from the database. The ordering should be consistent with the table from which Ambar will read. | 
**SequenceNumColumn** | **string** | A case sensitive string name of the column which is the incrementing sequence number for records on writes. | 
**SequenceIdColumn** | **string** | A case sensitive string name of the column which is the ID which Ambar will partition message sequences by. | 

## Methods

### NewCreateDataSourceRequestDataSourceConfig

`func NewCreateDataSourceRequestDataSourceConfig(username string, password string, hostname string, hostPort string, databaseName string, tableName string, globallyUniqueColumnName string, incrementingColumnName string, partitioningColumnName string, additionalColumns string, binLogReplicationServerId string, tlsTerminationOverrideHost string, publicationName string, columnNames string, sequenceNumColumn string, sequenceIdColumn string, ) *CreateDataSourceRequestDataSourceConfig`

NewCreateDataSourceRequestDataSourceConfig instantiates a new CreateDataSourceRequestDataSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDataSourceRequestDataSourceConfigWithDefaults

`func NewCreateDataSourceRequestDataSourceConfigWithDefaults() *CreateDataSourceRequestDataSourceConfig`

NewCreateDataSourceRequestDataSourceConfigWithDefaults instantiates a new CreateDataSourceRequestDataSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CreateDataSourceRequestDataSourceConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateDataSourceRequestDataSourceConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateDataSourceRequestDataSourceConfig) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CreateDataSourceRequestDataSourceConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateDataSourceRequestDataSourceConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateDataSourceRequestDataSourceConfig) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetHostname

`func (o *CreateDataSourceRequestDataSourceConfig) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CreateDataSourceRequestDataSourceConfig) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CreateDataSourceRequestDataSourceConfig) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetHostPort

`func (o *CreateDataSourceRequestDataSourceConfig) GetHostPort() string`

GetHostPort returns the HostPort field if non-nil, zero value otherwise.

### GetHostPortOk

`func (o *CreateDataSourceRequestDataSourceConfig) GetHostPortOk() (*string, bool)`

GetHostPortOk returns a tuple with the HostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPort

`func (o *CreateDataSourceRequestDataSourceConfig) SetHostPort(v string)`

SetHostPort sets HostPort field to given value.


### GetDatabaseName

`func (o *CreateDataSourceRequestDataSourceConfig) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *CreateDataSourceRequestDataSourceConfig) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *CreateDataSourceRequestDataSourceConfig) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetTableName

`func (o *CreateDataSourceRequestDataSourceConfig) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *CreateDataSourceRequestDataSourceConfig) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *CreateDataSourceRequestDataSourceConfig) SetTableName(v string)`

SetTableName sets TableName field to given value.


### GetGloballyUniqueColumnName

`func (o *CreateDataSourceRequestDataSourceConfig) GetGloballyUniqueColumnName() string`

GetGloballyUniqueColumnName returns the GloballyUniqueColumnName field if non-nil, zero value otherwise.

### GetGloballyUniqueColumnNameOk

`func (o *CreateDataSourceRequestDataSourceConfig) GetGloballyUniqueColumnNameOk() (*string, bool)`

GetGloballyUniqueColumnNameOk returns a tuple with the GloballyUniqueColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGloballyUniqueColumnName

`func (o *CreateDataSourceRequestDataSourceConfig) SetGloballyUniqueColumnName(v string)`

SetGloballyUniqueColumnName sets GloballyUniqueColumnName field to given value.


### GetIncrementingColumnName

`func (o *CreateDataSourceRequestDataSourceConfig) GetIncrementingColumnName() string`

GetIncrementingColumnName returns the IncrementingColumnName field if non-nil, zero value otherwise.

### GetIncrementingColumnNameOk

`func (o *CreateDataSourceRequestDataSourceConfig) GetIncrementingColumnNameOk() (*string, bool)`

GetIncrementingColumnNameOk returns a tuple with the IncrementingColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementingColumnName

`func (o *CreateDataSourceRequestDataSourceConfig) SetIncrementingColumnName(v string)`

SetIncrementingColumnName sets IncrementingColumnName field to given value.


### GetPartitioningColumnName

`func (o *CreateDataSourceRequestDataSourceConfig) GetPartitioningColumnName() string`

GetPartitioningColumnName returns the PartitioningColumnName field if non-nil, zero value otherwise.

### GetPartitioningColumnNameOk

`func (o *CreateDataSourceRequestDataSourceConfig) GetPartitioningColumnNameOk() (*string, bool)`

GetPartitioningColumnNameOk returns a tuple with the PartitioningColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitioningColumnName

`func (o *CreateDataSourceRequestDataSourceConfig) SetPartitioningColumnName(v string)`

SetPartitioningColumnName sets PartitioningColumnName field to given value.


### GetAdditionalColumns

`func (o *CreateDataSourceRequestDataSourceConfig) GetAdditionalColumns() string`

GetAdditionalColumns returns the AdditionalColumns field if non-nil, zero value otherwise.

### GetAdditionalColumnsOk

`func (o *CreateDataSourceRequestDataSourceConfig) GetAdditionalColumnsOk() (*string, bool)`

GetAdditionalColumnsOk returns a tuple with the AdditionalColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalColumns

`func (o *CreateDataSourceRequestDataSourceConfig) SetAdditionalColumns(v string)`

SetAdditionalColumns sets AdditionalColumns field to given value.


### GetBinLogReplicationServerId

`func (o *CreateDataSourceRequestDataSourceConfig) GetBinLogReplicationServerId() string`

GetBinLogReplicationServerId returns the BinLogReplicationServerId field if non-nil, zero value otherwise.

### GetBinLogReplicationServerIdOk

`func (o *CreateDataSourceRequestDataSourceConfig) GetBinLogReplicationServerIdOk() (*string, bool)`

GetBinLogReplicationServerIdOk returns a tuple with the BinLogReplicationServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinLogReplicationServerId

`func (o *CreateDataSourceRequestDataSourceConfig) SetBinLogReplicationServerId(v string)`

SetBinLogReplicationServerId sets BinLogReplicationServerId field to given value.


### GetTlsTerminationOverrideHost

`func (o *CreateDataSourceRequestDataSourceConfig) GetTlsTerminationOverrideHost() string`

GetTlsTerminationOverrideHost returns the TlsTerminationOverrideHost field if non-nil, zero value otherwise.

### GetTlsTerminationOverrideHostOk

`func (o *CreateDataSourceRequestDataSourceConfig) GetTlsTerminationOverrideHostOk() (*string, bool)`

GetTlsTerminationOverrideHostOk returns a tuple with the TlsTerminationOverrideHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsTerminationOverrideHost

`func (o *CreateDataSourceRequestDataSourceConfig) SetTlsTerminationOverrideHost(v string)`

SetTlsTerminationOverrideHost sets TlsTerminationOverrideHost field to given value.


### GetPublicationName

`func (o *CreateDataSourceRequestDataSourceConfig) GetPublicationName() string`

GetPublicationName returns the PublicationName field if non-nil, zero value otherwise.

### GetPublicationNameOk

`func (o *CreateDataSourceRequestDataSourceConfig) GetPublicationNameOk() (*string, bool)`

GetPublicationNameOk returns a tuple with the PublicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationName

`func (o *CreateDataSourceRequestDataSourceConfig) SetPublicationName(v string)`

SetPublicationName sets PublicationName field to given value.


### GetColumnNames

`func (o *CreateDataSourceRequestDataSourceConfig) GetColumnNames() string`

GetColumnNames returns the ColumnNames field if non-nil, zero value otherwise.

### GetColumnNamesOk

`func (o *CreateDataSourceRequestDataSourceConfig) GetColumnNamesOk() (*string, bool)`

GetColumnNamesOk returns a tuple with the ColumnNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnNames

`func (o *CreateDataSourceRequestDataSourceConfig) SetColumnNames(v string)`

SetColumnNames sets ColumnNames field to given value.


### GetSequenceNumColumn

`func (o *CreateDataSourceRequestDataSourceConfig) GetSequenceNumColumn() string`

GetSequenceNumColumn returns the SequenceNumColumn field if non-nil, zero value otherwise.

### GetSequenceNumColumnOk

`func (o *CreateDataSourceRequestDataSourceConfig) GetSequenceNumColumnOk() (*string, bool)`

GetSequenceNumColumnOk returns a tuple with the SequenceNumColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumColumn

`func (o *CreateDataSourceRequestDataSourceConfig) SetSequenceNumColumn(v string)`

SetSequenceNumColumn sets SequenceNumColumn field to given value.


### GetSequenceIdColumn

`func (o *CreateDataSourceRequestDataSourceConfig) GetSequenceIdColumn() string`

GetSequenceIdColumn returns the SequenceIdColumn field if non-nil, zero value otherwise.

### GetSequenceIdColumnOk

`func (o *CreateDataSourceRequestDataSourceConfig) GetSequenceIdColumnOk() (*string, bool)`

GetSequenceIdColumnOk returns a tuple with the SequenceIdColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceIdColumn

`func (o *CreateDataSourceRequestDataSourceConfig) SetSequenceIdColumn(v string)`

SetSequenceIdColumn sets SequenceIdColumn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


