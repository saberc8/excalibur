<template>
  <div class="app">
    <a-layout>
      <a-layout-header>
        <div class="header-content">
          <span class="logo">Excalibur Project</span>
          <a-menu mode="horizontal">
            <a-menu-item key="1">
              <router-link to="/">首页</router-link>
            </a-menu-item>
            <a-menu-item key="2">
              <router-link to="/page2">页面2</router-link>
            </a-menu-item>
          </a-menu>
        </div>
      </a-layout-header>
      <a-layout-content>
        <router-view></router-view>
      </a-layout-content>
    </a-layout>
  </div>
</template>

<script setup>
</script>

<style lang="less">
.app {
  height: 100vh;
  .header-content {
    display: flex;
    align-items: center;
    .logo {
      color: #fff;
      margin-right: 20px;
    }
  }
  .arco-layout-content {
    background: #f5f6f7;
    min-height: calc(100vh - 60px);
  }
}
</style>
