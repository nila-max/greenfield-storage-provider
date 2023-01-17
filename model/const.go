package model

// define storage provider support service name.
const (
	GatewayService    = "Gateway"
	UploaderService   = "Uploader"
	DownloaderService = "Downloader"
	ChallengeService  = "Challenge"
	StoneHubService   = "StoneHub"
	StoneNodeService  = "StoneNode"
	SyncerService     = "Syncer"
)

// define payload data redundancy size.
const (
	InlineSize  = 1 * 1024 * 1024
	SegmentSize = 16 * 1024 * 1024
	EC_M        = 4
	EC_K        = 2
)

type PieceType int32

const (
	SegmentPieceType PieceType = 0
	ECPieceType      PieceType = 1
)

// define piece store constants.
const (
	BufPoolSize  = 32 << 10
	ChecksumAlgo = "Crc32c"
	OctetStream  = "application/octet-stream"
)

// RPC config
const (
	// server and client max send or recv msg size
	MaxCallMsgSize = 40 * 1024 * 1024
)

// Gateway
const (
	// path
	AdminPath          = "/greenfield/admin/v1/"
	GetApprovalSubPath = "get-approval"

	// query key
	TransactionQuery = "transaction"
	PutObjectV2Query = "putobjectv2"
	ActionQuery      = "action"

	// bfs header key
	BFSRequestIDHeader       = "X-Bfs-Request-ID"
	BFSContentLengthHeader   = "X-Bfs-Content-Length"
	BFSContentTypeHeader     = "X-Bfs-Content-Type"
	BFSChecksumHeader        = "X-Bfs-Checksum"
	BFSIsPrivateHeader       = "X-Bfs-Is-Private"
	BFSTransactionHashHeader = "X-Bfs-Txn-Hash"
	BFSResourceHeader        = "X-Bfs-Resource"
	BFSPreSignatureHeader    = "X-Bfs-Pre-Signature"
	// BFSRedundancyTypeHeader can be EC or Replica, EC is default
	BFSRedundancyTypeHeader = "X-Bfs-Redundancy-Type"

	// http header key
	ContentTypeHeader   = "Content-Type"
	ETagHeader          = "ETag"
	ContentLengthHeader = "Content-Length"

	// header value
	ContentTypeXMLHeaderValue        = "application/xml"
	ReplicaRedundancyTypeHeaderValue = "Replica"
)

// define backend store type name.
const (
	MemoryDB string = "memory"
	MySqlDB  string = "mysql"
	LevelDB  string = "leveldb"
)