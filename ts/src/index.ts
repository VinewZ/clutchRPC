import { createClient as connectCreateClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import type { UseShellRequest, UseShellResponse } from "./gen/clutch/v1/use_shell_pb";
import type { ToggleWindowResponse } from "./gen/clutch/v1/toggle_window_pb";
import type { GreetRequest, GreetResponse } from "./gen/clutch/v1/greet_pb";
import { UseShellService } from "./gen/clutch/v1/use_shell_pb";
import { ToggleWindowService } from "./gen/clutch/v1/toggle_window_pb";
import { GreetService } from "./gen/clutch/v1/greet_pb";

export function createClient(port: number) {
  const transport = createConnectTransport({
    baseUrl: `http://127.0.0.1:${port}`,
  });

  const shellClient = connectCreateClient(UseShellService, transport);
  const windowClient = connectCreateClient(ToggleWindowService, transport);
  const greetClient = connectCreateClient(GreetService, transport);

  return {
    async greet({ name }: Pick<GreetRequest, "name">): Promise<GreetResponse> {
      try {
        const response = await greetClient.greet({ name });
        return response;
      } catch (error) {
        console.error("Error using shell:", error);
        return Promise.reject(error);
      }
    },

    async toggleWindow(): Promise<ToggleWindowResponse> {
      try {
        const response = await windowClient.toggleWindow({});
        return response;
      } catch (error) {
        console.error("Error toggling window:", error);
        return Promise.reject(error);
      }
    },

    async useShell({ appName, command }: Pick<UseShellRequest, "appName" | "command">): Promise<UseShellResponse> {
      try {
        const response = await shellClient.useShell({ appName, command });
        return response;
      } catch (error) {
        console.error("Error using shell:", error);
        return Promise.reject(error);
      }
    },

  };
}
