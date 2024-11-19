<script setup lang="ts">
defineProps<{ field?: string; label?: string; span?: string; }>()
</script>

<template>
	<template v-if="field">
		<a-form-item :field="field" :label="label"
					 :rules="[{required:true,message:'请输入'+label}]"
					 :validate-trigger="['blur']">


			<span v-if="span">{{ span }}</span>
			<slot></slot>
		</a-form-item>
	</template>
	<template v-else>
		<a-form-item :label="label">
			<span v-if="span">{{ span }}</span>
			<slot></slot>
		</a-form-item>
	</template>
</template>
