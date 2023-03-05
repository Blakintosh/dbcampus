<script>
	import SmallButton from "../../../../common/SmallButton.svelte";
	import SettingField from "../components/SettingField.svelte";

</script>
<SettingField title="Rename Project" description="Enter a new name for the project.">
    <input type="text" class=" bg-slate-800 text-slate-100 py-2 px-4 rounded-sm" placeholder="Project Name"/>
</SettingField>

<SettingField title="Delete Project" description="This action is irreversible.">
    <SmallButton type="danger" label="Delete Project"/>
</SettingField>