<script setup lang="ts">
import { onMounted, ref } from 'vue'

const status = ref('正在检查服务')

onMounted(async () => {
  try {
    const response = await fetch('/healthz')
    const health = response.ok ? await response.json() : null
    status.value = health?.status === 'ok' ? '服务在线' : '服务不可用'
  } catch {
    status.value = '服务不可用'
  }
})
</script>

<template>
  <main>
    <h1>研改</h1>
    <p>课程作业互评与修订平台</p>
    <p role="status">{{ status }}</p>
  </main>
</template>

<style>
body {
  margin: 2rem;
  font-family: 'Instrument Sans', system-ui, sans-serif;
}
</style>
