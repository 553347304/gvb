<template>
	<div>
		<a-modal
			modal-class="gvb_login_modal"
			width="400px" title="用户登录"
			:hide-title="true"
			:visible="props.visible" :footer="false" @cancel="close"
		>
			<Gvb_login_form @ok="close" ref="gvbLoginForm"></Gvb_login_form>
		</a-modal>
	</div>
</template>

<script setup lang="ts">
import Gvb_login_form from "@/components/common/gvb_login_form.vue";
import {ref} from "vue";

const props = defineProps({
	visible: {type: Boolean,}
})
const emits = defineEmits(["update:visible"])
const gvbLoginForm = ref()

function close() {
	emits("update:visible", false)
	gvbLoginForm.value.formReset()
}

</script>

<style scoped lang="scss">
</style>
