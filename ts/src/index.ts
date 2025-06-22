import { createClient as connectCreateClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { UseShellService } from "./gen/clutch/v1/use_shell_pb";
import { ToggleWindowService } from "./gen/clutch/v1/toggle_window_pb";
import { SayHiService } from "./gen/clutch/v1/say_hi_pb";

type SayHiReq = {
  greet: string
}

type UseShellReq = {
  appName: string
  command: string
}

type ToggleWindowReq = {
  isVisible: boolean
}

export function createClient(port: number) {
  const transport = createConnectTransport({
    baseUrl: `http://127.0.0.1:${port}`,
  });

  const shellClient = connectCreateClient(UseShellService, transport);
  const windowClient = connectCreateClient(ToggleWindowService, transport);
  const sayHiClient = connectCreateClient(SayHiService, transport);

  return {
    async sayHi({ greet }: SayHiReq) {
      const request: SayHiReq = { greet }

      try {
        const response = await sayHiClient.sayHi(request);
        return response;
      } catch (error) {
        console.error("Error using shell:", error);
        throw error;
      }
    },

    async useShell({ appName, command }: UseShellReq) {
      const request: UseShellReq = { appName, command };

      try {
        const response = await shellClient.useShell(request);
        return response;
      } catch (error) {
        console.error("Error using shell:", error);
        throw error;
      }
    },
    async toggleWindow({ isVisible }: ToggleWindowReq) {
      const request: ToggleWindowReq = { isVisible }
      try {
        const response = await windowClient.toggleWindow(request);
        return response;
      } catch (error) {
        console.error("Error toggling window:", error);
        throw error;
      }
    }
  };
}

