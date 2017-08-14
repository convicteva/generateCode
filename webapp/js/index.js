/**
 * Created by Jamin on 2017/8/10.
 */
$(function(){
    //设置布局宽高
    var width = screen.width-25;
    var height = screen.height-150;
    $("#main").css({"width":width+"px","height":height+"px"});

    //数据源下拉
    $('#node_select').combo({
        required:true,
        editable:false
    });
    $('#node_div').appendTo($('#node_select').combo('panel'));
    //选中数据源
    $('#node_div input').click(function(){
        var v = $(this).val();
        var s = $(this).next('span').text();
        $('#node_select').combo('setValue', v).combo('setText', s).combo('hidePanel');

        //加载数据源对应的表
        loadTable(v)
    });

    //加载数据源对应的表
    function loadTable(node) {
        selectiveNode = node;
        $.get("/getTable/"+node,function (data) {
            $("#table_name_slice_ul").html("");
            for(var i in data){
                var t = $("#template").clone();
                $($(t).find("input")[0]).val(data[i]);
                $($(t).find("a")[0]).text(data[i]);
                $("#table_name_slice_ul").append(t);
            }
        })
    }

});

//选中结点
var selectiveNode="";

//点击表事件
function tab_click() {

}

//导出代码
function exportCode(){
    //1、选中的表名
    var tableNameNodes = $("input[name='tableName']");
    var tableNames = [];
    for(var i in tableNameNodes){
        if(tableNameNodes[i].checked){
            tableNames.push(tableNameNodes[i].value)
        }
    }
    //包名
    var packageName = $("#packageName").val();
    if(!$.trim(packageName)){
        alert("包名不能为空");
        return false;
    }
    //生成代码
    var param = {};
    param["packageName"] = packageName;
    param["node"] = selectiveNode;
    param["tableSlice"] = JSON.stringify(tableNames);

    $.post("/generateCode",param,function (data) {
        console.log(data);
    })

}
    
