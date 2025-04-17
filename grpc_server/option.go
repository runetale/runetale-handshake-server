package grpcserver

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	keepAliveTime    = 30 * time.Second // Ping interval for clients
	keepAliveTimeout = 10 * time.Second // Timeout for client's response
)

func NewGrpcServerOption() []grpc.ServerOption {
	var opts []grpc.ServerOption

	enforcementPol := keepalive.EnforcementPolicy{
		// MinTime: クライアントが keepalive ping を送信する前に待機する必要がある最小時間を指定します。
		//この時間が経過する前にクライアントが ping を送信すると、サーバーは接続を閉じる可能性があります。デフォルト値は5分です。
		MinTime: 5 * time.Minute,
		// PermitWithoutStream: このフィールドが true の場合、サーバーはアクティブなストリーム（RPC）がない場合でも keepalive ping を許可します。
		//false の場合、アクティブなストリームがない状態でクライアントが ping を送信すると、サーバーは GOAWAY メッセージを送信して接続を閉じます。デフォルト値は false です。
		PermitWithoutStream: true,
	}

	serverParams := keepalive.ServerParameters{
		// Time: サーバーがクライアントに対して keepalive ping を送信するまでのアイドル時間を指定します。
		//この時間が経過すると、サーバーはクライアントに ping を送信して接続がまだ生きているかどうかを確認します。デフォルト値は2時間です。
		Time: keepAliveTime,
		// Timeout: サーバーが keepalive ping を送信した後、クライアントからの応答を待つ時間を指定します。
		// この時間内にクライアントからの応答がない場合、サーバーは接続を閉じます。デフォルト値は20秒です。
		Timeout: keepAliveTimeout,
	}
	opts = append(opts, grpc.KeepaliveEnforcementPolicy(enforcementPol), grpc.KeepaliveParams(serverParams))

	return opts
}
