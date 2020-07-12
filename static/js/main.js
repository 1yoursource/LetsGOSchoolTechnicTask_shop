"use strict"

$(function() {
    change_header_bg();
    //signin form
    $(".user").click(() => {
        popUp($(".userForm"), $(".pop-up-container"));
    })
    //add product form
    $(".productItem").click(() => {
        popUp($(".productAddForm"), $(".pop-up-container-add-product-2"));
    })
    //change product form
    $(".product-change").click( function () {
        popUp($(".productChangeForm"), $("#"+$(this).data("id")));
    })
    //add basket form
    $(".basketItem").click(() => {
        popUp($(".basketAddForm"), $(".pop-up-container-add-basket-1"));
    })

    //add product to basket form
    $(".product-basket").click( function () {
        let product_id = $(this).closest(".product")[0].id
        popUp($(".productAddToBasketForm"), $(".pop-up-container-add-to-basket-a"));
        $(".productToBasket").click(function (){
            let basket_id = $("option:selected")[0].id
            //var url = "/auth/add_product_to_basket/product_id/"+product_id+"/basket_id/"+basket_id
            $.ajax({
                method: "POST",
                url: "/auth/add_product_to_basket/product_id/"+product_id+"/basket_id/"+basket_id,
                success: function (result) {
                    console.log(result.responseJSON);
                    //window.location.replace(url)
                    location.reload();
                },
                error : function (result) {
                    console.log(result.responseJSON);

                }
            });
        })
    })

    $(".sidebar-menu nav li").click(toggle_menu)

    $( "#spinner" ).spinner({
        max: 50,
        min: 0
    });

    /*ADMIN page handlers*/
    //product add
    $(".product-add").click(function (){
        $.ajax({
            type: "POST",
            url: "/add_product",
            success: function (result) {
                console.log(result.responseJSON);
                location.reload();
            },
            error : function (result) {
                console.log(result.responseJSON);
            }
        });
    })
    //product delete
    $(".product-delete").click(function (){
        let id = $(this).closest(".product")[0].id
        $.ajax({
            method: "POST",
            url: "/admin/delete_product/"+id,
            success: function (result) {
                location.reload();
                console.log("s"+result.responseJSON);
                },
            error : function (result) {
                console.log("f"+result.responseJSON);

            }
        });
    })
    //product change
    $(".product-change-submit").click(function (){
        $.ajax({
            method: $(this).attr("method"),
            url: $(this).attr("action"),
            success: function (result) {
                location.reload();
                console.log("s"+result.responseJSON);
            },
            error : function (result) {
                console.log("f"+result.responseJSON);
            }
        });
    })
    //sort product by category
    $(".admin-sort-button-category").click(function (){
        let category = $(this).closest(".admin-sort-button-category")[0].value
        $.ajax({
            method: "GET",
            url: "/admin/sort_product_by_category/"+category,
            success: function (result) {
                //location.reload();
                console.log("s"+result.responseJSON);
            },
            error : function (result) {
                console.log("f"+result.responseJSON);
            }
        });
    })
    //sort product by price
    $(".admin-sort-button-price").click(function (){
        let boolValue = $(this).closest(".admin-sort-button-price")[0].value
        $.ajax({
            method: "GET",
            url: "/admin/sort_product_by_price/"+boolValue,
            success: function (result) {
                location.reload()
                console.log("s "+result.responseJSON);
                console.log("s "+result.responseText);
            },
            error : function (result) {
                console.log("f"+result.responseJSON);
            }
        });
    })
    //order delete
    $(".order-delete").click(function (){
        let id = $(this).closest(".order")[0].id
        $.ajax({
            type: "POST",
            url: "/delete_order?id="+id,
            success: function (result) {
                console.log(result.responseJSON);
                location.reload();
            },
            error : function (result) {
                console.log(result.responseJSON);

            }
        });
    })
    /*ADMIN page handlers*/
    /*USER Product page handlers*/
    //add basket
    $(".basket-add").click(function (){
        $.ajax({
            type: "POST",
            url: "/auth/add_basket",
            success: function (result) {
                console.log(result.responseJSON);
                location.reload();
            },
            error : function (result) {
                console.log(result.responseJSON);
            }
        });
    })
    //delete basket
    $(".basket-delete").click(function (){
        let id = $(this).closest(".basket")[0].id
        $.ajax({
            type: "POST",
            url: "/auth/delete_basket/"+id,
            success: function (result) {
                location.reload();
                console.log(result.responseJSON);
            },
            error : function (result) {
                console.log(result.responseJSON);
            }
        });
    })
    //view basket
    $(".basket-view").click(function (){
        let id = $(this).closest(".basket")[0].id
        $.ajax({
            type: "GET",
            url: "/auth/basket/"+id,
            success: function (result) {
                console.log(result.responseJSON);
            },
            error : function (result) {
                console.log(result.responseJSON);
            }
        });
    })
    //delete ptoduct from basket
    $(".deleteProductFromBasket").click(function (){
        let id = $(this).closest(".CartItem")[0].id
        $.ajax({
            type: "POST",
            url: "/auth/delete_product_from_basket"+id,
            success: function (result) {
                location.reload();
                console.log(result.responseJSON);
            },
            error : function (result) {
                console.log(result.responseJSON);
            }
        });
    })
    //sort by category
   /* $(".sort-button-category").click(function (){
        let category = $(this).closest(".sort-button-category")[0].value
        $.ajax({
            method: "GET",
            url: "",
            success: function (result) {
                location.reload();
                console.log("s"+result.responseJSON);
            },
            error : function (result) {
                console.log("f"+result.responseJSON);
            }
        });
    })*/
    /*USER Product page handlers*/

    // $(document).mouseup(function(e) { // событие клика по веб-документу
    //     let div = $(".userForm"); // тут указываем ID элемента
    //     if (!div.is(e.target) && div.has(e.target).length === 0 && $(".pop-up-container").hasClass("show__pop-up-container")) { // и не по его дочерним элементам
    //         toggle_popUp(".pop-up-container"); // скрываем его
    //     }
    // })
});

function change_header_bg() {
    setInterval(() => {
        let width = document.documentElement.clientWidth;
        let left = $(".bg-container").css("left");
        if (+left.slice(1, -2) > (width * 2.5)) {
            $(".bg-container").animate({ "left": "0vw" }, 1000)
        } else {
            $(".bg-container").animate({ "left": "-=100vw" }, 500)
        }
    }, 5000)
}

// function toggle_popUp(popUp) {
//     if ($(popUp).hasClass("show__pop-up-container")) {
//         $(popUp).removeClass("show__pop-up-container");
//     } else {
//         $(popUp).addClass("show__pop-up-container");
//     }
// }

$(".userForm");
$(".productAddForm");
$(".productChangeForm");
$(".basketAddForm");
$(".pop-up-container")
$(".pop-up-container-add-product-2")
$(".pop-up-container-change-product")
$(".pop-up-container-add-basket")
$(".pop-up-container-add-to-basket")

function popUp(popUp_content, popUp_container) {
    let elem = $(popUp_content);
    let container = $(popUp_container)

    if (!container.hasClass("show__pop-up-container")) {
        container.addClass("show__pop-up-container");
    }

    $(document).on("mouseup", function em(e) {
        if (!elem.is(e.target) && elem.has(e.target).length === 0 && container.hasClass("show__pop-up-container")) {
            $(document).off("mouseup", em, true);
            container.removeClass("show__pop-up-container");
        }
    })
}

function toggle_menu() {
    let sub_menu = $(this).find(".sub-menu");

    if (sub_menu.hasClass("sub-menu_active")) {
        sub_menu.removeClass("sub-menu_active");
        $(this).removeClass("active_sub_item");
    } else {
        sub_menu.addClass("sub-menu_active");
        $(this).addClass("active_sub_item");
    }
}


//  jqeury UI


$(function(){

    $( "#slider-range" ).slider({
        range: true,
        min: 0,
        max: 500,
        values: [ 0, 200 ],
        slide: function( event, ui ) {
            $( "#price" ).val( "$" + ui.values[ 0 ] + " - $" + ui.values[ 1 ] );
        }
    });
    $( "#price" ).val( "$" + $( "#slider-range" ).slider( "values", 0 ) +
        " - $" + $( "#slider-range" ).slider( "values", 1 ) );

});
