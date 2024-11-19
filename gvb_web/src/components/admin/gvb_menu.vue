<template>
	<div class="gvb_menu">
		<a-menu @menu-item-click="clickMenu" @collapse="collapse"
				v-model:selected-keys="selectKey" v-model:open-keys="openKey"
				show-collapse-button>

			<template v-for="item in global.menuList" :key="item.name">
				<a-menu-item :key="item.name" v-if="item.child?.length === 0">
					{{ item.title }}
					<template #icon>
						<component :is="item.icon"></component>
					</template>
				</a-menu-item>

				<a-sub-menu v-if="item.child?.length !==0" :key="item.name">
					<template #icon>
						<component :is="item.icon"></component>
					</template>
					<template #title>{{ item.title }}</template>
					<a-menu-item :key="sub.name" v-for="sub in item.child">
						<component :is="sub.icon"/>
						{{ sub.title }}
					</a-menu-item>
				</a-sub-menu>
			</template>
		</a-menu>
	</div>
</template>
<script setup lang="ts">

import {type Component, watch} from "vue";
import {ref} from "vue";
import {useRouter} from "vue-router"
import {useRoute} from "vue-router"
import {useStore} from "@/stores";
import {global} from "@/global/global";

const store = useStore()
const router = useRouter()
const route = useRoute()

const selectKey = ref([route.name])
const openKey = ref([route.matched[1].name])

// 侧边栏收缩
function collapse(collapsed: boolean) {
	store.setCollapsed(collapsed)
}


function clickMenu(name: string) {
	router.push({name: name,})
}

// 监听标签栏
watch(() => route.name, () => {
	selectKey.value = [route.name]
	openKey.value = [route.matched[1].name]
})
</script>

<style lang="scss">
.gvb_menu {

	.arco-menu {
		position: inherit;
	}

	.arco-menu-collapse-button {
		position: absolute;
		right: 0;
		top: 50%;
		transform: translate(50%, 50%); // 自身50%
		opacity: 0;
		transition: all .3s;
	}
}

aside:hover {
	.gvb_menu {
		.arco-menu-collapse-button {
			opacity: 1;
		}
	}
}

</style>