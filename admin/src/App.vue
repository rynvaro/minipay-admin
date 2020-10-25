<template>
  <div id="app">

    <el-container>
      <el-header>{{navList[navIndex].navItem}}</el-header>
      <el-container>
        <el-aside width = "150px">
          <el-col :span="2">
              <el-menu :default-active="this.$route.path" router
                 class="el-menu-vertical-demo"
                 background-color="#545c64"
                 text-color="#fff"
                 active-text-color="#ffd04b">
                 <template v-for="(item,i) in navList">
                   <el-menu-item v-if="!item.subItems" :key="i" :index="item.name" @click="itemclick(i)">
                     <template slot="title"><i class="el-icon-s-platform"></i><span> {{ item.navItem }}</span></template>
                   </el-menu-item>
                   <el-submenu v-else>
                     <template slot="title"><i class="el-icon-s-platform"></i><span> {{ item.navItem }}</span></template>
                     <el-menu-item v-for="(item,i) in item.subItems" :key="i" :index="item.name" @click="itemclick(i)">
                       <template slot="title"><i class="el-icon-s-platform"></i><span> {{ item.navItem }}</span></template>
                     </el-menu-item>
                   </el-submenu>
                 </template>

              </el-menu>
            </el-col>

        </el-aside>
        <el-main>
          <router-view />
        </el-main>
      </el-container>
    </el-container>

  </div>
</template>
<!-- ,subItems: [
          {name:'/components/fcashflows', navItem: '财务流水'},
          {name:'/components/forders', navItem: '订单列表'},
          {name:'/components/fcoupons', navItem: '优惠券'},
          {name:'/components/fsubsidy', navItem: '商家补贴'},
          {name:'/components/fusers', navItem: '用户增长'}
        ] -->
<script>
export default {
  data() {
    return {
      navIndex: 0,
      navList: [
        {name:'/components/pubstores',navItem:'商家发布'},
        {name:'/components/stores',navItem:'商家列表'},
        {name:'/components/storegroups',navItem:'商家分组'},
        {name:'/components/plates',navItem:'推荐板块'},
        {name:'/components/coupons',navItem:'板块详情'},
        {name:'/components/users',navItem:'用户列表'},
        {name:'/components/events',navItem:'活动列表'},
        {name:'/components/finance',navItem:'财务板块'},
        {name:'/components/withdraws',navItem:'提现板块'},
        {name:'/components/orders',navItem:'订单列表'},
        {name:'/components/others',navItem:'其他杂项'},
      ]
    }
  },
    methods: {
      itemclick(e){
        this.navIndex = e
      }
    },
    mounted() {
    }
  }
</script>

<style>

  .el-header, .el-footer {
      background-color: #B3C0D1;
      color: #333;
      text-align: center;
      line-height: 60px;
    }

    .el-aside {
      background-color: #545c64;
      color: #333;
      text-align: center;
      line-height: 200px;
    }

    .el-main {
      background-color: #E9EEF3;
      color: #333333;
      text-align: center;
    }

    body > .el-container {
      margin-bottom: 40px;
    }

    .el-container:nth-child(5) .el-aside,
    .el-container:nth-child(6) .el-aside {
      line-height: 260px;
    }

    .el-container:nth-child(7) .el-aside {
      line-height: 320px;
    }
</style>
