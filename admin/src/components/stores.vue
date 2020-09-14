<template>

  <div>
    <el-row>
      <el-col :span="10">
        <el-input placeholder="搜索商家" v-model="q" @change="search" clearable></el-input>
      </el-col>
    </el-row>
    <el-table
          :data="stores"
          style="width: 100%">
          <el-table-column
            prop="_id"
            label="ID"
            width="280">
          </el-table-column>
          <el-table-column
            prop="storeName"
            label="商家名称"
            width="180">
          </el-table-column>
          <el-table-column
            prop="storeDesc"
            label="描述">
          </el-table-column>
          <el-table-column
            prop="address"
            label="地址">
          </el-table-column>
          <el-table-column
            fixed="right"
            label="操作"
            width="100">
            <template slot-scope="scope">
              <el-button @click="handleClick(scope.row)" type="text" size="small">编辑</el-button>
            </template>
          </el-table-column>
        </el-table>
  </div>
</template>

<script>
  import axios from 'axios'
  export default {
    data() {
      return {
        stores:[],
        q: '',
      }
    },
    methods: {
      handleClick(row) {
        this.$router.push({name: 'pubstores', params: {_id: row._id}})
      },
      search() {
        axios.get('http://localhost:8090/stores?q='+this.q).then(resp => {
          this.stores = resp.data.data
        })
      }
    },
    mounted() {
      axios.get('http://localhost:8090/stores?q='+this.q).then(resp => {
        this.stores = resp.data.data
      })
    }
  }
</script>

<style>
  .el-row {
    margin-bottom: 20px;

    &:last-child {
      margin-bottom: 0;
    }
  }

  .el-col {
    border-radius: 4px;
  }

  .title {
    margin-top: 10px;
  }

</style>
