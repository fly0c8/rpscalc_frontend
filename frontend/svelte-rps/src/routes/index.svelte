<script>
  import { onMount } from "svelte";
  import { connect, StringCodec } from "nats.ws";

  let nc = null;
  const sc = StringCodec();

  let connecting = true;
  let sending = false;
  let error = null;
  let responseText = "";

  const NATS_URL = "ws://localhost:8081"; // change if needed

  const payload = {
    Varset: "TO_MAX_CALC_21",
    Config: {
      Flaps: 1,
      AntiIcing: 0,
      AirCond: 0
    },
    Rwy: {
      Elevation: 777,
      Tora: 12000,
      Toda: 12000,
      Asda: 12000,
      Width: 150,
      SlopeToda: 0.2,
      SlopeAsda: 0.1
    },
    Wx: {
      Oat: 15,
      Qnh: 29.92,
      Wc: 8,
      Xwc: 8
    },
    Ptow: 190000
  };

  onMount(async () => {
    try {
      nc = await connect({ servers: NATS_URL });
    } catch (err) {
      error = "Failed to connect to NATS: " + err.message;
    } finally {
      connecting = false;
    }

    return () => {
      if (nc) {
        nc.close();
      }
    };
  });

  async function sendRequest() {
    if (!nc) return;

    sending = true;
    error = null;
    responseText = "";

    try {
      const jsonText = JSON.stringify(payload);

      const msg = await nc.request(
        "calc.flex",
        sc.encode(jsonText),
        { timeout: 60000 }   // 60 seconds
      );

      responseText = sc.decode(msg.data);
    } catch (err) {
      if (err.code === "TIMEOUT") {
        error = "Request timed out.";
      } else {
        error = err.message;
      }
    } finally {
      sending = false;
    }
  }
</script>

<div class="space-y-4 max-w-2xl">
  <h1 class="text-3xl font-bold mb-4">Flex Calculation via NATS</h1>

  {#if connecting}
    <p class="text-slate-400">Connecting to NATS…</p>
  {:else if error && !nc}
    <p class="text-red-400">Error: {error}</p>
  {:else}

    <button
      class="inline-flex items-center rounded-md border border-slate-700 px-4 py-2 text-lg font-medium hover:bg-slate-800 disabled:opacity-50"
      on:click={sendRequest}
      disabled={sending}
    >
      {#if sending}
        Sending…
      {:else}
        Send Request
      {/if}
    </button>

    {#if error}
      <p class="text-red-400 font-semibold">{error}</p>
    {/if}

    {#if responseText}
      <div class="mt-4">
        <h2 class="text-xl font-semibold mb-2">Response</h2>
        <pre class="text-sm bg-slate-950 border border-slate-800 rounded-lg p-4 overflow-auto">
{responseText}
        </pre>
      </div>
    {/if}

  {/if}
</div>
