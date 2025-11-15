<script>
  import { isActiveLink } from "sv-router";

  let { children } = $props();
</script>

<div class="h-screen w-screen flex overflow-hidden bg-slate-900 text-slate-100">
  <!-- LEFT SIDEBAR -->
  <aside class="w-56 bg-slate-950 border-r border-slate-800 flex flex-col">
    <div class="p-4 text-2xl font-bold tracking-tight">
      RPS
    </div>

    <!-- text-lg = larger font for menu items -->
    <nav class="flex-1 px-2 space-y-1 text-lg">
      <a
        href="/"
        use:isActiveLink={{ className: "bg-slate-800 text-cyan-400 font-semibold" }}
        class="block px-3 py-2 rounded-lg transition hover:bg-slate-800"
      >
        Home
      </a>

      <a
        href="/about"
        use:isActiveLink={{ className: "bg-slate-800 text-cyan-400 font-semibold" }}
        class="block px-3 py-2 rounded-lg transition hover:bg-slate-800"
      >
        About
      </a>

      <a
        href="/settings"
        use:isActiveLink={{ className: "bg-slate-800 text-cyan-400 font-semibold" }}
        class="block px-3 py-2 rounded-lg transition hover:bg-slate-800"
      >
        Settings
      </a>
    </nav>

    <div class="p-4 text-xs text-slate-500">
      © 2025 zeroc8
    </div>
  </aside>

  <!-- MAIN CONTENT AREA -->
  <main class="flex-1 overflow-auto p-6">
    {@render children()}
  </main>
</div>
