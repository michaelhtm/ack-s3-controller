// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type AnalyticsS3ExportFileFormat string

const (
	AnalyticsS3ExportFileFormat_CSV AnalyticsS3ExportFileFormat = "CSV"
)

type ArchiveStatus string

const (
	ArchiveStatus_ARCHIVE_ACCESS      ArchiveStatus = "ARCHIVE_ACCESS"
	ArchiveStatus_DEEP_ARCHIVE_ACCESS ArchiveStatus = "DEEP_ARCHIVE_ACCESS"
)

type BucketAccelerateStatus string

const (
	BucketAccelerateStatus_Enabled   BucketAccelerateStatus = "Enabled"
	BucketAccelerateStatus_Suspended BucketAccelerateStatus = "Suspended"
)

type BucketCannedACL string

const (
	BucketCannedACL_authenticated_read BucketCannedACL = "authenticated-read"
	BucketCannedACL_private            BucketCannedACL = "private"
	BucketCannedACL_public_read        BucketCannedACL = "public-read"
	BucketCannedACL_public_read_write  BucketCannedACL = "public-read-write"
)

type BucketLocationConstraint string

const (
	BucketLocationConstraint_EU             BucketLocationConstraint = "EU"
	BucketLocationConstraint_af_south_1     BucketLocationConstraint = "af-south-1"
	BucketLocationConstraint_ap_east_1      BucketLocationConstraint = "ap-east-1"
	BucketLocationConstraint_ap_northeast_1 BucketLocationConstraint = "ap-northeast-1"
	BucketLocationConstraint_ap_northeast_2 BucketLocationConstraint = "ap-northeast-2"
	BucketLocationConstraint_ap_northeast_3 BucketLocationConstraint = "ap-northeast-3"
	BucketLocationConstraint_ap_south_1     BucketLocationConstraint = "ap-south-1"
	BucketLocationConstraint_ap_south_2     BucketLocationConstraint = "ap-south-2"
	BucketLocationConstraint_ap_southeast_1 BucketLocationConstraint = "ap-southeast-1"
	BucketLocationConstraint_ap_southeast_2 BucketLocationConstraint = "ap-southeast-2"
	BucketLocationConstraint_ap_southeast_3 BucketLocationConstraint = "ap-southeast-3"
	BucketLocationConstraint_ca_central_1   BucketLocationConstraint = "ca-central-1"
	BucketLocationConstraint_cn_north_1     BucketLocationConstraint = "cn-north-1"
	BucketLocationConstraint_cn_northwest_1 BucketLocationConstraint = "cn-northwest-1"
	BucketLocationConstraint_eu_central_1   BucketLocationConstraint = "eu-central-1"
	BucketLocationConstraint_eu_north_1     BucketLocationConstraint = "eu-north-1"
	BucketLocationConstraint_eu_south_1     BucketLocationConstraint = "eu-south-1"
	BucketLocationConstraint_eu_south_2     BucketLocationConstraint = "eu-south-2"
	BucketLocationConstraint_eu_west_1      BucketLocationConstraint = "eu-west-1"
	BucketLocationConstraint_eu_west_2      BucketLocationConstraint = "eu-west-2"
	BucketLocationConstraint_eu_west_3      BucketLocationConstraint = "eu-west-3"
	BucketLocationConstraint_me_south_1     BucketLocationConstraint = "me-south-1"
	BucketLocationConstraint_sa_east_1      BucketLocationConstraint = "sa-east-1"
	BucketLocationConstraint_us_east_2      BucketLocationConstraint = "us-east-2"
	BucketLocationConstraint_us_gov_east_1  BucketLocationConstraint = "us-gov-east-1"
	BucketLocationConstraint_us_gov_west_1  BucketLocationConstraint = "us-gov-west-1"
	BucketLocationConstraint_us_west_1      BucketLocationConstraint = "us-west-1"
	BucketLocationConstraint_us_west_2      BucketLocationConstraint = "us-west-2"
)

type BucketLogsPermission string

const (
	BucketLogsPermission_FULL_CONTROL BucketLogsPermission = "FULL_CONTROL"
	BucketLogsPermission_READ         BucketLogsPermission = "READ"
	BucketLogsPermission_WRITE        BucketLogsPermission = "WRITE"
)

type BucketType string

const (
	BucketType_Directory BucketType = "Directory"
)

type BucketVersioningStatus string

const (
	BucketVersioningStatus_Enabled   BucketVersioningStatus = "Enabled"
	BucketVersioningStatus_Suspended BucketVersioningStatus = "Suspended"
)

type ChecksumAlgorithm string

const (
	ChecksumAlgorithm_CRC32  ChecksumAlgorithm = "CRC32"
	ChecksumAlgorithm_CRC32C ChecksumAlgorithm = "CRC32C"
	ChecksumAlgorithm_SHA1   ChecksumAlgorithm = "SHA1"
	ChecksumAlgorithm_SHA256 ChecksumAlgorithm = "SHA256"
)

type ChecksumMode string

const (
	ChecksumMode_ENABLED ChecksumMode = "ENABLED"
)

type CompressionType string

const (
	CompressionType_BZIP2 CompressionType = "BZIP2"
	CompressionType_GZIP  CompressionType = "GZIP"
	CompressionType_NONE  CompressionType = "NONE"
)

type DataRedundancy string

const (
	DataRedundancy_SingleAvailabilityZone DataRedundancy = "SingleAvailabilityZone"
	DataRedundancy_SingleLocalZone        DataRedundancy = "SingleLocalZone"
)

type DeleteMarkerReplicationStatus string

const (
	DeleteMarkerReplicationStatus_Disabled DeleteMarkerReplicationStatus = "Disabled"
	DeleteMarkerReplicationStatus_Enabled  DeleteMarkerReplicationStatus = "Enabled"
)

type EncodingType string

const (
	EncodingType_url EncodingType = "url"
)

type Event string

const (
	Event_s3_IntelligentTiering                            Event = "s3:IntelligentTiering"
	Event_s3_LifecycleExpiration__                         Event = "s3:LifecycleExpiration:*"
	Event_s3_LifecycleExpiration_Delete                    Event = "s3:LifecycleExpiration:Delete"
	Event_s3_LifecycleExpiration_DeleteMarkerCreated       Event = "s3:LifecycleExpiration:DeleteMarkerCreated"
	Event_s3_LifecycleTransition                           Event = "s3:LifecycleTransition"
	Event_s3_ObjectAcl_Put                                 Event = "s3:ObjectAcl:Put"
	Event_s3_ObjectCreated__                               Event = "s3:ObjectCreated:*"
	Event_s3_ObjectCreated_CompleteMultipartUpload         Event = "s3:ObjectCreated:CompleteMultipartUpload"
	Event_s3_ObjectCreated_Copy                            Event = "s3:ObjectCreated:Copy"
	Event_s3_ObjectCreated_Post                            Event = "s3:ObjectCreated:Post"
	Event_s3_ObjectCreated_Put                             Event = "s3:ObjectCreated:Put"
	Event_s3_ObjectRemoved__                               Event = "s3:ObjectRemoved:*"
	Event_s3_ObjectRemoved_Delete                          Event = "s3:ObjectRemoved:Delete"
	Event_s3_ObjectRemoved_DeleteMarkerCreated             Event = "s3:ObjectRemoved:DeleteMarkerCreated"
	Event_s3_ObjectRestore__                               Event = "s3:ObjectRestore:*"
	Event_s3_ObjectRestore_Completed                       Event = "s3:ObjectRestore:Completed"
	Event_s3_ObjectRestore_Delete                          Event = "s3:ObjectRestore:Delete"
	Event_s3_ObjectRestore_Post                            Event = "s3:ObjectRestore:Post"
	Event_s3_ObjectTagging__                               Event = "s3:ObjectTagging:*"
	Event_s3_ObjectTagging_Delete                          Event = "s3:ObjectTagging:Delete"
	Event_s3_ObjectTagging_Put                             Event = "s3:ObjectTagging:Put"
	Event_s3_ReducedRedundancyLostObject                   Event = "s3:ReducedRedundancyLostObject"
	Event_s3_Replication__                                 Event = "s3:Replication:*"
	Event_s3_Replication_OperationFailedReplication        Event = "s3:Replication:OperationFailedReplication"
	Event_s3_Replication_OperationMissedThreshold          Event = "s3:Replication:OperationMissedThreshold"
	Event_s3_Replication_OperationNotTracked               Event = "s3:Replication:OperationNotTracked"
	Event_s3_Replication_OperationReplicatedAfterThreshold Event = "s3:Replication:OperationReplicatedAfterThreshold"
)

type ExistingObjectReplicationStatus string

const (
	ExistingObjectReplicationStatus_Disabled ExistingObjectReplicationStatus = "Disabled"
	ExistingObjectReplicationStatus_Enabled  ExistingObjectReplicationStatus = "Enabled"
)

type ExpirationStatus string

const (
	ExpirationStatus_Disabled ExpirationStatus = "Disabled"
	ExpirationStatus_Enabled  ExpirationStatus = "Enabled"
)

type ExpressionType string

const (
	ExpressionType_SQL ExpressionType = "SQL"
)

type FileHeaderInfo string

const (
	FileHeaderInfo_IGNORE FileHeaderInfo = "IGNORE"
	FileHeaderInfo_NONE   FileHeaderInfo = "NONE"
	FileHeaderInfo_USE    FileHeaderInfo = "USE"
)

type FilterRuleName string

const (
	FilterRuleName_prefix FilterRuleName = "prefix"
	FilterRuleName_suffix FilterRuleName = "suffix"
)

type IntelligentTieringAccessTier string

const (
	IntelligentTieringAccessTier_ARCHIVE_ACCESS      IntelligentTieringAccessTier = "ARCHIVE_ACCESS"
	IntelligentTieringAccessTier_DEEP_ARCHIVE_ACCESS IntelligentTieringAccessTier = "DEEP_ARCHIVE_ACCESS"
)

type IntelligentTieringStatus string

const (
	IntelligentTieringStatus_Disabled IntelligentTieringStatus = "Disabled"
	IntelligentTieringStatus_Enabled  IntelligentTieringStatus = "Enabled"
)

type InventoryFormat string

const (
	InventoryFormat_CSV     InventoryFormat = "CSV"
	InventoryFormat_ORC     InventoryFormat = "ORC"
	InventoryFormat_Parquet InventoryFormat = "Parquet"
)

type InventoryFrequency string

const (
	InventoryFrequency_Daily  InventoryFrequency = "Daily"
	InventoryFrequency_Weekly InventoryFrequency = "Weekly"
)

type InventoryIncludedObjectVersions string

const (
	InventoryIncludedObjectVersions_All     InventoryIncludedObjectVersions = "All"
	InventoryIncludedObjectVersions_Current InventoryIncludedObjectVersions = "Current"
)

type InventoryOptionalField string

const (
	InventoryOptionalField_BucketKeyStatus              InventoryOptionalField = "BucketKeyStatus"
	InventoryOptionalField_ChecksumAlgorithm            InventoryOptionalField = "ChecksumAlgorithm"
	InventoryOptionalField_ETag                         InventoryOptionalField = "ETag"
	InventoryOptionalField_EncryptionStatus             InventoryOptionalField = "EncryptionStatus"
	InventoryOptionalField_IntelligentTieringAccessTier InventoryOptionalField = "IntelligentTieringAccessTier"
	InventoryOptionalField_IsMultipartUploaded          InventoryOptionalField = "IsMultipartUploaded"
	InventoryOptionalField_LastModifiedDate             InventoryOptionalField = "LastModifiedDate"
	InventoryOptionalField_ObjectAccessControlList      InventoryOptionalField = "ObjectAccessControlList"
	InventoryOptionalField_ObjectLockLegalHoldStatus    InventoryOptionalField = "ObjectLockLegalHoldStatus"
	InventoryOptionalField_ObjectLockMode               InventoryOptionalField = "ObjectLockMode"
	InventoryOptionalField_ObjectLockRetainUntilDate    InventoryOptionalField = "ObjectLockRetainUntilDate"
	InventoryOptionalField_ObjectOwner                  InventoryOptionalField = "ObjectOwner"
	InventoryOptionalField_ReplicationStatus            InventoryOptionalField = "ReplicationStatus"
	InventoryOptionalField_Size                         InventoryOptionalField = "Size"
	InventoryOptionalField_StorageClass                 InventoryOptionalField = "StorageClass"
)

type JSONType string

const (
	JSONType_DOCUMENT JSONType = "DOCUMENT"
	JSONType_LINES    JSONType = "LINES"
)

type LocationType string

const (
	LocationType_AvailabilityZone LocationType = "AvailabilityZone"
	LocationType_LocalZone        LocationType = "LocalZone"
)

type MFADelete string

const (
	MFADelete_Disabled MFADelete = "Disabled"
	MFADelete_Enabled  MFADelete = "Enabled"
)

type MFADeleteStatus string

const (
	MFADeleteStatus_Disabled MFADeleteStatus = "Disabled"
	MFADeleteStatus_Enabled  MFADeleteStatus = "Enabled"
)

type MetadataDirective string

const (
	MetadataDirective_COPY    MetadataDirective = "COPY"
	MetadataDirective_REPLACE MetadataDirective = "REPLACE"
)

type MetricsStatus string

const (
	MetricsStatus_Disabled MetricsStatus = "Disabled"
	MetricsStatus_Enabled  MetricsStatus = "Enabled"
)

type ObjectAttributes string

const (
	ObjectAttributes_Checksum     ObjectAttributes = "Checksum"
	ObjectAttributes_ETag         ObjectAttributes = "ETag"
	ObjectAttributes_ObjectParts  ObjectAttributes = "ObjectParts"
	ObjectAttributes_ObjectSize   ObjectAttributes = "ObjectSize"
	ObjectAttributes_StorageClass ObjectAttributes = "StorageClass"
)

type ObjectCannedACL string

const (
	ObjectCannedACL_authenticated_read        ObjectCannedACL = "authenticated-read"
	ObjectCannedACL_aws_exec_read             ObjectCannedACL = "aws-exec-read"
	ObjectCannedACL_bucket_owner_full_control ObjectCannedACL = "bucket-owner-full-control"
	ObjectCannedACL_bucket_owner_read         ObjectCannedACL = "bucket-owner-read"
	ObjectCannedACL_private                   ObjectCannedACL = "private"
	ObjectCannedACL_public_read               ObjectCannedACL = "public-read"
	ObjectCannedACL_public_read_write         ObjectCannedACL = "public-read-write"
)

type ObjectLockEnabled string

const (
	ObjectLockEnabled_Enabled ObjectLockEnabled = "Enabled"
)

type ObjectLockLegalHoldStatus string

const (
	ObjectLockLegalHoldStatus_OFF ObjectLockLegalHoldStatus = "OFF"
	ObjectLockLegalHoldStatus_ON  ObjectLockLegalHoldStatus = "ON"
)

type ObjectLockMode string

const (
	ObjectLockMode_COMPLIANCE ObjectLockMode = "COMPLIANCE"
	ObjectLockMode_GOVERNANCE ObjectLockMode = "GOVERNANCE"
)

type ObjectLockRetentionMode string

const (
	ObjectLockRetentionMode_COMPLIANCE ObjectLockRetentionMode = "COMPLIANCE"
	ObjectLockRetentionMode_GOVERNANCE ObjectLockRetentionMode = "GOVERNANCE"
)

type ObjectOwnership string

const (
	ObjectOwnership_BucketOwnerEnforced  ObjectOwnership = "BucketOwnerEnforced"
	ObjectOwnership_BucketOwnerPreferred ObjectOwnership = "BucketOwnerPreferred"
	ObjectOwnership_ObjectWriter         ObjectOwnership = "ObjectWriter"
)

type ObjectStorageClass string

const (
	ObjectStorageClass_DEEP_ARCHIVE        ObjectStorageClass = "DEEP_ARCHIVE"
	ObjectStorageClass_EXPRESS_ONEZONE     ObjectStorageClass = "EXPRESS_ONEZONE"
	ObjectStorageClass_GLACIER             ObjectStorageClass = "GLACIER"
	ObjectStorageClass_GLACIER_IR          ObjectStorageClass = "GLACIER_IR"
	ObjectStorageClass_INTELLIGENT_TIERING ObjectStorageClass = "INTELLIGENT_TIERING"
	ObjectStorageClass_ONEZONE_IA          ObjectStorageClass = "ONEZONE_IA"
	ObjectStorageClass_OUTPOSTS            ObjectStorageClass = "OUTPOSTS"
	ObjectStorageClass_REDUCED_REDUNDANCY  ObjectStorageClass = "REDUCED_REDUNDANCY"
	ObjectStorageClass_SNOW                ObjectStorageClass = "SNOW"
	ObjectStorageClass_STANDARD            ObjectStorageClass = "STANDARD"
	ObjectStorageClass_STANDARD_IA         ObjectStorageClass = "STANDARD_IA"
)

type ObjectVersionStorageClass string

const (
	ObjectVersionStorageClass_STANDARD ObjectVersionStorageClass = "STANDARD"
)

type OptionalObjectAttributes string

const (
	OptionalObjectAttributes_RestoreStatus OptionalObjectAttributes = "RestoreStatus"
)

type OwnerOverride string

const (
	OwnerOverride_Destination OwnerOverride = "Destination"
)

type PartitionDateSource string

const (
	PartitionDateSource_DeliveryTime PartitionDateSource = "DeliveryTime"
	PartitionDateSource_EventTime    PartitionDateSource = "EventTime"
)

type Payer string

const (
	Payer_BucketOwner Payer = "BucketOwner"
	Payer_Requester   Payer = "Requester"
)

type Permission string

const (
	Permission_FULL_CONTROL Permission = "FULL_CONTROL"
	Permission_READ         Permission = "READ"
	Permission_READ_ACP     Permission = "READ_ACP"
	Permission_WRITE        Permission = "WRITE"
	Permission_WRITE_ACP    Permission = "WRITE_ACP"
)

type Protocol string

const (
	Protocol_http  Protocol = "http"
	Protocol_https Protocol = "https"
)

type QuoteFields string

const (
	QuoteFields_ALWAYS   QuoteFields = "ALWAYS"
	QuoteFields_ASNEEDED QuoteFields = "ASNEEDED"
)

type ReplicaModificationsStatus string

const (
	ReplicaModificationsStatus_Disabled ReplicaModificationsStatus = "Disabled"
	ReplicaModificationsStatus_Enabled  ReplicaModificationsStatus = "Enabled"
)

type ReplicationRuleStatus string

const (
	ReplicationRuleStatus_Disabled ReplicationRuleStatus = "Disabled"
	ReplicationRuleStatus_Enabled  ReplicationRuleStatus = "Enabled"
)

type ReplicationStatus string

const (
	ReplicationStatus_COMPLETE  ReplicationStatus = "COMPLETE"
	ReplicationStatus_COMPLETED ReplicationStatus = "COMPLETED"
	ReplicationStatus_FAILED    ReplicationStatus = "FAILED"
	ReplicationStatus_PENDING   ReplicationStatus = "PENDING"
	ReplicationStatus_REPLICA   ReplicationStatus = "REPLICA"
)

type ReplicationTimeStatus string

const (
	ReplicationTimeStatus_Disabled ReplicationTimeStatus = "Disabled"
	ReplicationTimeStatus_Enabled  ReplicationTimeStatus = "Enabled"
)

type RequestCharged string

const (
	RequestCharged_requester RequestCharged = "requester"
)

type RequestPayer string

const (
	RequestPayer_requester RequestPayer = "requester"
)

type RestoreRequestType string

const (
	RestoreRequestType_SELECT RestoreRequestType = "SELECT"
)

type SSEKMSEncryptedObjectsStatus string

const (
	SSEKMSEncryptedObjectsStatus_Disabled SSEKMSEncryptedObjectsStatus = "Disabled"
	SSEKMSEncryptedObjectsStatus_Enabled  SSEKMSEncryptedObjectsStatus = "Enabled"
)

type ServerSideEncryption string

const (
	ServerSideEncryption_AES256       ServerSideEncryption = "AES256"
	ServerSideEncryption_aws_kms      ServerSideEncryption = "aws:kms"
	ServerSideEncryption_aws_kms_dsse ServerSideEncryption = "aws:kms:dsse"
)

type SessionMode string

const (
	SessionMode_ReadOnly  SessionMode = "ReadOnly"
	SessionMode_ReadWrite SessionMode = "ReadWrite"
)

type StorageClass string

const (
	StorageClass_DEEP_ARCHIVE        StorageClass = "DEEP_ARCHIVE"
	StorageClass_EXPRESS_ONEZONE     StorageClass = "EXPRESS_ONEZONE"
	StorageClass_GLACIER             StorageClass = "GLACIER"
	StorageClass_GLACIER_IR          StorageClass = "GLACIER_IR"
	StorageClass_INTELLIGENT_TIERING StorageClass = "INTELLIGENT_TIERING"
	StorageClass_ONEZONE_IA          StorageClass = "ONEZONE_IA"
	StorageClass_OUTPOSTS            StorageClass = "OUTPOSTS"
	StorageClass_REDUCED_REDUNDANCY  StorageClass = "REDUCED_REDUNDANCY"
	StorageClass_SNOW                StorageClass = "SNOW"
	StorageClass_STANDARD            StorageClass = "STANDARD"
	StorageClass_STANDARD_IA         StorageClass = "STANDARD_IA"
)

type StorageClassAnalysisSchemaVersion string

const (
	StorageClassAnalysisSchemaVersion_V_1 StorageClassAnalysisSchemaVersion = "V_1"
)

type TaggingDirective string

const (
	TaggingDirective_COPY    TaggingDirective = "COPY"
	TaggingDirective_REPLACE TaggingDirective = "REPLACE"
)

type Tier string

const (
	Tier_Bulk      Tier = "Bulk"
	Tier_Expedited Tier = "Expedited"
	Tier_Standard  Tier = "Standard"
)

type TransitionDefaultMinimumObjectSize string

const (
	TransitionDefaultMinimumObjectSize_all_storage_classes_128K TransitionDefaultMinimumObjectSize = "all_storage_classes_128K"
	TransitionDefaultMinimumObjectSize_varies_by_storage_class  TransitionDefaultMinimumObjectSize = "varies_by_storage_class"
)

type TransitionStorageClass string

const (
	TransitionStorageClass_DEEP_ARCHIVE        TransitionStorageClass = "DEEP_ARCHIVE"
	TransitionStorageClass_GLACIER             TransitionStorageClass = "GLACIER"
	TransitionStorageClass_GLACIER_IR          TransitionStorageClass = "GLACIER_IR"
	TransitionStorageClass_INTELLIGENT_TIERING TransitionStorageClass = "INTELLIGENT_TIERING"
	TransitionStorageClass_ONEZONE_IA          TransitionStorageClass = "ONEZONE_IA"
	TransitionStorageClass_STANDARD_IA         TransitionStorageClass = "STANDARD_IA"
)

type Type string

const (
	Type_AmazonCustomerByEmail Type = "AmazonCustomerByEmail"
	Type_CanonicalUser         Type = "CanonicalUser"
	Type_Group                 Type = "Group"
)
