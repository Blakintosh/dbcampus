<script>
	import SmallButton from "../../../../common/SmallButton.svelte";
	import CurrencyTextField from "../components/CurrencyTextField.svelte";
	import SettingField from "../components/SettingField.svelte";

</script>
<SettingField title="Regular Project Cost" description="This is the expected regular cost of the project that will be deducted every month.">
	<CurrencyTextField placeholder="Monthly Cost"/>
</SettingField>

<SettingField title="Add Unforeseen Cost" description="Deduct an additional, one-off cost from the project budget." >
    <CurrencyTextField placeholder="Add Cost"/>
</SettingField>