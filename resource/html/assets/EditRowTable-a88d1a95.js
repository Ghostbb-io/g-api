var m=(i,a,n)=>new Promise((l,d)=>{var p=t=>{try{o(n.next(t))}catch(e){d(e)}},s=t=>{try{o(n.throw(t))}catch(e){d(e)}},o=t=>t.done?l(t.value):Promise.resolve(t.value).then(p,s);o((n=n.apply(i,a)).next())});import{B as b}from"./BasicTable-9834a2fb.js";import{T as h}from"./componentMap-c5f962c4.js";import"./TableImg.vue_vue_type_style_index_0_lang-3cd25fdb.js";import{u as R}from"./useTable-b33575de.js";import{o as c}from"./select-93eac05e.js";import{d as v}from"./table-9286518d.js";import{t as g}from"./tree-164e6087.js";import{b as x,_ as I}from"./index.js";import{d as k}from"./antd-12f11a56.js";import{d as T,f as _,a8 as f,_ as w,$ as y,l as A,aa as E,a9 as P,ac as F}from"./vue-5c68ae35.js";import"./BasicForm-33def721.js";import"./FormItem.vue_vue_type_script_lang-bb880c91.js";import"./helper-736f1539.js";import"./BasicForm.vue_vue_type_style_index_0_lang-4cc18aab.js";import"./index-cf0c6833.js";import"./useWindowSizeFn-997fa1d0.js";import"./useForm-61e3194d.js";import"./RadioButtonGroup-81fea792.js";import"./useFormItem-16924efd.js";import"./uuid-31b8b5a4.js";import"./onMountedOrActivated-97a1bb6a.js";import"./download-2647fd8d.js";import"./base64Conver-39fc0d26.js";import"./index-cf464bf9.js";import"./index-7f09b489.js";import"./sortable.esm-4ae27e0b.js";const M=[{title:"输入框",dataIndex:"name-group",editRow:!0,children:[{title:"输入框",dataIndex:"name",editRow:!0,editComponentProps:{prefix:"$"},width:150},{title:"默认输入状态",dataIndex:"name7",editRow:!0,width:150},{title:"输入框校验",dataIndex:"name1",editRow:!0,align:"left",editRule:!0,width:150},{title:"输入框函数校验",dataIndex:"name2",editRow:!0,align:"right",editRule:i=>m(void 0,null,function*(){return i==="2"?"不能输入该值":""})},{title:"数字输入框",dataIndex:"id",editRow:!0,editRule:!0,editComponent:"InputNumber",width:150}]},{title:"下拉框",dataIndex:"name3",editRow:!0,editComponent:"Select",editComponentProps:{options:[{label:"Option1",value:"1"},{label:"Option2",value:"2"},{label:"Option3",value:"3"}]},width:200},{title:"远程下拉",dataIndex:"name4",editRow:!0,editComponent:"ApiSelect",editComponentProps:{api:c,resultField:"list",labelField:"name",valueField:"id"},width:200},{title:"远程下拉树",dataIndex:"name8",editRow:!0,editComponent:"ApiTreeSelect",editRule:!1,editComponentProps:{api:g,resultField:"list"},width:200},{title:"日期选择",dataIndex:"date",editRow:!0,editComponent:"DatePicker",editComponentProps:{valueFormat:"YYYY-MM-DD",format:"YYYY-MM-DD"},width:150},{title:"时间选择",dataIndex:"time",editRow:!0,editComponent:"TimePicker",editComponentProps:{valueFormat:"HH:mm",format:"HH:mm"},width:100},{title:"勾选框",dataIndex:"name5",editRow:!0,editComponent:"Checkbox",editValueMap:i=>i?"是":"否",width:100},{title:"开关",dataIndex:"name6",editRow:!0,editComponent:"Switch",editValueMap:i=>i?"开":"关",width:100},{title:"单选框",dataIndex:"radio1",editRow:!0,editComponent:"RadioGroup",editComponentProps:{options:[{label:"选项1",value:"1"},{label:"选项2",value:"2"}]},width:200},{title:"单选按钮框",dataIndex:"radio2",editRow:!0,editComponent:"RadioButtonGroup",editComponentProps:{options:[{label:"选项1",value:"1"},{label:"选项2",value:"2"}]},width:200},{title:"远程单选框",dataIndex:"radio3",editRow:!0,editComponent:"ApiRadioGroup",editComponentProps:{api:c,resultField:"list",labelField:"name",valueField:"id"},width:200}],B=T({components:{BasicTable:b,TableAction:h},setup(){const{createMessage:i}=x(),a=_(""),[n]=R({title:"可编辑行示例",titleHelpMessage:["本例中修改[数字输入框]这一列时，同一行的[远程下拉]列的当前编辑数据也会同步发生改变"],api:v,columns:M,showIndexColumn:!1,showTableSetting:!0,tableSetting:{fullScreen:!0},actionColumn:{width:160,title:"Action",dataIndex:"action"}});function l(t){var e;a.value=t.key,(e=t.onEdit)==null||e.call(t,!0)}function d(t){var e;a.value="",(e=t.onEdit)==null||e.call(t,!1,!1)}function p(t){return m(this,null,function*(){var r,u;if(i.loading({content:"正在保存...",duration:0,key:"saving"}),yield(r=t.onValid)==null?void 0:r.call(t))try{const C=k(t.editValueRefs);(yield(u=t.onEdit)==null?void 0:u.call(t,!1,!0))&&(a.value=""),i.success({content:"数据已保存",key:"saving"})}catch(C){i.error({content:"保存失败",key:"saving"})}else i.error({content:"请填写正确的数据",key:"saving"})})}function s(t,e){return t.editable?[{label:"保存",onClick:p.bind(null,t,e)},{label:"取消",popConfirm:{title:"是否取消编辑",confirm:d.bind(null,t,e)}}]:[{label:"编辑",disabled:a.value?a.value!==t.key:!1,onClick:l.bind(null,t)}]}function o({column:t,value:e,record:r}){t.dataIndex==="id"&&(r.editValueRefs.name4.value=`${e}`)}return{registerTable:n,handleEdit:l,createActions:s,onEditChange:o}}}),S={class:"p-4"};function Y(i,a,n,l,d,p){const s=f("TableAction"),o=f("BasicTable");return w(),y("div",S,[A(o,{onRegister:i.registerTable,onEditChange:i.onEditChange},{bodyCell:E(({column:t,record:e})=>[t.key==="action"?(w(),P(s,{key:0,actions:i.createActions(e,t)},null,8,["actions"])):F("",!0)]),_:1},8,["onRegister","onEditChange"])])}const rt=I(B,[["render",Y]]);export{rt as default};
