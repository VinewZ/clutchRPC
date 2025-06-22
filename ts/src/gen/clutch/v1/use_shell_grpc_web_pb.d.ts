import * as grpcWeb from 'grpc-web';

import * as clutch_v1_use_shell_pb from '../../clutch/v1/use_shell_pb'; // proto import: "clutch/v1/use_shell.proto"


export class UseShellServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  useShell(
    request: clutch_v1_use_shell_pb.UseShellRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: clutch_v1_use_shell_pb.UseShellResponse) => void
  ): grpcWeb.ClientReadableStream<clutch_v1_use_shell_pb.UseShellResponse>;

}

export class UseShellServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  useShell(
    request: clutch_v1_use_shell_pb.UseShellRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<clutch_v1_use_shell_pb.UseShellResponse>;

}

