<template>

  <div>

    <el-table
        :data="plates"
        style="width: 100%">
        <el-table-column
          label="标题"
          width="100">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.title }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="位置"
          width="100">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.index }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="800">
          <template slot-scope="scope">
            <el-button type="text" @click="handleDelete(scope.$index, plates)">编辑</el-button>
            <el-button type="text" @click="handleDelete(scope.$index, plates)">删除</el-button>

            <el-upload class="upload-demo" ref="p1Item" action="http://localhost:8090/upload"
               :on-success="onUploadSuccess1" :auto-upload="false">
              <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
              <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload(2)">上传到服务器</el-button>
              <div slot="tip" class="el-upload__tip"></div>
            </el-upload>

          </template>
        </el-table-column>
      </el-table>
      <el-button type="primary" style="width: 100%;margin-top: 5px;" icon="el-icon-plus" @click="addPlate">新增板块</el-button>

    <el-row>
      <el-col :span="14">
        <div class="title">超实惠</div>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="2">
        <div class="title">板块图片：</div>
      </el-col>
      <el-col :span="4">
        <el-upload class="upload-demo" ref="p1" action="http://localhost:8090/upload" :on-success="onP1Success" :auto-upload="false">
          <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
          <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload(1)">上传到服务器</el-button>
          <div slot="tip" class="el-upload__tip"></div>
        </el-upload>
      </el-col>
    </el-row>

    <el-table
        :data="plates1.items"
        style="width: 100%">
        <el-table-column
          label="商家ID"
          width="250">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.storeId }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="推广文案"
          width="500">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.desc }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button
              size="mini"
              type="danger"
              @click="handleDelete(scope.$index, plates1.items)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-row>
        <el-col :span="3">
          <el-button @click="addNew">新增一条推荐</el-button>
        </el-col>
        <el-col :span="4">
          <el-input placeholder="商家ID" v-model="id1" clearable></el-input>
        </el-col>
        <el-col :span="8">
          <el-input placeholder="文案" v-model="desc1" clearable></el-input>
        </el-col>
        <el-col :span="4">
          <el-upload class="upload-demo" ref="p1Item" action="http://localhost:8090/upload"
             :on-success="onUploadSuccess1" :auto-upload="false">
            <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
            <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload(2)">上传到服务器</el-button>
            <div slot="tip" class="el-upload__tip"></div>
          </el-upload>
        </el-col>
      </el-row>

    <el-row>
      <el-col :span="18">
        <el-button type="primary" @click="doUpdate">更新超实惠板块</el-button>
      </el-col>
    </el-row>


    <el-row>

       <el-divider><i class="el-icon-mobile-phone"></i></el-divider>
       <el-divider><i class="el-icon-mobile-phone"></i></el-divider>

    </el-row>
    <!-- 0000000000000 -->

    <el-row>
      <el-col :span="20">
        <div class="title">最新鲜</div>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="2">
        <div class="title">板块图片：</div>
      </el-col>
      <el-col :span="4">
        <el-upload class="upload-demo" ref="p2" action="http://localhost:8090/upload" :on-success="onP2Success" :auto-upload="false">
          <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
          <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload(3)">上传到服务器</el-button>
          <div slot="tip" class="el-upload__tip"></div>
        </el-upload>
      </el-col>
    </el-row>

    <el-table
        :data="plates2.items"
        style="width: 100%">
        <el-table-column
          label="商家ID"
          width="250">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.storeId }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="推广文案"
          width="500">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.desc }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button
              size="mini"
              type="danger"
              @click="handleDelete(scope.$index, plates2.items)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-row>
        <el-col :span="3">
          <el-button @click="addNew2">新增一条推荐</el-button>
        </el-col>
        <el-col :span="4">
          <el-input placeholder="商家ID" v-model="id2" clearable></el-input>
        </el-col>
        <el-col :span="8">
          <el-input placeholder="文案" v-model="desc2" clearable></el-input>
        </el-col>
        <el-col :span="4">
          <el-upload class="upload-demo" ref="p2Item" action="http://localhost:8090/upload"
             :on-success="onUploadSuccess2" :auto-upload="false">
            <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
            <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload(4)">上传到服务器</el-button>
            <div slot="tip" class="el-upload__tip"></div>
          </el-upload>
        </el-col>
      </el-row>

    <el-row>
      <el-col :span="18">
        <el-button type="primary" @click="doUpdate2">更新最新鲜板块</el-button>
      </el-col>
    </el-row>

    <!-- 0000000000000 -->

    <el-row>

       <el-divider><i class="el-icon-mobile-phone"></i></el-divider>
       <el-divider><i class="el-icon-mobile-phone"></i></el-divider>

    </el-row>

    <!-- 0000000000000 -->

    <el-row>
      <el-col :span="20">
        <div class="title">生活丽人</div>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="2">
        <div class="title">板块图片：</div>
      </el-col>
      <el-col :span="4">
        <el-upload class="upload-demo" ref="p3" action="http://localhost:8090/upload" :on-success="onP3Success" :auto-upload="false">
          <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
          <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload(5)">上传到服务器</el-button>
          <div slot="tip" class="el-upload__tip"></div>
        </el-upload>
      </el-col>
    </el-row>

    <el-table
        :data="plates3.items"
        style="width: 100%">
        <el-table-column
          label="商家ID"
          width="250">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.storeId }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="推广文案"
          width="500">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.desc }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button
              size="mini"
              type="danger"
              @click="handleDelete(scope.$index, plates3.items)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-row>
        <el-col :span="3">
          <el-button @click="addNew3">新增一条推荐</el-button>
        </el-col>
        <el-col :span="4">
          <el-input placeholder="商家ID" v-model="id3" clearable></el-input>
        </el-col>
        <el-col :span="8">
          <el-input placeholder="文案" v-model="desc3" clearable></el-input>
        </el-col>
        <el-col :span="4">
          <el-upload class="upload-demo" ref="p3Item" action="http://localhost:8090/upload"
             :on-success="onUploadSuccess3" :auto-upload="false">
            <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
            <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload(6)">上传到服务器</el-button>
            <div slot="tip" class="el-upload__tip"></div>
          </el-upload>
        </el-col>
      </el-row>

    <el-row>
      <el-col :span="18">
        <el-button type="primary" @click="doUpdate3">更新生活丽人板块</el-button>
      </el-col>
    </el-row>

    <!-- 0000000000000 -->


  </div>
</template>

<script>
  import axios from 'axios'
  export default {
    data() {
      return {
        plates: [],
        plates1: {},
        plates2: [],
        plates3: [],

        id1: '',
        desc1: '',
        image1: '',

        id2: '',
        desc2: '',
        image2: '',

        id3: '',
        desc3: '',
        image3: '',
      }
    },
    methods: {
      handleDelete(index, rows) {
        rows.splice(index, 1)
      },
      addPlate() {
        this.plates.push({})
      },
      addNew(){
        if (!this.id1) {
          this.$message.error("请填写商家ID")
          return
        }
        if (!this.desc1) {
          this.$message.error("请填推广文案")
          return
        }
        if (!this.image1) {
          this.$message.error("上传推广图片")
          return
        }
        this.plates1.items.push({storeId:this.id1,desc:this.desc1,image: this.image1})
        this.id1 = ''
        this.desc1 = ''
        this.image1 = ''
      },

      addNew2(){
        if (!this.id2) {
          this.$message.error("请填写商家ID")
          return
        }
        if (!this.desc2) {
          this.$message.error("请填推广文案")
          return
        }
        if (!this.image2) {
          this.$message.error("上传推广图片")
          return
        }
        this.plates2.items.push({storeId:this.id2,desc:this.desc2,image: this.image2})
        this.id2 = ''
        this.desc2 = ''
        this.image2 = ''
      },

      addNew3(){
        if (!this.id3) {
          this.$message.error("请填写商家ID")
          return
        }
        if (!this.desc3) {
          this.$message.error("请填推广文案")
          return
        }
        if (!this.image3) {
          this.$message.error("上传推广图片")
          return
        }
        this.plates3.items.push({storeId:this.id3,desc:this.desc3,image: this.image3})
        this.id3 = ''
        this.desc3 = ''
        this.image3 = ''
      },

      // plate 1
      onP1Success(e){
        this.plates1.image = e
        this.$message.success("上传超实惠板块图片成功")
      },
      onUploadSuccess1(e) {
        this.image1 = e
        this.$message.success("上传超实惠板块图片成功")
      },

      // plate 2
      onP2Success(e){
        this.plates2.image = e
        this.$message.success("上传最新鲜板块图片成功")
      },
      onUploadSuccess2(e) {
        this.image2 = e
        this.$message.success("上传最新鲜板块图片成功")
      },

      // plate 2
      onP3Success(e){
        this.plates3.image = e
        this.$message.success("上传生活丽人板块图片成功")
      },
      onUploadSuccess3(e) {
        this.image3 = e
        this.$message.success("上传生活丽人板块图片成功")
      },

      doUpdate(){
        this.plates1.tp = 2
        axios.post('http://localhost:8090/updateplate',this.plates1).then(resp => {
          if (resp.status != 200 ){
            this.$message.error("更新失败")
          }else {
            this.$message.success("成功")
          }
        })
      },

      doUpdate2(){
        this.plates2.tp = 2
        axios.post('http://localhost:8090/updateplate',this.plates2).then(resp => {
          if (resp.status != 200 ){
            this.$message.error("更新失败")
          }else {
            this.$message.success("成功")
          }
        })
      },

      doUpdate3(){
        this.plates3.tp = 2
        axios.post('http://localhost:8090/updateplate',this.plates3).then(resp => {
          if (resp.status != 200 ){
            this.$message.error("更新失败")
          }else {
            this.$message.success("成功")
          }
        })
      },

      submitUpload(index) {
        switch (index) {
          case 1:
            this.$refs.p1.submit();
            break
          case 2:
            this.$refs.p1Item.submit();
            break
          case 3:
            this.$refs.p2.submit();
            break
          case 4:
            this.$refs.p2Item.submit();
            break
          case 5:
            this.$refs.p3.submit();
            break
          case 6:
            this.$refs.p3Item.submit();
            break
        }
      }
    },
    mounted() {
      axios.get('http://localhost:8090/plates').then(resp => {
        this.plates = resp.data.data
        this.plates1 = resp.data.data[0]
        this.plates2 = resp.data.data[1]
        this.plates3 = resp.data.data[2]
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
