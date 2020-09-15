<template>

  <div>
    <el-table
        :data="plates"
        style="width: 100%">
        <el-table-column
          label="标题"
          width="200">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.title }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="位置"
          width="200">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.index }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="500">
          <template slot-scope="scope">
            <el-button type="text" @click="edit(scope.row)">编辑</el-button>
          </template>
        </el-table-column>
      </el-table>

	  <el-dialog title="新增板块" :visible.sync="dialogFormVisible">

      <el-row>
        <el-col :span="2">
          <div class="title">标题：</div>
        </el-col>
        <el-col :span="4">
          <el-input placeholder="请输入店铺名称" v-model="newPlate.title" clearable></el-input>
        </el-col>

        <el-col :span="2">
          <div class="title">位置：</div>
        </el-col>
        <el-col :span="4">
          <el-input placeholder="请输入位置" v-model="newPlate.index" clearable></el-input>
        </el-col>

        <el-col :span="8">

           <el-upload class="upload-demo" ref="newPlate" action="http://localhost:8090/upload" :on-success="onNewPlateImageUploadSuccess" :auto-upload="false">
             <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
             <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload(10)">上传到服务器</el-button>
             <div slot="tip" class="el-upload__tip"></div>
           </el-upload>

        </el-col>
      </el-row>

      <el-row>
        <el-col :span="3">
          <div class="title">跳转连接：</div>
        </el-col>
        <el-col :span="4">
          <el-select v-model="newPlate.linkTo" filterable placeholder="跳转连接">
              <el-option
                v-for="item in links"
                :key="item.id"
                :label="item.label"
                :value="item.value">
              </el-option>
            </el-select>
        </el-col>
      </el-row>



	    <div slot="footer" class="dialog-footer">
	      <el-button @click="dialogFormVisible = false">取 消</el-button>
	      <el-button type="primary" @click="doAddNewPlate">确 定</el-button>
	    </div>
	  </el-dialog>

    <el-button type="primary" style="width: 100%;margin-top: 5px;" icon="el-icon-plus" @click="addPlate">新增板块</el-button>
  </div>
</template>

<script>
  import axios from 'axios'
  import { Loading } from 'element-ui';
  export default {
    data() {
      return {
        plates: [],
        dialogFormVisible: false,

        newPlate: {
          title: '',
          index: -1,
          image: '',
          linkTo: '',
        },

        links: [
          {
            value: '1',
            label: '超实惠'
          }, {
            value: '2',
            label: '最新鲜'
          }, {
            value: '3',
            label: '生活丽人'
          }, {
            value: '4',
            label: '食在上饶'
          }, {
            value: '5',
            label: '玩乐途游'
          },
          {
            value: '6',
            label: '附近商家'
          }]
      }
    },
    methods: {
      doAddNewPlate() {
        if (!this.newPlate.title) {
          this.$message.error("请输入板块标题")
          return
        }
        if (this.newPlate.index == -1) {
          this.$message.error("请输入展示位置")
          return
        }
        if (!this.newPlate.image) {
          this.$message.error("请上传板块图片")
          return
        }
        if (!this.newPlate.linkTo) {
          this.$message.error("请选择跳转连接")
          return
        }

        let loadingInstance = Loading.service({ fullscreen: true });
        this.newPlate.tp = 1
        axios.post('http://localhost:8090/addplate',this.newPlate).then(resp => {
          loadingInstance.close()
          if (resp.status != 200 ){
            this.$message.error("更新失败")
          }else {
            this.dialogFormVisible = false
            console.log(resp.data)
            this.plates.push({_id: resp.data._id, title: resp.data.title, index: resp.data.index, items: resp.data.items})
            this.$message.success("成功")
          }
          this.newPlate.title = ''
          this.newPlate.index = -1
          this.newPlate.image = ''
          this.newPlate.linkTo = ''
        })

      },

      onNewPlateImageUploadSuccess(e) {
        this.newPlate.image = e
      },


      edit(row) {
        this.$router.push({name: 'coupons', params: {_id: row._id}})
      },

      addPlate() {
        this.dialogFormVisible = true
      },

      submitUpload(index) {
        switch (index) {
          case 10:
            this.$refs.newPlate.submit();
        }
      }
    },
    mounted() {
      axios.get('http://localhost:8090/plates').then(resp => {
        this.plates = resp.data.data
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
