<script>
	import SmallButton from "../../../../common/SmallButton.svelte";
	import SettingField from "../components/SettingField.svelte";

</script>
<SettingField title="Regular Project Cost" description="This is the expected regular cost of the project that will be deducted every month.">
    <input type="text" class=" bg-slate-800 text-slate-100 py-2 px-4 rounded-sm" placeholder="Monthly Cost"/>
</SettingField>

<SettingField title="Add Unforeseen Cost" description="Deduct an additional, one-off cost from the project budget." >
    <input type="text" class=" bg-slate-800 text-slate-100 py-2 px-4 rounded-sm" placeholder="Add Cost"/>
</SettingField>