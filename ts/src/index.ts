import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { UseShellService } from "./gen/clutch/v1/use_shell_pb"

const transport = createConnectTransport({
  baseUrl: "http://localhost:5051",
});


const client = createClient(UseShellService, transport);

export async function useShell(appName: string, command: string) {
  const request = {
    appName,
    command,
  };

  try {
    const response = await client.useShell(request);
    return response;
  } catch (error) {
    console.error("Error using shell:", error);
    throw error;
  }
}

async function testUseShell() {
  try {
    const response = await useShell("exampleApp", "echo Hello, World!");
    console.log("Shell command output:", response.output);
  } catch (error) {
    console.error("Error executing shell command:", error);
  }
}

testUseShell()
