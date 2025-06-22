import * as grpcWeb from 'grpc-web';

import * as clutch_v1_toggle_window_pb from '../../clutch/v1/toggle_window_pb'; // proto import: "clutch/v1/toggle_window.proto"


export class ToggleWindowServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  toggleWindow(
    request: clutch_v1_toggle_window_pb.ToggleWindowRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: clutch_v1_toggle_window_pb.ToggleWindowResponse) => void
  ): grpcWeb.ClientReadableStream<clutch_v1_toggle_window_pb.ToggleWindowResponse>;

}

export class ToggleWindowServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  toggleWindow(
    request: clutch_v1_toggle_window_pb.ToggleWindowRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<clutch_v1_toggle_window_pb.ToggleWindowResponse>;

}

