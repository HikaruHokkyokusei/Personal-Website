<script lang="ts">
    import { goto } from "$app/navigation";
    import { genericDataStore } from "$lib/stores/GenericDataStore";

    if ($genericDataStore.defaultRoute && $genericDataStore.defaultRoute != "/") {
        goto($genericDataStore.defaultRoute, { "replaceState": true })
    }
</script>
