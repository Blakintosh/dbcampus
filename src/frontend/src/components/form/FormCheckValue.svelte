<script lang="ts">
    export let key: string;
    export let value: string;
    export let divide: boolean = false;
</script>

<tr>
    <th class="p-2 px-4 font-semibold bg-slate-600 text-slate-50 text-sm {divide ? "border-b-2" : "border-b"} border-b-slate-500">{key}</th>
    <td class="p-1 px-2 bg-slate-100 text-sm {divide ? "border-b" : "border-b"} border-b-slate-300">{value}</td>
</tr>