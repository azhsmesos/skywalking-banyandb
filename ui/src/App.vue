<!--
  ~ Licensed to Apache Software Foundation (ASF) under one or more contributor
  ~ license agreements. See the NOTICE file distributed with
  ~ this work for additional information regarding copyright
  ~ ownership. Apache Software Foundation (ASF) licenses this file to you under
  ~ the Apache License, Version 2.0 (the "License"); you may
  ~ not use this file except in compliance with the License.
  ~ You may obtain a copy of the License at
  ~
  ~     http://www.apache.org/licenses/LICENSE-2.0
  ~
  ~ Unless required by applicable law or agreed to in writing,
  ~ software distributed under the License is distributed on an
  ~ "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
  ~ KIND, either express or implied.  See the License for the
  ~ specific language governing permissions and limitations
  ~ under the License.
-->

<template>
  <div id="app" @mousedown="appMouseDown">
    <el-container>
      <el-header v-if="show">
        <header-component :active="activePath" :showButton="showButton"></header-component>
      </el-header>
      <el-main>
        <keep-alive>
          <router-view v-if="$route.meta.keepAlive" />
        </keep-alive>
        <router-view v-if="!$route.meta.keepAlive"></router-view>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import HeaderComponent from "./components/HeaderComponent.vue";
export default {
  data() {
    return {
      activePath: "",
      show: true,
      showButton: true,
    }
  },

  components: {
    HeaderComponent
  },

  beforeCreate() {
    this.$loading.create()
  },

  created() {
    let path = this.$route.path
    let name = this.$route.name
    if (name == "NotFound") {
      this.show = false
    } else {
      this.activePath = path
    }
    if (name == "Database") {
      this.$store.commit('changeShowButton', true)
    } else {
      this.$store.commit('changeShowButton', false)
    }
  },

  mounted() {
    this.$loading.close()
  },

  methods: {
    appMouseDown() {
      this.$store.commit("changeShowRightMenu", false)
    }
  }
}
</script>

<style lang="scss">
html,
body,
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  height: 100%;
  margin: 0;
  padding: 0;
}

.el-container {
  height: 100%;
  padding: 0;
  margin: 0;

  .el-header {
    background-color: #ffffff;
    padding: 0;
    margin: 0;
  }

  .el-main {
    background-color: #eaedf1;
    padding: 0;
    margin: 0;
  }
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
