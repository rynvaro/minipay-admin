<template>

  <div>
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
        </el-table>
  </div>
</template>

<script>
  import axios from 'axios'
  export default {
    data() {
      return {
        stores:[]
      }
    },
    methods: {

    },
    mounted() {
      axios.get('http://localhost:8090/stores').then(resp => {
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
