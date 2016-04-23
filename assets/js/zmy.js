$(document).ready(function(){

    map = new AMap.Map('map-container',{
        zoom: 4,
        center: [116.39,39.9]
    });

    $("#map-container").click(panel_close);


    $("#sb-locate-list").click(function(){
        $("#map-panel").load("/admin/location/list", function(){
            $("#map-panel").css("right",0);
            $(".p-d-btn").click(function(){
                var id = $(this).data("id");
                $.get("/admin/location/delete/"+id, function(){
                    $("#sb-locate-list").trigger("click");
                });
            });
        });
    });


    $("#sb-locate").click(function(){
        $.ajax({
            url: "/ajax/location",
            type: "GET",
            dataType: "json",
            success: function(data){
                console.log(data);
                var ll = AMap.LngLat(data.Lng, data.Lat).offset(-500,0);
                map.clearMap();
                map.setZoom(16);
                map.setCenter(ll);

                var marker = new AMap.Marker({
                    position: ll,
                    title: data.CreateAt
                });
                marker.setMap(map);
           
            }
        });
    });


    $("#sb-locate-all").click(function(){
            $.ajax({
            url: "/ajax/location/all",
            type: "GET",
            dataType: "json",
            success: function(data){
                console.log(data);
                map.clearMap();
                for(point in data){
                    if(point == 0){
                        map.setCenter([data[point].Lng, data[point].Lat]);
                        map.setZoom(13);
                    }
                    var marker = new AMap.Marker({
                        position: [data[point].Lng, data[point].Lat],
                        title: data[point].CreateAt
                    });
                    marker.setMap(map);
                }
            
            }
        });
    });


    
      

});

function panel_close(){
    console.log("close");
    $("#map-panel").css("right","-35%");
}