import { createClient as connectCreateClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { UseShellService, type ConfirmShellRequest, type UseShellRequest } from "./gen/clutch/v1/use_shell_pb";
import { ToggleWindowService } from "./gen/clutch/v1/toggle_window_pb";
import { GreetService, type GreetRequest } from "./gen/clutch/v1/greet_pb";

export function createClient(port: number) {
  const transport = createConnectTransport({
    baseUrl: `http://127.0.0.1:${port}`,
  });

  const shellClient = connectCreateClient(UseShellService, transport);
  const windowClient = connectCreateClient(ToggleWindowService, transport);
  const greetClient = connectCreateClient(GreetService, transport);

  return {
    async greet({ name }: Pick<GreetRequest, "name">) {
      try {
        const response = await greetClient.greet({ name });
        return response;
      } catch (error) {
        console.error("Error using shell:", error);
        throw error;
      }
    },

    async toggleWindow() {
      try {
        const response = await windowClient.toggleWindow({});
        return response;
      } catch (error) {
        console.error("Error toggling window:", error);
        throw error;
      }
    },

    async useShell({ appName, command }: Pick<UseShellRequest, "appName" | "command">) {
      try {
        const response = await shellClient.useShell({ appName, command });
        return response;
      } catch (error) {
        console.error("Error using shell:", error);
        throw error;
      }
    },

    async confirmShell({ allow }: Pick<ConfirmShellRequest, "allow">) {
      try {
        const response = await shellClient.confirmShell({ allow });
        return response;
      } catch (error) {
        console.error("Error confirming shell:", error);
        throw error;
      }
    },

  };
}
