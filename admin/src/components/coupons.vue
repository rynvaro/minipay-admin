<template>
    <div>
      <el-row>
        <el-col :span="2">
          <div class="title">板块名称：</div>
        </el-col>
        <el-col :span="12">
          <el-input placeholder="请输入店铺名称" v-model="plate.title" clearable></el-input>
        </el-col>
      </el-row>

      <el-row>
        <el-col :span="2">
          <div class="title">跳转连接：</div>
        </el-col>
        <el-col :span="2">
          <el-select v-model="plate.linkTo" filterable placeholder="跳转连接">
              <el-option
                v-for="item in links"
                :key="item.id"
                :label="item.label"
                :value="item.value">
              </el-option>
            </el-select>
        </el-col>
      </el-row>




      <el-row>
        <el-col :span="2">
          <div class="title">板块位置：</div>
        </el-col>
        <el-col :span="12">
          <el-input placeholder="请输入店铺名称" v-model="plate.index" clearable></el-input>
        </el-col>
      </el-row>

      <el-row>
        <el-col :span="2">
          <div class="title">板块图片：</div>
        </el-col>
        <el-col :span="4">
          <el-upload class="upload-demo" ref="plateImage" action="http://localhost:8090/upload" :multiple="true" :limit="6"
            :file-list="plateImageList" :on-success="onUploadPlateImageSuccess" :auto-upload="false">
            <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
            <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload(1)">上传到服务器</el-button>
            <div slot="tip" class="el-upload__tip"></div>
          </el-upload>
        </el-col>
      </el-row>

      <el-row>
        <el-col :span="20">推荐列表</el-col>
      </el-row>

      <el-table
          :data="plate.items"
          style="width: 100%">
          <el-table-column
            label="商家ID"
            width="300">
            <template slot-scope="scope">
              <span style="margin-left: 10px">{{ scope.row.storeId }}</span>
            </template>
          </el-table-column>
          <el-table-column
            label="名称"
            width="200">
            <template slot-scope="scope">
              <span style="margin-left: 10px">{{ scope.row.title }}</span>
            </template>
          </el-table-column>
          <el-table-column
            label="折扣"
            width="100">
            <template slot-scope="scope">
              <span style="margin-left: 10px">{{ scope.row.discount }}</span>
            </template>
          </el-table-column>
          <el-table-column
            label="描述信息"
            width="400">
            <template slot-scope="scope">
              <span style="margin-left: 10px">{{ scope.row.desc }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作">
            <template slot-scope="scope">
              <el-button
                size="mini"
                type="danger"
                @click="handleDelete(scope.$index, plate.items)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>

        <el-dialog title="新增板块" :visible.sync="dialogFormVisible">

            <el-row>
              <el-col :span="8">
                <el-select v-model="store" filterable placeholder="搜索商家">
                    <el-option
                      v-for="item in stores"
                      :key="item._id"
                      :label="item.storeName"
                      :value="item">
                    </el-option>
                  </el-select>
              </el-col>

            </el-row>

          <el-row>
            <el-col :span="8">
               <el-upload class="upload-demo" ref="item" action="http://localhost:8090/upload" :on-success="onItemUploadSuccess" :auto-upload="false">
                 <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
                 <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload(2)">上传到服务器</el-button>
                 <div slot="tip" class="el-upload__tip"></div>
               </el-upload>

            </el-col>
          </el-row>



          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogFormVisible = false">取 消</el-button>
            <el-button type="primary" @click="doAddNewList">确 定</el-button>
          </div>
        </el-dialog>

        <el-button style="width: 100%;margin-top: 5px;" icon="el-icon-plus" @click="addNewList">新增一条推荐</el-button>

        <el-button type="primary" style="margin-top: 15px;" @click="doUpdate">更新本页</el-button>
    </div>
</template>

<script>
  import axios from 'axios'
  export default {
    data() {
      return {
        id: this.$route.params._id,
        q: '',
        plate: {
          items: []
        },
        stores: [],
        store: {},
        dialogFormVisible: false,

        itemImage: '',
        plateImageList: [],

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
          }],
        }
    },
    methods: {
      onSelectChange(e) {
        console.log(e)
      },
      handleDelete(index, rows) {
        rows.splice(index, 1)
      },
      addNewList() {
        axios.get('http://localhost:8090/stores?q='+this.q).then(resp => {
          this.stores = resp.data.data
        })
        this.dialogFormVisible = true
      },
      //
      doAddNewList() {
        if (!this.store._id) {
          this.$message.error("请选择商家")
          return
        }
        if (!this.itemImage) {
          this.$message.error("请上传推广图片")
          return
        }
        this.plate.items.push({storeId: this.store._id,title: this.store.storeName, discount: this.store.discount, image: this.itemImage, desc: this.store.storeDesc})
        this.selectedValue = ''
        this.itemImage = ''
        this.dialogFormVisible =false
      },
      onItemUploadSuccess(e){
        this.itemImage = e
      },
      onUploadPlateImageSuccess(e) {
        this.plate.image = e
      },
      submitUpload(index) {
        switch (index) {
          case 1:
            this.$refs.plateImage.submit();
            break
          case 2:
            this.$refs.item.submit();
            break
        }
      },

      doUpdate(){
        if (!this.id) {
          this.$message.error("请从推荐板块进入此页面")
          return
        }
        this.plate.tp = 2
        axios.post('http://localhost:8090/updateplate',this.plate).then(resp => {
          if (resp.status != 200 ){
            this.$message.error("更新失败")
          }else {
            this.$message.success("成功")
          }
        })
      }
    },
    mounted() {
      if(this.id) {
        axios.get('http://localhost:8090/plate?id='+this.id).then(resp => {
          this.plate = resp.data.data
        })
      }
    }
  }
</script>

<style>
</style>
