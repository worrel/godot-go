package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <gdnative_api_struct.gen.h>

const extern godot_gdnative_core_api_struct *api;

godot_aabb *new_godot_aabb();
godot_aabb *new_godot_aabb() {
	return malloc(sizeof(godot_aabb));
}
godot_array *new_godot_array();
godot_array *new_godot_array() {
	return malloc(sizeof(godot_array));
}
godot_basis *new_godot_basis();
godot_basis *new_godot_basis() {
	return malloc(sizeof(godot_basis));
}
godot_color *new_godot_color();
godot_color *new_godot_color() {
	return malloc(sizeof(godot_color));
}
godot_dictionary *new_godot_dictionary();
godot_dictionary *new_godot_dictionary() {
	return malloc(sizeof(godot_dictionary));
}
godot_error *new_godot_error();
godot_error *new_godot_error() {
	return malloc(sizeof(godot_error));
}
godot_node_path *new_godot_node_path();
godot_node_path *new_godot_node_path() {
	return malloc(sizeof(godot_node_path));
}
godot_plane *new_godot_plane();
godot_plane *new_godot_plane() {
	return malloc(sizeof(godot_plane));
}
godot_pool_byte_array *new_godot_pool_byte_array();
godot_pool_byte_array *new_godot_pool_byte_array() {
	return malloc(sizeof(godot_pool_byte_array));
}
godot_pool_color_array *new_godot_pool_color_array();
godot_pool_color_array *new_godot_pool_color_array() {
	return malloc(sizeof(godot_pool_color_array));
}
godot_pool_int_array *new_godot_pool_int_array();
godot_pool_int_array *new_godot_pool_int_array() {
	return malloc(sizeof(godot_pool_int_array));
}
godot_pool_real_array *new_godot_pool_real_array();
godot_pool_real_array *new_godot_pool_real_array() {
	return malloc(sizeof(godot_pool_real_array));
}
godot_pool_string_array *new_godot_pool_string_array();
godot_pool_string_array *new_godot_pool_string_array() {
	return malloc(sizeof(godot_pool_string_array));
}
godot_pool_vector2_array *new_godot_pool_vector2_array();
godot_pool_vector2_array *new_godot_pool_vector2_array() {
	return malloc(sizeof(godot_pool_vector2_array));
}
godot_pool_vector3_array *new_godot_pool_vector3_array();
godot_pool_vector3_array *new_godot_pool_vector3_array() {
	return malloc(sizeof(godot_pool_vector3_array));
}
godot_quat *new_godot_quat();
godot_quat *new_godot_quat() {
	return malloc(sizeof(godot_quat));
}
godot_rid *new_godot_rid();
godot_rid *new_godot_rid() {
	return malloc(sizeof(godot_rid));
}
godot_rect2 *new_godot_rect2();
godot_rect2 *new_godot_rect2() {
	return malloc(sizeof(godot_rect2));
}
godot_string *new_godot_string();
godot_string *new_godot_string() {
	return malloc(sizeof(godot_string));
}
godot_transform *new_godot_transform();
godot_transform *new_godot_transform() {
	return malloc(sizeof(godot_transform));
}
godot_transform2d *new_godot_transform2d();
godot_transform2d *new_godot_transform2d() {
	return malloc(sizeof(godot_transform2d));
}
godot_variant *new_godot_variant();
godot_variant *new_godot_variant() {
	return malloc(sizeof(godot_variant));
}
godot_vector2 *new_godot_vector2();
godot_vector2 *new_godot_vector2() {
	return malloc(sizeof(godot_vector2));
}
godot_vector3 *new_godot_vector3();
godot_vector3 *new_godot_vector3() {
	return malloc(sizeof(godot_vector3));
}


int godot_call_int_variant(godot_object *instance, godot_method_bind *mb,godot_variant *arg0);
godot_dictionary *godot_call_dictionary(godot_object *instance, godot_method_bind *mb);
godot_transform2d *godot_call_transform2d(godot_object *instance, godot_method_bind *mb);
godot_variant *godot_call_variant_string_string_variant(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,godot_variant *arg2);
int godot_call_int_string_int_int_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,int arg2,int arg3);
godot_object *godot_call_object_bool(godot_object *instance, godot_method_bind *mb,bool arg0);
godot_variant *godot_call_variant_array_array_int_array(godot_object *instance, godot_method_bind *mb,godot_array *arg0,godot_array *arg1,int arg2,godot_array *arg3);
godot_pool_vector3_array *godot_call_poolvector3array_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1);
bool godot_call_bool_transform2d_vector2_object_transform2d_vector2(godot_object *instance, godot_method_bind *mb,godot_transform2d *arg0,godot_vector2 *arg1,godot_object *arg2,godot_transform2d *arg3,godot_vector2 *arg4);
const char *godot_call_string_string(godot_object *instance, godot_method_bind *mb,char *arg0);
godot_dictionary *godot_call_dictionary_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0);
void godot_call_void_int_bool_bool(godot_object *instance, godot_method_bind *mb,int arg0,bool arg1,bool arg2);
void godot_call_void_dictionary(godot_object *instance, godot_method_bind *mb,godot_dictionary *arg0);
int godot_call_int_string_string(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1);
bool godot_call_bool_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1);
void godot_call_void_object_rect2_bool_color_bool_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_rect2 *arg1,bool arg2,godot_color *arg3,bool arg4,godot_object *arg5);
void godot_call_void_int_int_bool_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,bool arg2,int arg3);
godot_transform2d *godot_call_transform2d_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1);
const char *godot_call_string_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0);
void godot_call_void_poolvector3array_bool_bool(godot_object *instance, godot_method_bind *mb,godot_pool_vector3_array *arg0,bool arg1,bool arg2);
godot_variant *godot_call_variant_string_bool(godot_object *instance, godot_method_bind *mb,char *arg0,bool arg1);
const char *godot_call_string_string_string_bool_int_int(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,bool arg2,int arg3,int arg4);
void godot_call_void_rid_rid_rid_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rid *arg1,godot_rid *arg2,godot_rid *arg3);
void godot_call_void_color(godot_object *instance, godot_method_bind *mb,godot_color *arg0);
float godot_call_float_float(godot_object *instance, godot_method_bind *mb,float arg0);
godot_color *godot_call_color_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1);
void godot_call_void_int_rect2(godot_object *instance, godot_method_bind *mb,int arg0,godot_rect2 *arg1);
void godot_call_void_int_plane(godot_object *instance, godot_method_bind *mb,int arg0,godot_plane *arg1);
void godot_call_void_string_int_vector2(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,godot_vector2 *arg2);
void godot_call_void_transform2d(godot_object *instance, godot_method_bind *mb,godot_transform2d *arg0);
void godot_call_void_string_int_string(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,char *arg2);
void godot_call_void_rid_int_array_array_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,godot_array *arg2,godot_array *arg3,int arg4);
void godot_call_void_rid_object_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_object *arg1,int arg2);
void godot_call_void_rid_int_transform(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,godot_transform *arg2);
godot_rect2 *godot_call_rect2(godot_object *instance, godot_method_bind *mb);
godot_pool_vector2_array *godot_call_poolvector2array_int_float(godot_object *instance, godot_method_bind *mb,int arg0,float arg1);
bool godot_call_bool_object_string_variant_object_string_float_int_int_float(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1,godot_variant *arg2,godot_object *arg3,char *arg4,float arg5,int arg6,int arg7,float arg8);
bool godot_call_bool_object_float_string_variant_variant_variant_variant_variant(godot_object *instance, godot_method_bind *mb,godot_object *arg0,float arg1,char *arg2,godot_variant *arg3,godot_variant *arg4,godot_variant *arg5,godot_variant *arg6,godot_variant *arg7);
void godot_call_void_int_nodepath(godot_object *instance, godot_method_bind *mb,int arg0,godot_node_path *arg1);
void godot_call_void_object_string(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1);
void godot_call_void_vector3_float(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0,float arg1);
void godot_call_void_rid_vector2_vector2_color_float_bool(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_vector2 *arg1,godot_vector2 *arg2,godot_color *arg3,float arg4,bool arg5);
void godot_call_void_nodepath_object_int(godot_object *instance, godot_method_bind *mb,godot_node_path *arg0,godot_object *arg1,int arg2);
void godot_call_void_rid_int_int_float(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,int arg2,float arg3);
godot_plane *godot_call_plane(godot_object *instance, godot_method_bind *mb);
godot_pool_int_array *godot_call_poolintarray_int_float_float(godot_object *instance, godot_method_bind *mb,int arg0,float arg1,float arg2);
void godot_call_void_string_int_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,int arg2);
godot_object *godot_call_object_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1);
bool godot_call_bool_object_nodepath_variant_variant_float_int_int_float(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_node_path *arg1,godot_variant *arg2,godot_variant *arg3,float arg4,int arg5,int arg6,float arg7);
void godot_call_void_rid_poolvector2array_poolcolorarray_poolvector2array_rid_rid_bool(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_pool_vector2_array *arg1,godot_pool_color_array *arg2,godot_pool_vector2_array *arg3,godot_rid *arg4,godot_rid *arg5,bool arg6);
godot_array *godot_call_array_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0);
bool godot_call_bool_int(godot_object *instance, godot_method_bind *mb,int arg0);
void godot_call_void_string_dictionary(godot_object *instance, godot_method_bind *mb,char *arg0,godot_dictionary *arg1);
void godot_call_void_rid_rect2_rid_bool_color_bool_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rect2 *arg1,godot_rid *arg2,bool arg3,godot_color *arg4,bool arg5,godot_rid *arg6);
int godot_call_int_string_int_bool_bool(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,bool arg2,bool arg3);
int godot_call_int_vector2_object(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_object *arg1);
void godot_call_void_poolvector3array(godot_object *instance, godot_method_bind *mb,godot_pool_vector3_array *arg0);
void godot_call_void_string_vector2(godot_object *instance, godot_method_bind *mb,char *arg0,godot_vector2 *arg1);
void godot_call_void_object_object_vector3_vector3_int(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_object *arg1,godot_vector3 *arg2,godot_vector3 *arg3,int arg4);
int godot_call_int_poolbytearray(godot_object *instance, godot_method_bind *mb,godot_pool_byte_array *arg0);
void godot_call_void_int_int_color(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,godot_color *arg2);
bool godot_call_bool_object_string(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1);
godot_vector3 *godot_call_vector3_vector3_vector3_float_int_float(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0,godot_vector3 *arg1,float arg2,int arg3,float arg4);
void godot_call_void_rid_vector2_float_color(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_vector2 *arg1,float arg2,godot_color *arg3);
int godot_call_int_string_object(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1);
bool godot_call_bool_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0);
void godot_call_void_int_poolrealarray(godot_object *instance, godot_method_bind *mb,int arg0,godot_pool_real_array *arg1);
void godot_call_void_rid_rid_transform2d(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rid *arg1,godot_transform2d *arg2);
void godot_call_void_rid_bool(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,bool arg1);
godot_array *godot_call_array_int_float(godot_object *instance, godot_method_bind *mb,int arg0,float arg1);
void godot_call_void_string_poolstringarray(godot_object *instance, godot_method_bind *mb,char *arg0,godot_pool_string_array *arg1);
void godot_call_void_int_variant_bool(godot_object *instance, godot_method_bind *mb,int arg0,godot_variant *arg1,bool arg2);
void godot_call_void_bool_float(godot_object *instance, godot_method_bind *mb,bool arg0,float arg1);
int godot_call_int_int_int_float(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,float arg2);
void godot_call_void_object_object_bool(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_object *arg1,bool arg2);
float godot_call_float_rid_int_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,int arg2);
void godot_call_void_object_string_poolstringarray(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1,godot_pool_string_array *arg2);
void godot_call_void_string_poolbytearray_bool(godot_object *instance, godot_method_bind *mb,char *arg0,godot_pool_byte_array *arg1,bool arg2);
void godot_call_void_int_color_bool(godot_object *instance, godot_method_bind *mb,int arg0,godot_color *arg1,bool arg2);
void godot_call_void_int_object_vector2(godot_object *instance, godot_method_bind *mb,int arg0,godot_object *arg1,godot_vector2 *arg2);
godot_pool_vector3_array *godot_call_poolvector3array_vector3_vector3_bool(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0,godot_vector3 *arg1,bool arg2);
int godot_call_int_bool(godot_object *instance, godot_method_bind *mb,bool arg0);
bool godot_call_bool_vector2_float_object(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,float arg1,godot_object *arg2);
float godot_call_float_int(godot_object *instance, godot_method_bind *mb,int arg0);
void godot_call_void_int_vector3_float(godot_object *instance, godot_method_bind *mb,int arg0,godot_vector3 *arg1,float arg2);
godot_array *godot_call_array_int(godot_object *instance, godot_method_bind *mb,int arg0);
void godot_call_void_vector3_vector3_vector3(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0,godot_vector3 *arg1,godot_vector3 *arg2);
void godot_call_void(godot_object *instance, godot_method_bind *mb);
godot_rid *godot_call_rid_vector2_vector2_vector2_rid_rid(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_vector2 *arg1,godot_vector2 *arg2,godot_rid *arg3,godot_rid *arg4);
void godot_call_void_int_string_string_variant(godot_object *instance, godot_method_bind *mb,int arg0,char *arg1,char *arg2,godot_variant *arg3);
void godot_call_void_rid_int_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,godot_rid *arg2);
godot_array *godot_call_array_nodepath(godot_object *instance, godot_method_bind *mb,godot_node_path *arg0);
float godot_call_float_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1);
int godot_call_int_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1);
void godot_call_void_object_vector2_color_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_vector2 *arg1,godot_color *arg2,godot_object *arg3);
void godot_call_void_int_int_rect2_vector2_float(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,godot_rect2 *arg2,godot_vector2 *arg3,float arg4);
void godot_call_void_rid_rid_rid_rid_int_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rid *arg1,godot_rid *arg2,godot_rid *arg3,int arg4,int arg5);
void godot_call_void_array(godot_object *instance, godot_method_bind *mb,godot_array *arg0);
void godot_call_void_object_rect2(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_rect2 *arg1);
int godot_call_int_int_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,int arg2);
void godot_call_void_object_rect2_vector2(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_rect2 *arg1,godot_vector2 *arg2);
void godot_call_void_int_string_variant(godot_object *instance, godot_method_bind *mb,int arg0,char *arg1,godot_variant *arg2);
godot_rid *godot_call_rid_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0);
godot_array *godot_call_array_string_int_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,int arg2);
bool godot_call_bool_vector3(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0);
godot_rid *godot_call_rid_int_bool(godot_object *instance, godot_method_bind *mb,int arg0,bool arg1);
bool godot_call_bool_string_object_string(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1,char *arg2);
void godot_call_void_rid_string_variant(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,char *arg1,godot_variant *arg2);
int godot_call_int_int_string_poolstringarray_string(godot_object *instance, godot_method_bind *mb,int arg0,char *arg1,godot_pool_string_array *arg2,char *arg3);
godot_vector2 *godot_call_vector2_int_int_object_vector2(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,godot_object *arg2,godot_vector2 *arg3);
godot_vector2 *godot_call_vector2_string(godot_object *instance, godot_method_bind *mb,char *arg0);
void godot_call_void_plane(godot_object *instance, godot_method_bind *mb,godot_plane *arg0);
bool godot_call_bool_vector2_variant_object(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_variant *arg1,godot_object *arg2);
void godot_call_void_int_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,int arg2);
int godot_call_int_nodepath(godot_object *instance, godot_method_bind *mb,godot_node_path *arg0);
godot_pool_vector2_array *godot_call_poolvector2array(godot_object *instance, godot_method_bind *mb);
bool godot_call_bool_string_string(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1);
void godot_call_void_string_string_object_object(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,godot_object *arg2,godot_object *arg3);
godot_vector2 *godot_call_vector2_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1);
godot_object *godot_call_object_int(godot_object *instance, godot_method_bind *mb,int arg0);
void godot_call_void_string_object_string_variant(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1,char *arg2,godot_variant *arg3);
godot_pool_vector2_array *godot_call_poolvector2array_vector2_vector2_bool(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_vector2 *arg1,bool arg2);
godot_rid *godot_call_rid_rid_string(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,char *arg1);
int godot_call_int_string_string_dictionary_array_array(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,godot_dictionary *arg2,godot_array *arg3,godot_array *arg4);
godot_array *godot_call_array_object_vector3(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_vector3 *arg1);
bool godot_call_bool_transform2d_object_transform2d(godot_object *instance, godot_method_bind *mb,godot_transform2d *arg0,godot_object *arg1,godot_transform2d *arg2);
int godot_call_int_int_int_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,int arg2,int arg3);
void godot_call_void_int_array(godot_object *instance, godot_method_bind *mb,int arg0,godot_array *arg1);
godot_node_path *godot_call_nodepath_int_bool(godot_object *instance, godot_method_bind *mb,int arg0,bool arg1);
godot_variant *godot_call_variant_vector2_object(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_object *arg1);
void godot_call_void_poolvector2array_color_float_bool(godot_object *instance, godot_method_bind *mb,godot_pool_vector2_array *arg0,godot_color *arg1,float arg2,bool arg3);
godot_transform2d *godot_call_transform2d_int(godot_object *instance, godot_method_bind *mb,int arg0);
void godot_call_void_int_int_bool_int_poolbytearray(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,bool arg2,int arg3,godot_pool_byte_array *arg4);
void godot_call_void_poolvector2array_poolintarray(godot_object *instance, godot_method_bind *mb,godot_pool_vector2_array *arg0,godot_pool_int_array *arg1);
const char *godot_call_string_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0);
void godot_call_void_string_nodepath_bool(godot_object *instance, godot_method_bind *mb,char *arg0,godot_node_path *arg1,bool arg2);
godot_dictionary *godot_call_dictionary_vector2_vector2_array_int(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_vector2 *arg1,godot_array *arg2,int arg3);
godot_vector2 *godot_call_vector2_vector3(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0);
godot_vector3 *godot_call_vector3_float(godot_object *instance, godot_method_bind *mb,float arg0);
void godot_call_void_int_int_poolstringarray_poolbytearray(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,godot_pool_string_array *arg2,godot_pool_byte_array *arg3);
int godot_call_int_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1);
void godot_call_void_object_object_rect2_vector2(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_object *arg1,godot_rect2 *arg2,godot_vector2 *arg3);
void godot_call_void_int_int_variant(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,godot_variant *arg2);
float godot_call_float_string_string(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1);
bool godot_call_bool_string_object(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1);
godot_dictionary *godot_call_dictionary_vector3_vector3_array_int(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0,godot_vector3 *arg1,godot_array *arg2,int arg3);
void godot_call_void_int_int_transform2d(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,godot_transform2d *arg2);
bool godot_call_bool_object_nodepath_object_nodepath_variant_float_int_int_float(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_node_path *arg1,godot_object *arg2,godot_node_path *arg3,godot_variant *arg4,float arg5,int arg6,int arg7,float arg8);
void godot_call_void_float(godot_object *instance, godot_method_bind *mb,float arg0);
void godot_call_void_object_object_int(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_object *arg1,int arg2);
void godot_call_void_vector2_variant(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_variant *arg1);
void godot_call_void_rid_transform(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_transform *arg1);
void godot_call_void_rid_rect2_bool_color_bool_object(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rect2 *arg1,bool arg2,godot_color *arg3,bool arg4,godot_object *arg5);
void godot_call_void_rid_string(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,char *arg1);
void godot_call_void_poolbytearray(godot_object *instance, godot_method_bind *mb,godot_pool_byte_array *arg0);
int godot_call_int_int_string_poolstringarray_poolbytearray(godot_object *instance, godot_method_bind *mb,int arg0,char *arg1,godot_pool_string_array *arg2,godot_pool_byte_array *arg3);
godot_object *godot_call_object_vector3(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0);
godot_object *godot_call_object_string_bool_bool(godot_object *instance, godot_method_bind *mb,char *arg0,bool arg1,bool arg2);
bool godot_call_bool_bool(godot_object *instance, godot_method_bind *mb,bool arg0);
void godot_call_void_rid_transform2d(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_transform2d *arg1);
godot_array *godot_call_array_poolbytearray(godot_object *instance, godot_method_bind *mb,godot_pool_byte_array *arg0);
void godot_call_void_rid_rect2_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rect2 *arg1,int arg2);
godot_vector3 *godot_call_vector3_int_float(godot_object *instance, godot_method_bind *mb,int arg0,float arg1);
void godot_call_void_poolintarray(godot_object *instance, godot_method_bind *mb,godot_pool_int_array *arg0);
void godot_call_void_string_bool(godot_object *instance, godot_method_bind *mb,char *arg0,bool arg1);
void godot_call_void_int_bool_int(godot_object *instance, godot_method_bind *mb,int arg0,bool arg1,int arg2);
void godot_call_void_rid_poolvector2array_poolcolorarray_poolvector2array_rid_float_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_pool_vector2_array *arg1,godot_pool_color_array *arg2,godot_pool_vector2_array *arg3,godot_rid *arg4,float arg5,godot_rid *arg6);
float godot_call_float(godot_object *instance, godot_method_bind *mb);
godot_aabb *godot_call_aabb(godot_object *instance, godot_method_bind *mb);
godot_pool_int_array *godot_call_poolintarray(godot_object *instance, godot_method_bind *mb);
godot_transform *godot_call_transform_int(godot_object *instance, godot_method_bind *mb,int arg0);
void godot_call_void_int_float_float_bool(godot_object *instance, godot_method_bind *mb,int arg0,float arg1,float arg2,bool arg3);
void godot_call_void_bool_bool_int_int(godot_object *instance, godot_method_bind *mb,bool arg0,bool arg1,int arg2,int arg3);
godot_dictionary *godot_call_dictionary_string(godot_object *instance, godot_method_bind *mb,char *arg0);
int godot_call_int_string(godot_object *instance, godot_method_bind *mb,char *arg0);
godot_vector3 *godot_call_vector3_vector3_vector3_bool(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0,godot_vector3 *arg1,bool arg2);
void godot_call_void_object_int_bool(godot_object *instance, godot_method_bind *mb,godot_object *arg0,int arg1,bool arg2);
void godot_call_void_string_int_int_int_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,int arg2,int arg3,int arg4);
void godot_call_void_object_bool_rid(godot_object *instance, godot_method_bind *mb,godot_object *arg0,bool arg1,godot_rid *arg2);
godot_object *godot_call_object(godot_object *instance, godot_method_bind *mb);
godot_transform2d *godot_call_transform2d_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0);
void godot_call_void_string_object_string(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1,char *arg2);
godot_transform *godot_call_transform_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1);
godot_vector3 *godot_call_vector3_int(godot_object *instance, godot_method_bind *mb,int arg0);
void godot_call_void_int_vector2(godot_object *instance, godot_method_bind *mb,int arg0,godot_vector2 *arg1);
void godot_call_void_vector2_vector2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_vector2 *arg1);
godot_object *godot_call_object_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1);
void godot_call_void_object_vector2_string_color_int(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_vector2 *arg1,char *arg2,godot_color *arg3,int arg4);
int godot_call_int_string_int_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,char *arg2,int arg3);
bool godot_call_bool_string_int_int_int_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,int arg2,int arg3,int arg4);
void godot_call_void_vector2_vector2_vector2_int(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_vector2 *arg1,godot_vector2 *arg2,int arg3);
godot_array *godot_call_array_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1);
void godot_call_void_float_bool(godot_object *instance, godot_method_bind *mb,float arg0,bool arg1);
void godot_call_void_string_int_bool(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,bool arg2);
void godot_call_void_rect2(godot_object *instance, godot_method_bind *mb,godot_rect2 *arg0);
void godot_call_void_rect2_bool(godot_object *instance, godot_method_bind *mb,godot_rect2 *arg0,bool arg1);
godot_variant *godot_call_variant_string_string_array(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,godot_array *arg2);
godot_dictionary *godot_call_dictionary_int(godot_object *instance, godot_method_bind *mb,int arg0);
void godot_call_void_int_object(godot_object *instance, godot_method_bind *mb,int arg0,godot_object *arg1);
void godot_call_void_rid_rect2_rect2_rid_vector2_vector2_int_int_bool_color_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rect2 *arg1,godot_rect2 *arg2,godot_rid *arg3,godot_vector2 *arg4,godot_vector2 *arg5,int arg6,int arg7,bool arg8,godot_color *arg9,godot_rid *arg10);
godot_object *godot_call_object_string_int_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,int arg2);
godot_vector3 *godot_call_vector3_int_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,int arg2);
void godot_call_void_bool_vector2_vector2(godot_object *instance, godot_method_bind *mb,bool arg0,godot_vector2 *arg1,godot_vector2 *arg2);
godot_transform *godot_call_transform(godot_object *instance, godot_method_bind *mb);
void godot_call_void_variant(godot_object *instance, godot_method_bind *mb,godot_variant *arg0);
godot_pool_byte_array *godot_call_poolbytearray_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1);
int godot_call_int_string_poolstringarray_bool_int_string(godot_object *instance, godot_method_bind *mb,char *arg0,godot_pool_string_array *arg1,bool arg2,int arg3,char *arg4);
godot_pool_real_array *godot_call_poolrealarray(godot_object *instance, godot_method_bind *mb);
void godot_call_void_poolstringarray_bool_string_int(godot_object *instance, godot_method_bind *mb,godot_pool_string_array *arg0,bool arg1,char *arg2,int arg3);
void godot_call_void_int_bool_int_color_bool_int_color_object_object(godot_object *instance, godot_method_bind *mb,int arg0,bool arg1,int arg2,godot_color *arg3,bool arg4,int arg5,godot_color *arg6,godot_object *arg7,godot_object *arg8);
int godot_call_int_int_string(godot_object *instance, godot_method_bind *mb,int arg0,char *arg1);
void godot_call_void_vector3(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0);
float godot_call_float_string(godot_object *instance, godot_method_bind *mb,char *arg0);
bool godot_call_bool_nodepath(godot_object *instance, godot_method_bind *mb,godot_node_path *arg0);
void godot_call_void_vector2_variant_object(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_variant *arg1,godot_object *arg2);
void godot_call_void_rid_int_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,int arg2);
godot_variant *godot_call_variant_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1);
void godot_call_void_object_int_vector2(godot_object *instance, godot_method_bind *mb,godot_object *arg0,int arg1,godot_vector2 *arg2);
void godot_call_void_transform2d_vector2(godot_object *instance, godot_method_bind *mb,godot_transform2d *arg0,godot_vector2 *arg1);
godot_variant *godot_call_variant_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1);
void godot_call_void_rid_vector2_string_color_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_vector2 *arg1,char *arg2,godot_color *arg3,int arg4);
void godot_call_void_int_float_bool_bool(godot_object *instance, godot_method_bind *mb,int arg0,float arg1,bool arg2,bool arg3);
float godot_call_float_rid_vector2_int_int_color(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_vector2 *arg1,int arg2,int arg3,godot_color *arg4);
void godot_call_void_string_int_object_vector2(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,godot_object *arg2,godot_vector2 *arg3);
godot_pool_string_array *godot_call_poolstringarray(godot_object *instance, godot_method_bind *mb);
const char *godot_call_string_vector2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0);
void godot_call_void_object_aabb(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_aabb *arg1);
void godot_call_void_string_int_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,char *arg2,int arg3);
void godot_call_void_rid_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rid *arg1);
void godot_call_void_string_string(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1);
int godot_call_int_int_float(godot_object *instance, godot_method_bind *mb,int arg0,float arg1);
godot_variant *godot_call_variant_string_array(godot_object *instance, godot_method_bind *mb,char *arg0,godot_array *arg1);
void godot_call_void_rid_rect2_color(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rect2 *arg1,godot_color *arg2);
godot_transform *godot_call_transform_bool(godot_object *instance, godot_method_bind *mb,bool arg0);
godot_object *godot_call_object_string_string(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1);
godot_variant *godot_call_variant_string(godot_object *instance, godot_method_bind *mb,char *arg0);
void godot_call_void_int_bool_string_string(godot_object *instance, godot_method_bind *mb,int arg0,bool arg1,char *arg2,char *arg3);
void godot_call_void_nodepath_variant(godot_object *instance, godot_method_bind *mb,godot_node_path *arg0,godot_variant *arg1);
const char *godot_call_string_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1);
int godot_call_int_object_transform_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_transform *arg1,godot_object *arg2);
void godot_call_void_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1);
void godot_call_void_rid_string_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,char *arg1,godot_rid *arg2);
void godot_call_void_poolvector3array_object_bool(godot_object *instance, godot_method_bind *mb,godot_pool_vector3_array *arg0,godot_object *arg1,bool arg2);
void godot_call_void_int_int_float_bool(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,float arg2,bool arg3);
godot_vector3 *godot_call_vector3_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0);
godot_transform2d *godot_call_transform2d_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1);
void godot_call_void_rid_bool_rect2(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,bool arg1,godot_rect2 *arg2);
void godot_call_void_int_transform(godot_object *instance, godot_method_bind *mb,int arg0,godot_transform *arg1);
godot_pool_vector3_array *godot_call_poolvector3array(godot_object *instance, godot_method_bind *mb);
void godot_call_void_string_color(godot_object *instance, godot_method_bind *mb,char *arg0,godot_color *arg1);
void godot_call_void_object_float(godot_object *instance, godot_method_bind *mb,godot_object *arg0,float arg1);
godot_aabb *godot_call_aabb_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1);
void godot_call_void_vector2_float_color(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,float arg1,godot_color *arg2);
const char *godot_call_string_dictionary(godot_object *instance, godot_method_bind *mb,godot_dictionary *arg0);
godot_object *godot_call_object_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0);
void godot_call_void_rid_int_variant(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,godot_variant *arg2);
godot_variant *godot_call_variant_nodepath(godot_object *instance, godot_method_bind *mb,godot_node_path *arg0);
void godot_call_void_int_color(godot_object *instance, godot_method_bind *mb,int arg0,godot_color *arg1);
bool godot_call_bool_object_string_variant_variant_float_int_int_float(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1,godot_variant *arg2,godot_variant *arg3,float arg4,int arg5,int arg6,float arg7);
void godot_call_void_string_variant_bool(godot_object *instance, godot_method_bind *mb,char *arg0,godot_variant *arg1,bool arg2);
void godot_call_void_string_variant(godot_object *instance, godot_method_bind *mb,char *arg0,godot_variant *arg1);
void godot_call_void_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0);
const char *godot_call_string_int(godot_object *instance, godot_method_bind *mb,int arg0);
godot_array *godot_call_array_array_int(godot_object *instance, godot_method_bind *mb,godot_array *arg0,int arg1);
void godot_call_void_string_object_bool(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1,bool arg2);
void godot_call_void_int_object_string(godot_object *instance, godot_method_bind *mb,int arg0,godot_object *arg1,char *arg2);
void godot_call_void_rid_poolintarray_poolvector2array_poolcolorarray_poolvector2array_rid_int_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_pool_int_array *arg1,godot_pool_vector2_array *arg2,godot_pool_color_array *arg3,godot_pool_vector2_array *arg4,godot_rid *arg5,int arg6,godot_rid *arg7);
godot_vector2 *godot_call_vector2_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1);
godot_node_path *godot_call_nodepath(godot_object *instance, godot_method_bind *mb);
godot_transform2d *godot_call_transform2d_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0);
void godot_call_void_string_string_object(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,godot_object *arg2);
void godot_call_void_string(godot_object *instance, godot_method_bind *mb,char *arg0);
void godot_call_void_int_float_variant_float(godot_object *instance, godot_method_bind *mb,int arg0,float arg1,godot_variant *arg2,float arg3);
void godot_call_void_int_array_array_int(godot_object *instance, godot_method_bind *mb,int arg0,godot_array *arg1,godot_array *arg2,int arg3);
void godot_call_void_poolvector2array_color_poolvector2array_object_object_bool(godot_object *instance, godot_method_bind *mb,godot_pool_vector2_array *arg0,godot_color *arg1,godot_pool_vector2_array *arg2,godot_object *arg3,godot_object *arg4,bool arg5);
void godot_call_void_string_string_variant(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,godot_variant *arg2);
bool godot_call_bool_object_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_object *arg1);
int godot_call_int_object_bool_string(godot_object *instance, godot_method_bind *mb,godot_object *arg0,bool arg1,char *arg2);
void godot_call_void_int_float_float_float_bool(godot_object *instance, godot_method_bind *mb,int arg0,float arg1,float arg2,float arg3,bool arg4);
godot_node_path *godot_call_nodepath_int(godot_object *instance, godot_method_bind *mb,int arg0);
godot_color *godot_call_color(godot_object *instance, godot_method_bind *mb);
godot_object *godot_call_object_rect2(godot_object *instance, godot_method_bind *mb,godot_rect2 *arg0);
void godot_call_void_object_string_array(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1,godot_array *arg2);
void godot_call_void_int_object_int(godot_object *instance, godot_method_bind *mb,int arg0,godot_object *arg1,int arg2);
godot_color *godot_call_color_string_string(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1);
void godot_call_void_int_int_int_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,int arg2,int arg3,int arg4);
void godot_call_void_object_int_transform(godot_object *instance, godot_method_bind *mb,godot_object *arg0,int arg1,godot_transform *arg2);
void godot_call_void_int_object_transform2d_bool_vector2(godot_object *instance, godot_method_bind *mb,int arg0,godot_object *arg1,godot_transform2d *arg2,bool arg3,godot_vector2 *arg4);
void godot_call_void_int(godot_object *instance, godot_method_bind *mb,int arg0);
void godot_call_void_int_object_bool(godot_object *instance, godot_method_bind *mb,int arg0,godot_object *arg1,bool arg2);
godot_rid *godot_call_rid_object_int(godot_object *instance, godot_method_bind *mb,godot_object *arg0,int arg1);
bool godot_call_bool_rid_transform2d_vector2_float_object(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_transform2d *arg1,godot_vector2 *arg2,float arg3,godot_object *arg4);
godot_vector2 *godot_call_vector2_int(godot_object *instance, godot_method_bind *mb,int arg0);
godot_object *godot_call_object_object_string(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1);
bool godot_call_bool_string_variant(godot_object *instance, godot_method_bind *mb,char *arg0,godot_variant *arg1);
godot_variant *godot_call_variant_transform2d_vector2_object_transform2d_vector2(godot_object *instance, godot_method_bind *mb,godot_transform2d *arg0,godot_vector2 *arg1,godot_object *arg2,godot_transform2d *arg3,godot_vector2 *arg4);
int godot_call_int_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0);
float godot_call_float_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1);
godot_variant *godot_call_variant_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0);
void godot_call_void_int_int_int_bool_bool_bool_vector2(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,int arg2,bool arg3,bool arg4,bool arg5,godot_vector2 *arg6);
void godot_call_void_object_object_string_variant(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_object *arg1,char *arg2,godot_variant *arg3);
void godot_call_void_string_object_variant(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1,godot_variant *arg2);
godot_rect2 *godot_call_rect2_object_int(godot_object *instance, godot_method_bind *mb,godot_object *arg0,int arg1);
godot_variant *godot_call_variant_rid_string(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,char *arg1);
void godot_call_void_int_int_float(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,float arg2);
void godot_call_void_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1);
void godot_call_void_object_rect2_rect2_color_bool_object_bool(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_rect2 *arg1,godot_rect2 *arg2,godot_color *arg3,bool arg4,godot_object *arg5,bool arg6);
godot_variant *godot_call_variant_int(godot_object *instance, godot_method_bind *mb,int arg0);
void godot_call_void_float_color(godot_object *instance, godot_method_bind *mb,float arg0,godot_color *arg1);
godot_rid *godot_call_rid_vector2_rid_rid(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_rid *arg1,godot_rid *arg2);
void godot_call_void_object_string_variant(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1,godot_variant *arg2);
void godot_call_void_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1);
void godot_call_void_string_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,int arg2);
void godot_call_void_string_string_float(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,float arg2);
void godot_call_void_object_int(godot_object *instance, godot_method_bind *mb,godot_object *arg0,int arg1);
godot_variant *godot_call_variant_array(godot_object *instance, godot_method_bind *mb,godot_array *arg0);
void godot_call_void_string_int_object(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,godot_object *arg2);
void godot_call_void_vector2_int_bool_bool_bool(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,int arg1,bool arg2,bool arg3,bool arg4);
bool godot_call_bool_object_nodepath_variant_object_nodepath_float_int_int_float(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_node_path *arg1,godot_variant *arg2,godot_object *arg3,godot_node_path *arg4,float arg5,int arg6,int arg7,float arg8);
godot_vector2 *godot_call_vector2(godot_object *instance, godot_method_bind *mb);
void godot_call_void_rid_vector3_vector3(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_vector3 *arg1,godot_vector3 *arg2);
void godot_call_void_float_float_float_float(godot_object *instance, godot_method_bind *mb,float arg0,float arg1,float arg2,float arg3);
void godot_call_void_object_string_int(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1,int arg2);
godot_vector2 *godot_call_vector2_int_float(godot_object *instance, godot_method_bind *mb,int arg0,float arg1);
godot_object *godot_call_object_transform2d_vector2(godot_object *instance, godot_method_bind *mb,godot_transform2d *arg0,godot_vector2 *arg1);
godot_rid *godot_call_rid(godot_object *instance, godot_method_bind *mb);
void godot_call_void_poolstringarray(godot_object *instance, godot_method_bind *mb,godot_pool_string_array *arg0);
void godot_call_void_rid_vector2_color_bool_object(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_vector2 *arg1,godot_color *arg2,bool arg3,godot_object *arg4);
godot_pool_int_array *godot_call_poolintarray_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1);
godot_vector3 *godot_call_vector3_vector2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0);
void godot_call_void_poolvector3array_poolvector2array_poolcolorarray_poolvector2array_poolvector3array_array(godot_object *instance, godot_method_bind *mb,godot_pool_vector3_array *arg0,godot_pool_vector2_array *arg1,godot_pool_color_array *arg2,godot_pool_vector2_array *arg3,godot_pool_vector3_array *arg4,godot_array *arg5);
godot_pool_int_array *godot_call_poolintarray_string_int_int_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,int arg2,int arg3);
void godot_call_void_int_int_object(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,godot_object *arg2);
void godot_call_void_rid_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rid *arg1,int arg2);
void godot_call_void_int_float(godot_object *instance, godot_method_bind *mb,int arg0,float arg1);
float godot_call_float_vector2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0);
void godot_call_void_string_string_color(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,godot_color *arg2);
void godot_call_void_basis(godot_object *instance, godot_method_bind *mb,godot_basis *arg0);
void godot_call_void_string_object(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1);
void godot_call_void_string_object_int_string_variant(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1,int arg2,char *arg3,godot_variant *arg4);
godot_vector2 *godot_call_vector2_vector2_vector2_float_int_float(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_vector2 *arg1,float arg2,int arg3,float arg4);
void godot_call_void_int_string(godot_object *instance, godot_method_bind *mb,int arg0,char *arg1);
bool godot_call_bool_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1);
int godot_call_int_int_float_vector3_quat_vector3(godot_object *instance, godot_method_bind *mb,int arg0,float arg1,godot_vector3 *arg2,godot_quat *arg3,godot_vector3 *arg4);
bool godot_call_bool_transform2d_vector2(godot_object *instance, godot_method_bind *mb,godot_transform2d *arg0,godot_vector2 *arg1);
void godot_call_void_string_int_int_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,int arg2,int arg3);
void godot_call_void_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0);
void godot_call_void_vector2_bool(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,bool arg1);
void godot_call_void_float_float_float(godot_object *instance, godot_method_bind *mb,float arg0,float arg1,float arg2);
godot_plane *godot_call_plane_int(godot_object *instance, godot_method_bind *mb,int arg0);
godot_transform *godot_call_transform_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0);
godot_pool_byte_array *godot_call_poolbytearray(godot_object *instance, godot_method_bind *mb);
godot_array *godot_call_array_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1);
bool godot_call_bool_string_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,int arg2);
void godot_call_void_poolcolorarray(godot_object *instance, godot_method_bind *mb,godot_pool_color_array *arg0);
void godot_call_void_poolvector2array_int(godot_object *instance, godot_method_bind *mb,godot_pool_vector2_array *arg0,int arg1);
godot_pool_string_array *godot_call_poolstringarray_int(godot_object *instance, godot_method_bind *mb,int arg0);
void godot_call_void_rid_rect2_rect2_color_bool_object_bool(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rect2 *arg1,godot_rect2 *arg2,godot_color *arg3,bool arg4,godot_object *arg5,bool arg6);
void godot_call_void_rid_rid_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rid *arg1,godot_rid *arg2);
const char *godot_call_string(godot_object *instance, godot_method_bind *mb);
void godot_call_void_bool(godot_object *instance, godot_method_bind *mb,bool arg0);
void godot_call_void_int_int_bool(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,bool arg2);
int godot_call_int_object_bool(godot_object *instance, godot_method_bind *mb,godot_object *arg0,bool arg1);
void godot_call_void_int_int_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,int arg2,int arg3);
godot_object *godot_call_object_float(godot_object *instance, godot_method_bind *mb,float arg0);
godot_node_path *godot_call_nodepath_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0);
int godot_call_int_int_string_int(godot_object *instance, godot_method_bind *mb,int arg0,char *arg1,int arg2);
godot_vector3 *godot_call_vector3(godot_object *instance, godot_method_bind *mb);
godot_variant *godot_call_variant_transform2d_object_transform2d(godot_object *instance, godot_method_bind *mb,godot_transform2d *arg0,godot_object *arg1,godot_transform2d *arg2);
godot_object *godot_call_object_vector2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0);
godot_pool_vector2_array *godot_call_poolvector2array_vector2_vector2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_vector2 *arg1);
godot_vector2 *godot_call_vector2_vector2_bool(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,bool arg1);
godot_variant *godot_call_variant_object_string(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1);
bool godot_call_bool_vector2_variant(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_variant *arg1);
bool godot_call_bool_float(godot_object *instance, godot_method_bind *mb,float arg0);
godot_object *godot_call_object_nodepath(godot_object *instance, godot_method_bind *mb,godot_node_path *arg0);
int godot_call_int_string_object_string_array_int(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1,char *arg2,godot_array *arg3,int arg4);
int godot_call_int_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0);
bool godot_call_bool_object_string_object_string_variant_float_int_int_float(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1,godot_object *arg2,char *arg3,godot_variant *arg4,float arg5,int arg6,int arg7,float arg8);
bool godot_call_bool_string_int_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,char *arg2,int arg3);
void godot_call_void_string_string_poolstringarray(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,godot_pool_string_array *arg2);
int godot_call_int_object_int(godot_object *instance, godot_method_bind *mb,godot_object *arg0,int arg1);
godot_vector2 *godot_call_vector2_float(godot_object *instance, godot_method_bind *mb,float arg0);
void godot_call_void_string_float(godot_object *instance, godot_method_bind *mb,char *arg0,float arg1);
void godot_call_void_int_int_poolbytearray(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,godot_pool_byte_array *arg2);
float godot_call_float_object_vector2_string_string_color(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_vector2 *arg1,char *arg2,char *arg3,godot_color *arg4);
void godot_call_void_rid_color(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_color *arg1);
void godot_call_void_vector2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0);
int godot_call_int_string_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,int arg2);
godot_rect2 *godot_call_rect2_int(godot_object *instance, godot_method_bind *mb,int arg0);
bool godot_call_bool_transform_vector3(godot_object *instance, godot_method_bind *mb,godot_transform *arg0,godot_vector3 *arg1);
void godot_call_void_rid_object_string(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_object *arg1,char *arg2);
godot_array *godot_call_array(godot_object *instance, godot_method_bind *mb);
godot_pool_color_array *godot_call_poolcolorarray(godot_object *instance, godot_method_bind *mb);
godot_vector2 *godot_call_vector2_vector2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0);
int godot_call_int_vector2_float_float_int_int(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,float arg1,float arg2,int arg3,int arg4);
void godot_call_void_poolrealarray(godot_object *instance, godot_method_bind *mb,godot_pool_real_array *arg0);
void godot_call_void_object_string_int_int(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1,int arg2,int arg3);
void godot_call_void_vector2_vector2_color_float_bool(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_vector2 *arg1,godot_color *arg2,float arg3,bool arg4);
godot_object *godot_call_object_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1);
godot_vector3 *godot_call_vector3_float_bool(godot_object *instance, godot_method_bind *mb,float arg0,bool arg1);
void godot_call_void_string_float_float_bool(godot_object *instance, godot_method_bind *mb,char *arg0,float arg1,float arg2,bool arg3);
void godot_call_void_rid_object_string_variant(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_object *arg1,char *arg2,godot_variant *arg3);
bool godot_call_bool_string_int_int_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,int arg2,int arg3);
void godot_call_void_rid_aabb(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_aabb *arg1);
godot_color *godot_call_color_float(godot_object *instance, godot_method_bind *mb,float arg0);
void godot_call_void_int_rid_int_int_int(godot_object *instance, godot_method_bind *mb,int arg0,godot_rid *arg1,int arg2,int arg3,int arg4);
void godot_call_void_rid_vector2_vector2(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_vector2 *arg1,godot_vector2 *arg2);
void godot_call_void_rid_int_float(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,float arg2);
godot_rid *godot_call_rid_rid_transform_rid_transform(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_transform *arg1,godot_rid *arg2,godot_transform *arg3);
void godot_call_void_string_object_int(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1,int arg2);
godot_object *godot_call_object_object_int(godot_object *instance, godot_method_bind *mb,godot_object *arg0,int arg1);
void godot_call_void_string_string_color_bool(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,godot_color *arg2,bool arg3);
int godot_call_int_int(godot_object *instance, godot_method_bind *mb,int arg0);
void godot_call_void_rid_float(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,float arg1);
godot_rid *godot_call_rid_rid_vector3_rid_vector3(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_vector3 *arg1,godot_rid *arg2,godot_vector3 *arg3);
void godot_call_void_rid_rect2(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rect2 *arg1);
void godot_call_void_int_object_int_bool_string(godot_object *instance, godot_method_bind *mb,int arg0,godot_object *arg1,int arg2,bool arg3,char *arg4);
void godot_call_void_int_vector3(godot_object *instance, godot_method_bind *mb,int arg0,godot_vector3 *arg1);
bool godot_call_bool_vector2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0);
void godot_call_void_poolvector2array_poolcolorarray_poolvector2array_object_object_bool(godot_object *instance, godot_method_bind *mb,godot_pool_vector2_array *arg0,godot_pool_color_array *arg1,godot_pool_vector2_array *arg2,godot_object *arg3,godot_object *arg4,bool arg5);
godot_array *godot_call_array_string(godot_object *instance, godot_method_bind *mb,char *arg0);
void godot_call_void_int_poolintarray(godot_object *instance, godot_method_bind *mb,int arg0,godot_pool_int_array *arg1);
godot_pool_vector2_array *godot_call_poolvector2array_int(godot_object *instance, godot_method_bind *mb,int arg0);
godot_vector3 *godot_call_vector3_vector3(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0);
void godot_call_void_vector2_float_vector2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,float arg1,godot_vector2 *arg2);
void godot_call_void_int_transform2d(godot_object *instance, godot_method_bind *mb,int arg0,godot_transform2d *arg1);
godot_color *godot_call_color_int(godot_object *instance, godot_method_bind *mb,int arg0);
void godot_call_void_rid_int_transform2d(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,godot_transform2d *arg2);
godot_pool_int_array *godot_call_poolintarray_int(godot_object *instance, godot_method_bind *mb,int arg0);
void godot_call_void_int_variant(godot_object *instance, godot_method_bind *mb,int arg0,godot_variant *arg1);
void godot_call_void_rid_rect2_rid_rect2_color_bool_rid_bool(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rect2 *arg1,godot_rid *arg2,godot_rect2 *arg3,godot_color *arg4,bool arg5,godot_rid *arg6,bool arg7);
godot_variant *godot_call_variant_variant(godot_object *instance, godot_method_bind *mb,godot_variant *arg0);
void godot_call_void_vector3_vector3(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0,godot_vector3 *arg1);
godot_variant *godot_call_variant(godot_object *instance, godot_method_bind *mb);
void godot_call_void_poolvector2array_poolcolorarray_poolvector2array_object_float_object(godot_object *instance, godot_method_bind *mb,godot_pool_vector2_array *arg0,godot_pool_color_array *arg1,godot_pool_vector2_array *arg2,godot_object *arg3,float arg4,godot_object *arg5);
bool godot_call_bool(godot_object *instance, godot_method_bind *mb);
void godot_call_void_poolvector2array(godot_object *instance, godot_method_bind *mb,godot_pool_vector2_array *arg0);
godot_array *godot_call_array_object_int(godot_object *instance, godot_method_bind *mb,godot_object *arg0,int arg1);
const char *godot_call_string_variant(godot_object *instance, godot_method_bind *mb,godot_variant *arg0);
bool godot_call_bool_vector2_rect2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_rect2 *arg1);
int godot_call_int_vector3(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0);
void godot_call_void_object_object_int_bool(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_object *arg1,int arg2,bool arg3);
godot_rid *godot_call_rid_int_int_float(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,float arg2);
int godot_call_int_int_float_bool(godot_object *instance, godot_method_bind *mb,int arg0,float arg1,bool arg2);
void godot_call_void_string_array(godot_object *instance, godot_method_bind *mb,char *arg0,godot_array *arg1);
void godot_call_void_rid_vector2(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_vector2 *arg1);
int godot_call_int(godot_object *instance, godot_method_bind *mb);
void godot_call_void_object_string_bool(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1,bool arg2);
godot_variant *godot_call_variant_int_string(godot_object *instance, godot_method_bind *mb,int arg0,char *arg1);
godot_rid *godot_call_rid_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1);
void godot_call_void_rid_vector3(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_vector3 *arg1);
godot_variant *godot_call_variant_string_string(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1);
godot_variant *godot_call_variant_int_string_string(godot_object *instance, godot_method_bind *mb,int arg0,char *arg1,char *arg2);
godot_object *godot_call_object_string_bool_string(godot_object *instance, godot_method_bind *mb,char *arg0,bool arg1,char *arg2);
void godot_call_void_rect2_color_bool(godot_object *instance, godot_method_bind *mb,godot_rect2 *arg0,godot_color *arg1,bool arg2);
godot_rid *godot_call_rid_int(godot_object *instance, godot_method_bind *mb,int arg0);
bool godot_call_bool_rid_int_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,int arg2);
void godot_call_void_transform(godot_object *instance, godot_method_bind *mb,godot_transform *arg0);
void godot_call_void_rid_rid_transform(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rid *arg1,godot_transform *arg2);
godot_object *godot_call_object_float_bool(godot_object *instance, godot_method_bind *mb,float arg0,bool arg1);
bool godot_call_bool_string_dictionary(godot_object *instance, godot_method_bind *mb,char *arg0,godot_dictionary *arg1);
void godot_call_void_rid_int_int_bool(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,int arg2,bool arg3);
void godot_call_void_int_bool(godot_object *instance, godot_method_bind *mb,int arg0,bool arg1);
void godot_call_void_variant_object(godot_object *instance, godot_method_bind *mb,godot_variant *arg0,godot_object *arg1);
int godot_call_int_object_transform2d_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_transform2d *arg1,godot_object *arg2);
godot_array *godot_call_array_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0);
godot_object *godot_call_object_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0);
void godot_call_void_nodepath(godot_object *instance, godot_method_bind *mb,godot_node_path *arg0);
void godot_call_void_poolstringarray_int(godot_object *instance, godot_method_bind *mb,godot_pool_string_array *arg0,int arg1);
float godot_call_float_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0);
void godot_call_void_int_string_int(godot_object *instance, godot_method_bind *mb,int arg0,char *arg1,int arg2);
void godot_call_void_bool_bool(godot_object *instance, godot_method_bind *mb,bool arg0,bool arg1);
const char *godot_call_string_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1);
godot_vector2 *godot_call_vector2_float_bool(godot_object *instance, godot_method_bind *mb,float arg0,bool arg1);
godot_array *godot_call_array_vector2_int_array_int(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,int arg1,godot_array *arg2,int arg3);
void godot_call_void_rid_variant(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_variant *arg1);
godot_object *godot_call_object_string(godot_object *instance, godot_method_bind *mb,char *arg0);
godot_rid *godot_call_rid_vector2_vector2_rid_rid(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_vector2 *arg1,godot_rid *arg2,godot_rid *arg3);
void godot_call_void_rid_poolvector2array_poolcolorarray_float_bool(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_pool_vector2_array *arg1,godot_pool_color_array *arg2,float arg3,bool arg4);
godot_basis *godot_call_basis(godot_object *instance, godot_method_bind *mb);
void godot_call_void_int_float_float_float(godot_object *instance, godot_method_bind *mb,int arg0,float arg1,float arg2,float arg3);
void godot_call_void_rid_poolvector2array(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_pool_vector2_array *arg1);
void godot_call_void_object_color_bool(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_color *arg1,bool arg2);
int godot_call_int_vector2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0);
bool godot_call_bool_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1);
void godot_call_void_rid_rid_vector2(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rid *arg1,godot_vector2 *arg2);
void godot_call_void_aabb(godot_object *instance, godot_method_bind *mb,godot_aabb *arg0);
godot_pool_vector3_array *godot_call_poolvector3array_int_float(godot_object *instance, godot_method_bind *mb,int arg0,float arg1);
const char *godot_call_string_string_object(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1);
void godot_call_void_poolvector2array_poolcolorarray_float_bool(godot_object *instance, godot_method_bind *mb,godot_pool_vector2_array *arg0,godot_pool_color_array *arg1,float arg2,bool arg3);
godot_pool_string_array *godot_call_poolstringarray_string(godot_object *instance, godot_method_bind *mb,char *arg0);
godot_dictionary *godot_call_dictionary_bool(godot_object *instance, godot_method_bind *mb,bool arg0);
int godot_call_int_vector2_bool(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,bool arg1);
void godot_call_void_string_array_bool(godot_object *instance, godot_method_bind *mb,char *arg0,godot_array *arg1,bool arg2);
void godot_call_void_rid_int_bool(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,bool arg2);
bool godot_call_bool_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0);
void godot_call_void_rid_poolvector2array_bool(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_pool_vector2_array *arg1,bool arg2);
int godot_call_int_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1);
godot_pool_real_array *godot_call_poolrealarray_int(godot_object *instance, godot_method_bind *mb,int arg0);
void godot_call_void_int_poolvector2array(godot_object *instance, godot_method_bind *mb,int arg0,godot_pool_vector2_array *arg1);
void godot_call_void_int_int_vector2_float(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,godot_vector2 *arg2,float arg3);
bool godot_call_bool_string(godot_object *instance, godot_method_bind *mb,char *arg0);
void godot_call_void_object_bool(godot_object *instance, godot_method_bind *mb,godot_object *arg0,bool arg1);
godot_aabb *godot_call_aabb_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0);
void godot_call_void_rid_int_int_int_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,int arg2,int arg3,int arg4);
void godot_call_void_vector3_vector3_vector3_int(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0,godot_vector3 *arg1,godot_vector3 *arg2,int arg3);


int godot_call_int_string_poolstringarray_bool_int_string(godot_object *instance, godot_method_bind *mb,char *arg0,godot_pool_string_array *arg1,bool arg2,int arg3,char *arg4) {
    godot_int ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg4;
    api->godot_string_new(&gArg4);
    api->godot_string_parse_utf8(&gArg4, arg4);
    const void *c_args[] = {&gArg0,arg1,&arg2,&arg3,&gArg4,};
    printf("CGO: godot_call_int_string_poolstringarray_bool_int_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_string_poolstringarray_bool_int_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg4);
    return ret;
}
void godot_call_void_variant(godot_object *instance, godot_method_bind *mb,godot_variant *arg0) {
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_void_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_variant ...godot_method_bind_ptrcall returned\n");

}
godot_pool_byte_array *godot_call_poolbytearray_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1) {
    godot_pool_byte_array *ret = new_godot_pool_byte_array();
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_poolbytearray_rid_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_poolbytearray_rid_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_vector3(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0) {
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_void_vector3 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_vector3 ...godot_method_bind_ptrcall returned\n");

}
godot_pool_real_array *godot_call_poolrealarray(godot_object *instance, godot_method_bind *mb) {
    godot_pool_real_array *ret = new_godot_pool_real_array();
    const void *c_args[] = {};
    printf("CGO: godot_call_poolrealarray calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_poolrealarray ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_poolstringarray_bool_string_int(godot_object *instance, godot_method_bind *mb,godot_pool_string_array *arg0,bool arg1,char *arg2,int arg3) {
    godot_string gArg2;
    api->godot_string_new(&gArg2);
    api->godot_string_parse_utf8(&gArg2, arg2);
    const void *c_args[] = {arg0,&arg1,&gArg2,&arg3,};
    printf("CGO: godot_call_void_poolstringarray_bool_string_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_poolstringarray_bool_string_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg2);

}
void godot_call_void_int_bool_int_color_bool_int_color_object_object(godot_object *instance, godot_method_bind *mb,int arg0,bool arg1,int arg2,godot_color *arg3,bool arg4,int arg5,godot_color *arg6,godot_object *arg7,godot_object *arg8) {
    const void *c_args[] = {&arg0,&arg1,&arg2,arg3,&arg4,&arg5,arg6,arg7,arg8,};
    printf("CGO: godot_call_void_int_bool_int_color_bool_int_color_object_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_bool_int_color_bool_int_color_object_object ...godot_method_bind_ptrcall returned\n");

}
int godot_call_int_int_string(godot_object *instance, godot_method_bind *mb,int arg0,char *arg1) {
    godot_int ret;
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&arg0,&gArg1,};
    printf("CGO: godot_call_int_int_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_int_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);
    return ret;
}
godot_variant *godot_call_variant_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1) {
    godot_variant *ret = new_godot_variant();
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_variant_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_variant_int_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
float godot_call_float_string(godot_object *instance, godot_method_bind *mb,char *arg0) {
    godot_real ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,};
    printf("CGO: godot_call_float_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_float_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
bool godot_call_bool_nodepath(godot_object *instance, godot_method_bind *mb,godot_node_path *arg0) {
    godot_bool ret;
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_bool_nodepath calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_nodepath ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_vector2_variant_object(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_variant *arg1,godot_object *arg2) {
    const void *c_args[] = {arg0,arg1,arg2,};
    printf("CGO: godot_call_void_vector2_variant_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_vector2_variant_object ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rid_int_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,int arg2) {
    const void *c_args[] = {arg0,&arg1,&arg2,};
    printf("CGO: godot_call_void_rid_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_int_int ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rid_vector2_string_color_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_vector2 *arg1,char *arg2,godot_color *arg3,int arg4) {
    godot_string gArg2;
    api->godot_string_new(&gArg2);
    api->godot_string_parse_utf8(&gArg2, arg2);
    const void *c_args[] = {arg0,arg1,&gArg2,arg3,&arg4,};
    printf("CGO: godot_call_void_rid_vector2_string_color_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_vector2_string_color_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg2);

}
void godot_call_void_object_int_vector2(godot_object *instance, godot_method_bind *mb,godot_object *arg0,int arg1,godot_vector2 *arg2) {
    const void *c_args[] = {arg0,&arg1,arg2,};
    printf("CGO: godot_call_void_object_int_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_int_vector2 ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_transform2d_vector2(godot_object *instance, godot_method_bind *mb,godot_transform2d *arg0,godot_vector2 *arg1) {
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_void_transform2d_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_transform2d_vector2 ...godot_method_bind_ptrcall returned\n");

}
godot_variant *godot_call_variant_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1) {
    godot_variant *ret = new_godot_variant();
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_variant_rid_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_variant_rid_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_pool_string_array *godot_call_poolstringarray(godot_object *instance, godot_method_bind *mb) {
    godot_pool_string_array *ret = new_godot_pool_string_array();
    const void *c_args[] = {};
    printf("CGO: godot_call_poolstringarray calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_poolstringarray ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_float_bool_bool(godot_object *instance, godot_method_bind *mb,int arg0,float arg1,bool arg2,bool arg3) {
    const void *c_args[] = {&arg0,&arg1,&arg2,&arg3,};
    printf("CGO: godot_call_void_int_float_bool_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_float_bool_bool ...godot_method_bind_ptrcall returned\n");

}
float godot_call_float_rid_vector2_int_int_color(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_vector2 *arg1,int arg2,int arg3,godot_color *arg4) {
    godot_real ret;
    const void *c_args[] = {arg0,arg1,&arg2,&arg3,arg4,};
    printf("CGO: godot_call_float_rid_vector2_int_int_color calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_float_rid_vector2_int_int_color ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_string_int_object_vector2(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,godot_object *arg2,godot_vector2 *arg3) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,arg2,arg3,};
    printf("CGO: godot_call_void_string_int_object_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_int_object_vector2 ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
void godot_call_void_string_string(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&gArg0,&gArg1,};
    printf("CGO: godot_call_void_string_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg1);

}
const char *godot_call_string_vector2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0) {
    godot_string *ret = new_godot_string();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_string_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_string_vector2 ...godot_method_bind_ptrcall returned\n");
    godot_char_string gcsRet = godot_string_utf8(ret);
    const char *cRet = api->godot_char_string_get_data(&gcsRet);
    api->godot_char_string_destroy(&gcsRet);
    return cRet;

}
void godot_call_void_object_aabb(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_aabb *arg1) {
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_void_object_aabb calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_aabb ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_string_int_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,char *arg2,int arg3) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg2;
    api->godot_string_new(&gArg2);
    api->godot_string_parse_utf8(&gArg2, arg2);
    const void *c_args[] = {&gArg0,&arg1,&gArg2,&arg3,};
    printf("CGO: godot_call_void_string_int_string_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_int_string_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg2);

}
void godot_call_void_rid_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rid *arg1) {
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_void_rid_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_rid ...godot_method_bind_ptrcall returned\n");

}
godot_transform *godot_call_transform_bool(godot_object *instance, godot_method_bind *mb,bool arg0) {
    godot_transform *ret = new_godot_transform();
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_transform_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_transform_bool ...godot_method_bind_ptrcall returned\n");
    return ret;
}
int godot_call_int_int_float(godot_object *instance, godot_method_bind *mb,int arg0,float arg1) {
    godot_int ret;
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_int_int_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_int_float ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_variant *godot_call_variant_string_array(godot_object *instance, godot_method_bind *mb,char *arg0,godot_array *arg1) {
    godot_variant *ret = new_godot_variant();
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,arg1,};
    printf("CGO: godot_call_variant_string_array calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_variant_string_array ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
void godot_call_void_rid_rect2_color(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rect2 *arg1,godot_color *arg2) {
    const void *c_args[] = {arg0,arg1,arg2,};
    printf("CGO: godot_call_void_rid_rect2_color calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_rect2_color ...godot_method_bind_ptrcall returned\n");

}
const char *godot_call_string_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1) {
    godot_string *ret = new_godot_string();
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_string_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_string_int_int ...godot_method_bind_ptrcall returned\n");
    godot_char_string gcsRet = godot_string_utf8(ret);
    const char *cRet = api->godot_char_string_get_data(&gcsRet);
    api->godot_char_string_destroy(&gcsRet);
    return cRet;

}
godot_object *godot_call_object_string_string(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1) {
    godot_object *ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&gArg0,&gArg1,};
    printf("CGO: godot_call_object_string_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_object_string_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg1);
    return ret;
}
godot_variant *godot_call_variant_string(godot_object *instance, godot_method_bind *mb,char *arg0) {
    godot_variant *ret = new_godot_variant();
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,};
    printf("CGO: godot_call_variant_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_variant_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
void godot_call_void_int_bool_string_string(godot_object *instance, godot_method_bind *mb,int arg0,bool arg1,char *arg2,char *arg3) {
    godot_string gArg2;
    api->godot_string_new(&gArg2);
    api->godot_string_parse_utf8(&gArg2, arg2);
    godot_string gArg3;
    api->godot_string_new(&gArg3);
    api->godot_string_parse_utf8(&gArg3, arg3);
    const void *c_args[] = {&arg0,&arg1,&gArg2,&gArg3,};
    printf("CGO: godot_call_void_int_bool_string_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_bool_string_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg2);
    godot_string_destroy(&gArg3);

}
void godot_call_void_nodepath_variant(godot_object *instance, godot_method_bind *mb,godot_node_path *arg0,godot_variant *arg1) {
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_void_nodepath_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_nodepath_variant ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_poolvector3array_object_bool(godot_object *instance, godot_method_bind *mb,godot_pool_vector3_array *arg0,godot_object *arg1,bool arg2) {
    const void *c_args[] = {arg0,arg1,&arg2,};
    printf("CGO: godot_call_void_poolvector3array_object_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_poolvector3array_object_bool ...godot_method_bind_ptrcall returned\n");

}
int godot_call_int_object_transform_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_transform *arg1,godot_object *arg2) {
    godot_int ret;
    const void *c_args[] = {arg0,arg1,arg2,};
    printf("CGO: godot_call_int_object_transform_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_object_transform_object ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1) {
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_void_rid_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_int ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rid_string_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,char *arg1,godot_rid *arg2) {
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {arg0,&gArg1,arg2,};
    printf("CGO: godot_call_void_rid_string_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_string_rid ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);

}
void godot_call_void_int_transform(godot_object *instance, godot_method_bind *mb,int arg0,godot_transform *arg1) {
    const void *c_args[] = {&arg0,arg1,};
    printf("CGO: godot_call_void_int_transform calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_transform ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_int_int_float_bool(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,float arg2,bool arg3) {
    const void *c_args[] = {&arg0,&arg1,&arg2,&arg3,};
    printf("CGO: godot_call_void_int_int_float_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_int_float_bool ...godot_method_bind_ptrcall returned\n");

}
godot_vector3 *godot_call_vector3_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0) {
    godot_vector3 *ret = new_godot_vector3();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_vector3_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector3_rid ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_transform2d *godot_call_transform2d_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1) {
    godot_transform2d *ret = new_godot_transform2d();
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_transform2d_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_transform2d_int_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_rid_bool_rect2(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,bool arg1,godot_rect2 *arg2) {
    const void *c_args[] = {arg0,&arg1,arg2,};
    printf("CGO: godot_call_void_rid_bool_rect2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_bool_rect2 ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_vector2_float_color(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,float arg1,godot_color *arg2) {
    const void *c_args[] = {arg0,&arg1,arg2,};
    printf("CGO: godot_call_void_vector2_float_color calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_vector2_float_color ...godot_method_bind_ptrcall returned\n");

}
godot_pool_vector3_array *godot_call_poolvector3array(godot_object *instance, godot_method_bind *mb) {
    godot_pool_vector3_array *ret = new_godot_pool_vector3_array();
    const void *c_args[] = {};
    printf("CGO: godot_call_poolvector3array calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_poolvector3array ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_string_color(godot_object *instance, godot_method_bind *mb,char *arg0,godot_color *arg1) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,arg1,};
    printf("CGO: godot_call_void_string_color calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_color ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
void godot_call_void_object_float(godot_object *instance, godot_method_bind *mb,godot_object *arg0,float arg1) {
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_void_object_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_float ...godot_method_bind_ptrcall returned\n");

}
godot_aabb *godot_call_aabb_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1) {
    godot_aabb *ret = new_godot_aabb();
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_aabb_rid_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_aabb_rid_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_object *godot_call_object_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0) {
    godot_object *ret;
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_object_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_object_object ...godot_method_bind_ptrcall returned\n");
    return ret;
}
const char *godot_call_string_dictionary(godot_object *instance, godot_method_bind *mb,godot_dictionary *arg0) {
    godot_string *ret = new_godot_string();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_string_dictionary calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_string_dictionary ...godot_method_bind_ptrcall returned\n");
    godot_char_string gcsRet = godot_string_utf8(ret);
    const char *cRet = api->godot_char_string_get_data(&gcsRet);
    api->godot_char_string_destroy(&gcsRet);
    return cRet;

}
godot_variant *godot_call_variant_nodepath(godot_object *instance, godot_method_bind *mb,godot_node_path *arg0) {
    godot_variant *ret = new_godot_variant();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_variant_nodepath calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_variant_nodepath ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_rid_int_variant(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,godot_variant *arg2) {
    const void *c_args[] = {arg0,&arg1,arg2,};
    printf("CGO: godot_call_void_rid_int_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_int_variant ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_string_variant_bool(godot_object *instance, godot_method_bind *mb,char *arg0,godot_variant *arg1,bool arg2) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,arg1,&arg2,};
    printf("CGO: godot_call_void_string_variant_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_variant_bool ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
void godot_call_void_int_color(godot_object *instance, godot_method_bind *mb,int arg0,godot_color *arg1) {
    const void *c_args[] = {&arg0,arg1,};
    printf("CGO: godot_call_void_int_color calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_color ...godot_method_bind_ptrcall returned\n");

}
bool godot_call_bool_object_string_variant_variant_float_int_int_float(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1,godot_variant *arg2,godot_variant *arg3,float arg4,int arg5,int arg6,float arg7) {
    godot_bool ret;
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {arg0,&gArg1,arg2,arg3,&arg4,&arg5,&arg6,&arg7,};
    printf("CGO: godot_call_bool_object_string_variant_variant_float_int_int_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_object_string_variant_variant_float_int_int_float ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);
    return ret;
}
const char *godot_call_string_int(godot_object *instance, godot_method_bind *mb,int arg0) {
    godot_string *ret = new_godot_string();
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_string_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_string_int ...godot_method_bind_ptrcall returned\n");
    godot_char_string gcsRet = godot_string_utf8(ret);
    const char *cRet = api->godot_char_string_get_data(&gcsRet);
    api->godot_char_string_destroy(&gcsRet);
    return cRet;

}
void godot_call_void_string_variant(godot_object *instance, godot_method_bind *mb,char *arg0,godot_variant *arg1) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,arg1,};
    printf("CGO: godot_call_void_string_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_variant ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
void godot_call_void_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0) {
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_void_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid ...godot_method_bind_ptrcall returned\n");

}
godot_vector2 *godot_call_vector2_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1) {
    godot_vector2 *ret = new_godot_vector2();
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_vector2_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector2_int_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_array *godot_call_array_array_int(godot_object *instance, godot_method_bind *mb,godot_array *arg0,int arg1) {
    godot_array *ret = new_godot_array();
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_array_array_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_array_array_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_string_object_bool(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1,bool arg2) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,arg1,&arg2,};
    printf("CGO: godot_call_void_string_object_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_object_bool ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
void godot_call_void_int_object_string(godot_object *instance, godot_method_bind *mb,int arg0,godot_object *arg1,char *arg2) {
    godot_string gArg2;
    api->godot_string_new(&gArg2);
    api->godot_string_parse_utf8(&gArg2, arg2);
    const void *c_args[] = {&arg0,arg1,&gArg2,};
    printf("CGO: godot_call_void_int_object_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_object_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg2);

}
void godot_call_void_rid_poolintarray_poolvector2array_poolcolorarray_poolvector2array_rid_int_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_pool_int_array *arg1,godot_pool_vector2_array *arg2,godot_pool_color_array *arg3,godot_pool_vector2_array *arg4,godot_rid *arg5,int arg6,godot_rid *arg7) {
    const void *c_args[] = {arg0,arg1,arg2,arg3,arg4,arg5,&arg6,arg7,};
    printf("CGO: godot_call_void_rid_poolintarray_poolvector2array_poolcolorarray_poolvector2array_rid_int_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_poolintarray_poolvector2array_poolcolorarray_poolvector2array_rid_int_rid ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_string(godot_object *instance, godot_method_bind *mb,char *arg0) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,};
    printf("CGO: godot_call_void_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
godot_node_path *godot_call_nodepath(godot_object *instance, godot_method_bind *mb) {
    godot_node_path *ret = new_godot_node_path();
    const void *c_args[] = {};
    printf("CGO: godot_call_nodepath calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_nodepath ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_transform2d *godot_call_transform2d_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0) {
    godot_transform2d *ret = new_godot_transform2d();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_transform2d_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_transform2d_rid ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_string_string_object(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,godot_object *arg2) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&gArg0,&gArg1,arg2,};
    printf("CGO: godot_call_void_string_string_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_string_object ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg1);

}
int godot_call_int_object_bool_string(godot_object *instance, godot_method_bind *mb,godot_object *arg0,bool arg1,char *arg2) {
    godot_int ret;
    godot_string gArg2;
    api->godot_string_new(&gArg2);
    api->godot_string_parse_utf8(&gArg2, arg2);
    const void *c_args[] = {arg0,&arg1,&gArg2,};
    printf("CGO: godot_call_int_object_bool_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_object_bool_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg2);
    return ret;
}
void godot_call_void_int_float_float_float_bool(godot_object *instance, godot_method_bind *mb,int arg0,float arg1,float arg2,float arg3,bool arg4) {
    const void *c_args[] = {&arg0,&arg1,&arg2,&arg3,&arg4,};
    printf("CGO: godot_call_void_int_float_float_float_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_float_float_float_bool ...godot_method_bind_ptrcall returned\n");

}
godot_node_path *godot_call_nodepath_int(godot_object *instance, godot_method_bind *mb,int arg0) {
    godot_node_path *ret = new_godot_node_path();
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_nodepath_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_nodepath_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_float_variant_float(godot_object *instance, godot_method_bind *mb,int arg0,float arg1,godot_variant *arg2,float arg3) {
    const void *c_args[] = {&arg0,&arg1,arg2,&arg3,};
    printf("CGO: godot_call_void_int_float_variant_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_float_variant_float ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_int_array_array_int(godot_object *instance, godot_method_bind *mb,int arg0,godot_array *arg1,godot_array *arg2,int arg3) {
    const void *c_args[] = {&arg0,arg1,arg2,&arg3,};
    printf("CGO: godot_call_void_int_array_array_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_array_array_int ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_poolvector2array_color_poolvector2array_object_object_bool(godot_object *instance, godot_method_bind *mb,godot_pool_vector2_array *arg0,godot_color *arg1,godot_pool_vector2_array *arg2,godot_object *arg3,godot_object *arg4,bool arg5) {
    const void *c_args[] = {arg0,arg1,arg2,arg3,arg4,&arg5,};
    printf("CGO: godot_call_void_poolvector2array_color_poolvector2array_object_object_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_poolvector2array_color_poolvector2array_object_object_bool ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_string_string_variant(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,godot_variant *arg2) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&gArg0,&gArg1,arg2,};
    printf("CGO: godot_call_void_string_string_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_string_variant ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg1);

}
bool godot_call_bool_object_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_object *arg1) {
    godot_bool ret;
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_bool_object_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_object_object ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_object_int(godot_object *instance, godot_method_bind *mb,int arg0,godot_object *arg1,int arg2) {
    const void *c_args[] = {&arg0,arg1,&arg2,};
    printf("CGO: godot_call_void_int_object_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_object_int ...godot_method_bind_ptrcall returned\n");

}
godot_color *godot_call_color(godot_object *instance, godot_method_bind *mb) {
    godot_color *ret = new_godot_color();
    const void *c_args[] = {};
    printf("CGO: godot_call_color calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_color ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_object *godot_call_object_rect2(godot_object *instance, godot_method_bind *mb,godot_rect2 *arg0) {
    godot_object *ret;
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_object_rect2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_object_rect2 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_object_string_array(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1,godot_array *arg2) {
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {arg0,&gArg1,arg2,};
    printf("CGO: godot_call_void_object_string_array calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_string_array ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);

}
void godot_call_void_int(godot_object *instance, godot_method_bind *mb,int arg0) {
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_void_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int ...godot_method_bind_ptrcall returned\n");

}
godot_color *godot_call_color_string_string(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1) {
    godot_color *ret = new_godot_color();
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&gArg0,&gArg1,};
    printf("CGO: godot_call_color_string_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_color_string_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg1);
    return ret;
}
void godot_call_void_int_int_int_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,int arg2,int arg3,int arg4) {
    const void *c_args[] = {&arg0,&arg1,&arg2,&arg3,&arg4,};
    printf("CGO: godot_call_void_int_int_int_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_int_int_int_int ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_object_int_transform(godot_object *instance, godot_method_bind *mb,godot_object *arg0,int arg1,godot_transform *arg2) {
    const void *c_args[] = {arg0,&arg1,arg2,};
    printf("CGO: godot_call_void_object_int_transform calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_int_transform ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_int_object_transform2d_bool_vector2(godot_object *instance, godot_method_bind *mb,int arg0,godot_object *arg1,godot_transform2d *arg2,bool arg3,godot_vector2 *arg4) {
    const void *c_args[] = {&arg0,arg1,arg2,&arg3,arg4,};
    printf("CGO: godot_call_void_int_object_transform2d_bool_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_object_transform2d_bool_vector2 ...godot_method_bind_ptrcall returned\n");

}
bool godot_call_bool_rid_transform2d_vector2_float_object(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_transform2d *arg1,godot_vector2 *arg2,float arg3,godot_object *arg4) {
    godot_bool ret;
    const void *c_args[] = {arg0,arg1,arg2,&arg3,arg4,};
    printf("CGO: godot_call_bool_rid_transform2d_vector2_float_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_rid_transform2d_vector2_float_object ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_object_bool(godot_object *instance, godot_method_bind *mb,int arg0,godot_object *arg1,bool arg2) {
    const void *c_args[] = {&arg0,arg1,&arg2,};
    printf("CGO: godot_call_void_int_object_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_object_bool ...godot_method_bind_ptrcall returned\n");

}
godot_rid *godot_call_rid_object_int(godot_object *instance, godot_method_bind *mb,godot_object *arg0,int arg1) {
    godot_rid *ret = new_godot_rid();
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_rid_object_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_rid_object_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_vector2 *godot_call_vector2_int(godot_object *instance, godot_method_bind *mb,int arg0) {
    godot_vector2 *ret = new_godot_vector2();
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_vector2_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector2_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
int godot_call_int_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0) {
    godot_int ret;
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_int_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_object ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_object *godot_call_object_object_string(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1) {
    godot_object *ret;
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {arg0,&gArg1,};
    printf("CGO: godot_call_object_object_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_object_object_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);
    return ret;
}
bool godot_call_bool_string_variant(godot_object *instance, godot_method_bind *mb,char *arg0,godot_variant *arg1) {
    godot_bool ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,arg1,};
    printf("CGO: godot_call_bool_string_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_string_variant ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
godot_variant *godot_call_variant_transform2d_vector2_object_transform2d_vector2(godot_object *instance, godot_method_bind *mb,godot_transform2d *arg0,godot_vector2 *arg1,godot_object *arg2,godot_transform2d *arg3,godot_vector2 *arg4) {
    godot_variant *ret = new_godot_variant();
    const void *c_args[] = {arg0,arg1,arg2,arg3,arg4,};
    printf("CGO: godot_call_variant_transform2d_vector2_object_transform2d_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_variant_transform2d_vector2_object_transform2d_vector2 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
float godot_call_float_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1) {
    godot_real ret;
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_float_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_float_int_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_object_object_string_variant(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_object *arg1,char *arg2,godot_variant *arg3) {
    godot_string gArg2;
    api->godot_string_new(&gArg2);
    api->godot_string_parse_utf8(&gArg2, arg2);
    const void *c_args[] = {arg0,arg1,&gArg2,arg3,};
    printf("CGO: godot_call_void_object_object_string_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_object_string_variant ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg2);

}
godot_variant *godot_call_variant_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0) {
    godot_variant *ret = new_godot_variant();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_variant_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_variant_rid ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_int_int_bool_bool_bool_vector2(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,int arg2,bool arg3,bool arg4,bool arg5,godot_vector2 *arg6) {
    const void *c_args[] = {&arg0,&arg1,&arg2,&arg3,&arg4,&arg5,arg6,};
    printf("CGO: godot_call_void_int_int_int_bool_bool_bool_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_int_int_bool_bool_bool_vector2 ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_int_int_float(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,float arg2) {
    const void *c_args[] = {&arg0,&arg1,&arg2,};
    printf("CGO: godot_call_void_int_int_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_int_float ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_string_object_variant(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1,godot_variant *arg2) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,arg1,arg2,};
    printf("CGO: godot_call_void_string_object_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_object_variant ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
godot_rect2 *godot_call_rect2_object_int(godot_object *instance, godot_method_bind *mb,godot_object *arg0,int arg1) {
    godot_rect2 *ret = new_godot_rect2();
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_rect2_object_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_rect2_object_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_variant *godot_call_variant_rid_string(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,char *arg1) {
    godot_variant *ret = new_godot_variant();
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {arg0,&gArg1,};
    printf("CGO: godot_call_variant_rid_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_variant_rid_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);
    return ret;
}
void godot_call_void_object_string_variant(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1,godot_variant *arg2) {
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {arg0,&gArg1,arg2,};
    printf("CGO: godot_call_void_object_string_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_string_variant ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);

}
void godot_call_void_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1) {
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_void_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_int ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,};
    printf("CGO: godot_call_void_string_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
void godot_call_void_object_rect2_rect2_color_bool_object_bool(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_rect2 *arg1,godot_rect2 *arg2,godot_color *arg3,bool arg4,godot_object *arg5,bool arg6) {
    const void *c_args[] = {arg0,arg1,arg2,arg3,&arg4,arg5,&arg6,};
    printf("CGO: godot_call_void_object_rect2_rect2_color_bool_object_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_rect2_rect2_color_bool_object_bool ...godot_method_bind_ptrcall returned\n");

}
godot_variant *godot_call_variant_int(godot_object *instance, godot_method_bind *mb,int arg0) {
    godot_variant *ret = new_godot_variant();
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_variant_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_variant_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_float_color(godot_object *instance, godot_method_bind *mb,float arg0,godot_color *arg1) {
    const void *c_args[] = {&arg0,arg1,};
    printf("CGO: godot_call_void_float_color calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_float_color ...godot_method_bind_ptrcall returned\n");

}
godot_rid *godot_call_rid_vector2_rid_rid(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_rid *arg1,godot_rid *arg2) {
    godot_rid *ret = new_godot_rid();
    const void *c_args[] = {arg0,arg1,arg2,};
    printf("CGO: godot_call_rid_vector2_rid_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_rid_vector2_rid_rid ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_string_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,int arg2) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&gArg0,&gArg1,&arg2,};
    printf("CGO: godot_call_void_string_string_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_string_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg1);

}
bool godot_call_bool_object_nodepath_variant_object_nodepath_float_int_int_float(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_node_path *arg1,godot_variant *arg2,godot_object *arg3,godot_node_path *arg4,float arg5,int arg6,int arg7,float arg8) {
    godot_bool ret;
    const void *c_args[] = {arg0,arg1,arg2,arg3,arg4,&arg5,&arg6,&arg7,&arg8,};
    printf("CGO: godot_call_bool_object_nodepath_variant_object_nodepath_float_int_int_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_object_nodepath_variant_object_nodepath_float_int_int_float ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_vector2 *godot_call_vector2(godot_object *instance, godot_method_bind *mb) {
    godot_vector2 *ret = new_godot_vector2();
    const void *c_args[] = {};
    printf("CGO: godot_call_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector2 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_string_string_float(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,float arg2) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&gArg0,&gArg1,&arg2,};
    printf("CGO: godot_call_void_string_string_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_string_float ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg1);

}
void godot_call_void_object_int(godot_object *instance, godot_method_bind *mb,godot_object *arg0,int arg1) {
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_void_object_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_int ...godot_method_bind_ptrcall returned\n");

}
godot_variant *godot_call_variant_array(godot_object *instance, godot_method_bind *mb,godot_array *arg0) {
    godot_variant *ret = new_godot_variant();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_variant_array calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_variant_array ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_string_int_object(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,godot_object *arg2) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,arg2,};
    printf("CGO: godot_call_void_string_int_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_int_object ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
void godot_call_void_vector2_int_bool_bool_bool(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,int arg1,bool arg2,bool arg3,bool arg4) {
    const void *c_args[] = {arg0,&arg1,&arg2,&arg3,&arg4,};
    printf("CGO: godot_call_void_vector2_int_bool_bool_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_vector2_int_bool_bool_bool ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_object_string_int(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1,int arg2) {
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {arg0,&gArg1,&arg2,};
    printf("CGO: godot_call_void_object_string_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_string_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);

}
void godot_call_void_rid_vector3_vector3(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_vector3 *arg1,godot_vector3 *arg2) {
    const void *c_args[] = {arg0,arg1,arg2,};
    printf("CGO: godot_call_void_rid_vector3_vector3 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_vector3_vector3 ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_float_float_float_float(godot_object *instance, godot_method_bind *mb,float arg0,float arg1,float arg2,float arg3) {
    const void *c_args[] = {&arg0,&arg1,&arg2,&arg3,};
    printf("CGO: godot_call_void_float_float_float_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_float_float_float_float ...godot_method_bind_ptrcall returned\n");

}
godot_rid *godot_call_rid(godot_object *instance, godot_method_bind *mb) {
    godot_rid *ret = new_godot_rid();
    const void *c_args[] = {};
    printf("CGO: godot_call_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_rid ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_vector2 *godot_call_vector2_int_float(godot_object *instance, godot_method_bind *mb,int arg0,float arg1) {
    godot_vector2 *ret = new_godot_vector2();
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_vector2_int_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector2_int_float ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_object *godot_call_object_transform2d_vector2(godot_object *instance, godot_method_bind *mb,godot_transform2d *arg0,godot_vector2 *arg1) {
    godot_object *ret;
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_object_transform2d_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_object_transform2d_vector2 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_pool_int_array *godot_call_poolintarray_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1) {
    godot_pool_int_array *ret = new_godot_pool_int_array();
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_poolintarray_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_poolintarray_int_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_poolstringarray(godot_object *instance, godot_method_bind *mb,godot_pool_string_array *arg0) {
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_void_poolstringarray calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_poolstringarray ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rid_vector2_color_bool_object(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_vector2 *arg1,godot_color *arg2,bool arg3,godot_object *arg4) {
    const void *c_args[] = {arg0,arg1,arg2,&arg3,arg4,};
    printf("CGO: godot_call_void_rid_vector2_color_bool_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_vector2_color_bool_object ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_int_float(godot_object *instance, godot_method_bind *mb,int arg0,float arg1) {
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_void_int_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_float ...godot_method_bind_ptrcall returned\n");

}
godot_vector3 *godot_call_vector3_vector2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0) {
    godot_vector3 *ret = new_godot_vector3();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_vector3_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector3_vector2 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_poolvector3array_poolvector2array_poolcolorarray_poolvector2array_poolvector3array_array(godot_object *instance, godot_method_bind *mb,godot_pool_vector3_array *arg0,godot_pool_vector2_array *arg1,godot_pool_color_array *arg2,godot_pool_vector2_array *arg3,godot_pool_vector3_array *arg4,godot_array *arg5) {
    const void *c_args[] = {arg0,arg1,arg2,arg3,arg4,arg5,};
    printf("CGO: godot_call_void_poolvector3array_poolvector2array_poolcolorarray_poolvector2array_poolvector3array_array calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_poolvector3array_poolvector2array_poolcolorarray_poolvector2array_poolvector3array_array ...godot_method_bind_ptrcall returned\n");

}
godot_pool_int_array *godot_call_poolintarray_string_int_int_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,int arg2,int arg3) {
    godot_pool_int_array *ret = new_godot_pool_int_array();
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,&arg2,&arg3,};
    printf("CGO: godot_call_poolintarray_string_int_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_poolintarray_string_int_int_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
void godot_call_void_int_int_object(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,godot_object *arg2) {
    const void *c_args[] = {&arg0,&arg1,arg2,};
    printf("CGO: godot_call_void_int_int_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_int_object ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rid_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rid *arg1,int arg2) {
    const void *c_args[] = {arg0,arg1,&arg2,};
    printf("CGO: godot_call_void_rid_rid_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_rid_int ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_basis(godot_object *instance, godot_method_bind *mb,godot_basis *arg0) {
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_void_basis calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_basis ...godot_method_bind_ptrcall returned\n");

}
float godot_call_float_vector2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0) {
    godot_real ret;
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_float_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_float_vector2 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_string_string_color(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,godot_color *arg2) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&gArg0,&gArg1,arg2,};
    printf("CGO: godot_call_void_string_string_color calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_string_color ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg1);

}
void godot_call_void_int_string(godot_object *instance, godot_method_bind *mb,int arg0,char *arg1) {
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&arg0,&gArg1,};
    printf("CGO: godot_call_void_int_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);

}
void godot_call_void_string_object(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,arg1,};
    printf("CGO: godot_call_void_string_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_object ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
void godot_call_void_string_object_int_string_variant(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1,int arg2,char *arg3,godot_variant *arg4) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg3;
    api->godot_string_new(&gArg3);
    api->godot_string_parse_utf8(&gArg3, arg3);
    const void *c_args[] = {&gArg0,arg1,&arg2,&gArg3,arg4,};
    printf("CGO: godot_call_void_string_object_int_string_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_object_int_string_variant ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg3);

}
godot_vector2 *godot_call_vector2_vector2_vector2_float_int_float(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_vector2 *arg1,float arg2,int arg3,float arg4) {
    godot_vector2 *ret = new_godot_vector2();
    const void *c_args[] = {arg0,arg1,&arg2,&arg3,&arg4,};
    printf("CGO: godot_call_vector2_vector2_vector2_float_int_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector2_vector2_vector2_float_int_float ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0) {
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_void_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object ...godot_method_bind_ptrcall returned\n");

}
bool godot_call_bool_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1) {
    godot_bool ret;
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_bool_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_int_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
int godot_call_int_int_float_vector3_quat_vector3(godot_object *instance, godot_method_bind *mb,int arg0,float arg1,godot_vector3 *arg2,godot_quat *arg3,godot_vector3 *arg4) {
    godot_int ret;
    const void *c_args[] = {&arg0,&arg1,arg2,arg3,arg4,};
    printf("CGO: godot_call_int_int_float_vector3_quat_vector3 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_int_float_vector3_quat_vector3 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
bool godot_call_bool_transform2d_vector2(godot_object *instance, godot_method_bind *mb,godot_transform2d *arg0,godot_vector2 *arg1) {
    godot_bool ret;
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_bool_transform2d_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_transform2d_vector2 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_string_int_int_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,int arg2,int arg3) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,&arg2,&arg3,};
    printf("CGO: godot_call_void_string_int_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_int_int_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
godot_pool_byte_array *godot_call_poolbytearray(godot_object *instance, godot_method_bind *mb) {
    godot_pool_byte_array *ret = new_godot_pool_byte_array();
    const void *c_args[] = {};
    printf("CGO: godot_call_poolbytearray calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_poolbytearray ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_vector2_bool(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,bool arg1) {
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_void_vector2_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_vector2_bool ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_float_float_float(godot_object *instance, godot_method_bind *mb,float arg0,float arg1,float arg2) {
    const void *c_args[] = {&arg0,&arg1,&arg2,};
    printf("CGO: godot_call_void_float_float_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_float_float_float ...godot_method_bind_ptrcall returned\n");

}
godot_plane *godot_call_plane_int(godot_object *instance, godot_method_bind *mb,int arg0) {
    godot_plane *ret = new_godot_plane();
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_plane_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_plane_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_transform *godot_call_transform_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0) {
    godot_transform *ret = new_godot_transform();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_transform_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_transform_rid ...godot_method_bind_ptrcall returned\n");
    return ret;
}
bool godot_call_bool_string_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,int arg2) {
    godot_bool ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&gArg0,&gArg1,&arg2,};
    printf("CGO: godot_call_bool_string_string_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_string_string_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg1);
    return ret;
}
godot_array *godot_call_array_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1) {
    godot_array *ret = new_godot_array();
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_array_rid_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_array_rid_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
const char *godot_call_string(godot_object *instance, godot_method_bind *mb) {
    godot_string *ret = new_godot_string();
    const void *c_args[] = {};
    printf("CGO: godot_call_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_string ...godot_method_bind_ptrcall returned\n");
    godot_char_string gcsRet = godot_string_utf8(ret);
    const char *cRet = api->godot_char_string_get_data(&gcsRet);
    api->godot_char_string_destroy(&gcsRet);
    return cRet;

}
void godot_call_void_poolcolorarray(godot_object *instance, godot_method_bind *mb,godot_pool_color_array *arg0) {
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_void_poolcolorarray calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_poolcolorarray ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_poolvector2array_int(godot_object *instance, godot_method_bind *mb,godot_pool_vector2_array *arg0,int arg1) {
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_void_poolvector2array_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_poolvector2array_int ...godot_method_bind_ptrcall returned\n");

}
godot_pool_string_array *godot_call_poolstringarray_int(godot_object *instance, godot_method_bind *mb,int arg0) {
    godot_pool_string_array *ret = new_godot_pool_string_array();
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_poolstringarray_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_poolstringarray_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_rid_rect2_rect2_color_bool_object_bool(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rect2 *arg1,godot_rect2 *arg2,godot_color *arg3,bool arg4,godot_object *arg5,bool arg6) {
    const void *c_args[] = {arg0,arg1,arg2,arg3,&arg4,arg5,&arg6,};
    printf("CGO: godot_call_void_rid_rect2_rect2_color_bool_object_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_rect2_rect2_color_bool_object_bool ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rid_rid_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rid *arg1,godot_rid *arg2) {
    const void *c_args[] = {arg0,arg1,arg2,};
    printf("CGO: godot_call_void_rid_rid_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_rid_rid ...godot_method_bind_ptrcall returned\n");

}
godot_node_path *godot_call_nodepath_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0) {
    godot_node_path *ret = new_godot_node_path();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_nodepath_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_nodepath_object ...godot_method_bind_ptrcall returned\n");
    return ret;
}
int godot_call_int_int_string_int(godot_object *instance, godot_method_bind *mb,int arg0,char *arg1,int arg2) {
    godot_int ret;
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&arg0,&gArg1,&arg2,};
    printf("CGO: godot_call_int_int_string_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_int_string_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);
    return ret;
}
godot_vector3 *godot_call_vector3(godot_object *instance, godot_method_bind *mb) {
    godot_vector3 *ret = new_godot_vector3();
    const void *c_args[] = {};
    printf("CGO: godot_call_vector3 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector3 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_bool(godot_object *instance, godot_method_bind *mb,bool arg0) {
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_void_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_bool ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_int_int_bool(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,bool arg2) {
    const void *c_args[] = {&arg0,&arg1,&arg2,};
    printf("CGO: godot_call_void_int_int_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_int_bool ...godot_method_bind_ptrcall returned\n");

}
int godot_call_int_object_bool(godot_object *instance, godot_method_bind *mb,godot_object *arg0,bool arg1) {
    godot_int ret;
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_int_object_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_object_bool ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_int_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,int arg2,int arg3) {
    const void *c_args[] = {&arg0,&arg1,&arg2,&arg3,};
    printf("CGO: godot_call_void_int_int_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_int_int_int ...godot_method_bind_ptrcall returned\n");

}
godot_object *godot_call_object_float(godot_object *instance, godot_method_bind *mb,float arg0) {
    godot_object *ret;
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_object_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_object_float ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_variant *godot_call_variant_transform2d_object_transform2d(godot_object *instance, godot_method_bind *mb,godot_transform2d *arg0,godot_object *arg1,godot_transform2d *arg2) {
    godot_variant *ret = new_godot_variant();
    const void *c_args[] = {arg0,arg1,arg2,};
    printf("CGO: godot_call_variant_transform2d_object_transform2d calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_variant_transform2d_object_transform2d ...godot_method_bind_ptrcall returned\n");
    return ret;
}
bool godot_call_bool_vector2_variant(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_variant *arg1) {
    godot_bool ret;
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_bool_vector2_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_vector2_variant ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_object *godot_call_object_vector2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0) {
    godot_object *ret;
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_object_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_object_vector2 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_pool_vector2_array *godot_call_poolvector2array_vector2_vector2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_vector2 *arg1) {
    godot_pool_vector2_array *ret = new_godot_pool_vector2_array();
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_poolvector2array_vector2_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_poolvector2array_vector2_vector2 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_vector2 *godot_call_vector2_vector2_bool(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,bool arg1) {
    godot_vector2 *ret = new_godot_vector2();
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_vector2_vector2_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector2_vector2_bool ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_variant *godot_call_variant_object_string(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1) {
    godot_variant *ret = new_godot_variant();
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {arg0,&gArg1,};
    printf("CGO: godot_call_variant_object_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_variant_object_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);
    return ret;
}
bool godot_call_bool_string_int_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,char *arg2,int arg3) {
    godot_bool ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg2;
    api->godot_string_new(&gArg2);
    api->godot_string_parse_utf8(&gArg2, arg2);
    const void *c_args[] = {&gArg0,&arg1,&gArg2,&arg3,};
    printf("CGO: godot_call_bool_string_int_string_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_string_int_string_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg2);
    return ret;
}
bool godot_call_bool_float(godot_object *instance, godot_method_bind *mb,float arg0) {
    godot_bool ret;
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_bool_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_float ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_object *godot_call_object_nodepath(godot_object *instance, godot_method_bind *mb,godot_node_path *arg0) {
    godot_object *ret;
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_object_nodepath calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_object_nodepath ...godot_method_bind_ptrcall returned\n");
    return ret;
}
int godot_call_int_string_object_string_array_int(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1,char *arg2,godot_array *arg3,int arg4) {
    godot_int ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg2;
    api->godot_string_new(&gArg2);
    api->godot_string_parse_utf8(&gArg2, arg2);
    const void *c_args[] = {&gArg0,arg1,&gArg2,arg3,&arg4,};
    printf("CGO: godot_call_int_string_object_string_array_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_string_object_string_array_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg2);
    return ret;
}
int godot_call_int_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0) {
    godot_int ret;
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_int_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_rid ...godot_method_bind_ptrcall returned\n");
    return ret;
}
bool godot_call_bool_object_string_object_string_variant_float_int_int_float(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1,godot_object *arg2,char *arg3,godot_variant *arg4,float arg5,int arg6,int arg7,float arg8) {
    godot_bool ret;
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    godot_string gArg3;
    api->godot_string_new(&gArg3);
    api->godot_string_parse_utf8(&gArg3, arg3);
    const void *c_args[] = {arg0,&gArg1,arg2,&gArg3,arg4,&arg5,&arg6,&arg7,&arg8,};
    printf("CGO: godot_call_bool_object_string_object_string_variant_float_int_int_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_object_string_object_string_variant_float_int_int_float ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);
    godot_string_destroy(&gArg3);
    return ret;
}
godot_vector2 *godot_call_vector2_float(godot_object *instance, godot_method_bind *mb,float arg0) {
    godot_vector2 *ret = new_godot_vector2();
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_vector2_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector2_float ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_string_string_poolstringarray(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,godot_pool_string_array *arg2) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&gArg0,&gArg1,arg2,};
    printf("CGO: godot_call_void_string_string_poolstringarray calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_string_poolstringarray ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg1);

}
int godot_call_int_object_int(godot_object *instance, godot_method_bind *mb,godot_object *arg0,int arg1) {
    godot_int ret;
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_int_object_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_object_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_vector2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0) {
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_void_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_vector2 ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_string_float(godot_object *instance, godot_method_bind *mb,char *arg0,float arg1) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,};
    printf("CGO: godot_call_void_string_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_float ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
void godot_call_void_int_int_poolbytearray(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,godot_pool_byte_array *arg2) {
    const void *c_args[] = {&arg0,&arg1,arg2,};
    printf("CGO: godot_call_void_int_int_poolbytearray calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_int_poolbytearray ...godot_method_bind_ptrcall returned\n");

}
float godot_call_float_object_vector2_string_string_color(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_vector2 *arg1,char *arg2,char *arg3,godot_color *arg4) {
    godot_real ret;
    godot_string gArg2;
    api->godot_string_new(&gArg2);
    api->godot_string_parse_utf8(&gArg2, arg2);
    godot_string gArg3;
    api->godot_string_new(&gArg3);
    api->godot_string_parse_utf8(&gArg3, arg3);
    const void *c_args[] = {arg0,arg1,&gArg2,&gArg3,arg4,};
    printf("CGO: godot_call_float_object_vector2_string_string_color calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_float_object_vector2_string_string_color ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg2);
    godot_string_destroy(&gArg3);
    return ret;
}
void godot_call_void_rid_color(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_color *arg1) {
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_void_rid_color calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_color ...godot_method_bind_ptrcall returned\n");

}
godot_array *godot_call_array(godot_object *instance, godot_method_bind *mb) {
    godot_array *ret = new_godot_array();
    const void *c_args[] = {};
    printf("CGO: godot_call_array calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_array ...godot_method_bind_ptrcall returned\n");
    return ret;
}
int godot_call_int_string_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,int arg2) {
    godot_int ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&gArg0,&gArg1,&arg2,};
    printf("CGO: godot_call_int_string_string_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_string_string_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg1);
    return ret;
}
godot_rect2 *godot_call_rect2_int(godot_object *instance, godot_method_bind *mb,int arg0) {
    godot_rect2 *ret = new_godot_rect2();
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_rect2_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_rect2_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
bool godot_call_bool_transform_vector3(godot_object *instance, godot_method_bind *mb,godot_transform *arg0,godot_vector3 *arg1) {
    godot_bool ret;
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_bool_transform_vector3 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_transform_vector3 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_rid_object_string(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_object *arg1,char *arg2) {
    godot_string gArg2;
    api->godot_string_new(&gArg2);
    api->godot_string_parse_utf8(&gArg2, arg2);
    const void *c_args[] = {arg0,arg1,&gArg2,};
    printf("CGO: godot_call_void_rid_object_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_object_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg2);

}
godot_pool_color_array *godot_call_poolcolorarray(godot_object *instance, godot_method_bind *mb) {
    godot_pool_color_array *ret = new_godot_pool_color_array();
    const void *c_args[] = {};
    printf("CGO: godot_call_poolcolorarray calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_poolcolorarray ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_vector2_vector2_color_float_bool(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_vector2 *arg1,godot_color *arg2,float arg3,bool arg4) {
    const void *c_args[] = {arg0,arg1,arg2,&arg3,&arg4,};
    printf("CGO: godot_call_void_vector2_vector2_color_float_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_vector2_vector2_color_float_bool ...godot_method_bind_ptrcall returned\n");

}
godot_vector2 *godot_call_vector2_vector2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0) {
    godot_vector2 *ret = new_godot_vector2();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_vector2_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector2_vector2 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
int godot_call_int_vector2_float_float_int_int(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,float arg1,float arg2,int arg3,int arg4) {
    godot_int ret;
    const void *c_args[] = {arg0,&arg1,&arg2,&arg3,&arg4,};
    printf("CGO: godot_call_int_vector2_float_float_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_vector2_float_float_int_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_poolrealarray(godot_object *instance, godot_method_bind *mb,godot_pool_real_array *arg0) {
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_void_poolrealarray calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_poolrealarray ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_object_string_int_int(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1,int arg2,int arg3) {
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {arg0,&gArg1,&arg2,&arg3,};
    printf("CGO: godot_call_void_object_string_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_string_int_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);

}
void godot_call_void_string_float_float_bool(godot_object *instance, godot_method_bind *mb,char *arg0,float arg1,float arg2,bool arg3) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,&arg2,&arg3,};
    printf("CGO: godot_call_void_string_float_float_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_float_float_bool ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
godot_object *godot_call_object_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1) {
    godot_object *ret;
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_object_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_object_int_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_vector3 *godot_call_vector3_float_bool(godot_object *instance, godot_method_bind *mb,float arg0,bool arg1) {
    godot_vector3 *ret = new_godot_vector3();
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_vector3_float_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector3_float_bool ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_color *godot_call_color_float(godot_object *instance, godot_method_bind *mb,float arg0) {
    godot_color *ret = new_godot_color();
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_color_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_color_float ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_rid_object_string_variant(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_object *arg1,char *arg2,godot_variant *arg3) {
    godot_string gArg2;
    api->godot_string_new(&gArg2);
    api->godot_string_parse_utf8(&gArg2, arg2);
    const void *c_args[] = {arg0,arg1,&gArg2,arg3,};
    printf("CGO: godot_call_void_rid_object_string_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_object_string_variant ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg2);

}
bool godot_call_bool_string_int_int_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,int arg2,int arg3) {
    godot_bool ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,&arg2,&arg3,};
    printf("CGO: godot_call_bool_string_int_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_string_int_int_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
void godot_call_void_rid_aabb(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_aabb *arg1) {
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_void_rid_aabb calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_aabb ...godot_method_bind_ptrcall returned\n");

}
godot_object *godot_call_object_object_int(godot_object *instance, godot_method_bind *mb,godot_object *arg0,int arg1) {
    godot_object *ret;
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_object_object_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_object_object_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_string_string_color_bool(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,godot_color *arg2,bool arg3) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&gArg0,&gArg1,arg2,&arg3,};
    printf("CGO: godot_call_void_string_string_color_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_string_color_bool ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg1);

}
int godot_call_int_int(godot_object *instance, godot_method_bind *mb,int arg0) {
    godot_int ret;
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_rid_int_int_int(godot_object *instance, godot_method_bind *mb,int arg0,godot_rid *arg1,int arg2,int arg3,int arg4) {
    const void *c_args[] = {&arg0,arg1,&arg2,&arg3,&arg4,};
    printf("CGO: godot_call_void_int_rid_int_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_rid_int_int_int ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rid_vector2_vector2(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_vector2 *arg1,godot_vector2 *arg2) {
    const void *c_args[] = {arg0,arg1,arg2,};
    printf("CGO: godot_call_void_rid_vector2_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_vector2_vector2 ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rid_int_float(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,float arg2) {
    const void *c_args[] = {arg0,&arg1,&arg2,};
    printf("CGO: godot_call_void_rid_int_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_int_float ...godot_method_bind_ptrcall returned\n");

}
godot_rid *godot_call_rid_rid_transform_rid_transform(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_transform *arg1,godot_rid *arg2,godot_transform *arg3) {
    godot_rid *ret = new_godot_rid();
    const void *c_args[] = {arg0,arg1,arg2,arg3,};
    printf("CGO: godot_call_rid_rid_transform_rid_transform calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_rid_rid_transform_rid_transform ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_string_object_int(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1,int arg2) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,arg1,&arg2,};
    printf("CGO: godot_call_void_string_object_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_object_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
void godot_call_void_int_vector3(godot_object *instance, godot_method_bind *mb,int arg0,godot_vector3 *arg1) {
    const void *c_args[] = {&arg0,arg1,};
    printf("CGO: godot_call_void_int_vector3 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_vector3 ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rid_float(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,float arg1) {
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_void_rid_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_float ...godot_method_bind_ptrcall returned\n");

}
godot_rid *godot_call_rid_rid_vector3_rid_vector3(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_vector3 *arg1,godot_rid *arg2,godot_vector3 *arg3) {
    godot_rid *ret = new_godot_rid();
    const void *c_args[] = {arg0,arg1,arg2,arg3,};
    printf("CGO: godot_call_rid_rid_vector3_rid_vector3 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_rid_rid_vector3_rid_vector3 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_rid_rect2(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rect2 *arg1) {
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_void_rid_rect2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_rect2 ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_int_object_int_bool_string(godot_object *instance, godot_method_bind *mb,int arg0,godot_object *arg1,int arg2,bool arg3,char *arg4) {
    godot_string gArg4;
    api->godot_string_new(&gArg4);
    api->godot_string_parse_utf8(&gArg4, arg4);
    const void *c_args[] = {&arg0,arg1,&arg2,&arg3,&gArg4,};
    printf("CGO: godot_call_void_int_object_int_bool_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_object_int_bool_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg4);

}
godot_vector3 *godot_call_vector3_vector3(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0) {
    godot_vector3 *ret = new_godot_vector3();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_vector3_vector3 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector3_vector3 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
bool godot_call_bool_vector2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0) {
    godot_bool ret;
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_bool_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_vector2 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_poolvector2array_poolcolorarray_poolvector2array_object_object_bool(godot_object *instance, godot_method_bind *mb,godot_pool_vector2_array *arg0,godot_pool_color_array *arg1,godot_pool_vector2_array *arg2,godot_object *arg3,godot_object *arg4,bool arg5) {
    const void *c_args[] = {arg0,arg1,arg2,arg3,arg4,&arg5,};
    printf("CGO: godot_call_void_poolvector2array_poolcolorarray_poolvector2array_object_object_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_poolvector2array_poolcolorarray_poolvector2array_object_object_bool ...godot_method_bind_ptrcall returned\n");

}
godot_array *godot_call_array_string(godot_object *instance, godot_method_bind *mb,char *arg0) {
    godot_array *ret = new_godot_array();
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,};
    printf("CGO: godot_call_array_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_array_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
void godot_call_void_int_poolintarray(godot_object *instance, godot_method_bind *mb,int arg0,godot_pool_int_array *arg1) {
    const void *c_args[] = {&arg0,arg1,};
    printf("CGO: godot_call_void_int_poolintarray calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_poolintarray ...godot_method_bind_ptrcall returned\n");

}
godot_pool_vector2_array *godot_call_poolvector2array_int(godot_object *instance, godot_method_bind *mb,int arg0) {
    godot_pool_vector2_array *ret = new_godot_pool_vector2_array();
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_poolvector2array_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_poolvector2array_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_pool_int_array *godot_call_poolintarray_int(godot_object *instance, godot_method_bind *mb,int arg0) {
    godot_pool_int_array *ret = new_godot_pool_int_array();
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_poolintarray_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_poolintarray_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_vector2_float_vector2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,float arg1,godot_vector2 *arg2) {
    const void *c_args[] = {arg0,&arg1,arg2,};
    printf("CGO: godot_call_void_vector2_float_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_vector2_float_vector2 ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_int_transform2d(godot_object *instance, godot_method_bind *mb,int arg0,godot_transform2d *arg1) {
    const void *c_args[] = {&arg0,arg1,};
    printf("CGO: godot_call_void_int_transform2d calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_transform2d ...godot_method_bind_ptrcall returned\n");

}
godot_color *godot_call_color_int(godot_object *instance, godot_method_bind *mb,int arg0) {
    godot_color *ret = new_godot_color();
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_color_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_color_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_rid_int_transform2d(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,godot_transform2d *arg2) {
    const void *c_args[] = {arg0,&arg1,arg2,};
    printf("CGO: godot_call_void_rid_int_transform2d calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_int_transform2d ...godot_method_bind_ptrcall returned\n");

}
godot_variant *godot_call_variant_variant(godot_object *instance, godot_method_bind *mb,godot_variant *arg0) {
    godot_variant *ret = new_godot_variant();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_variant_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_variant_variant ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_variant(godot_object *instance, godot_method_bind *mb,int arg0,godot_variant *arg1) {
    const void *c_args[] = {&arg0,arg1,};
    printf("CGO: godot_call_void_int_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_variant ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rid_rect2_rid_rect2_color_bool_rid_bool(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rect2 *arg1,godot_rid *arg2,godot_rect2 *arg3,godot_color *arg4,bool arg5,godot_rid *arg6,bool arg7) {
    const void *c_args[] = {arg0,arg1,arg2,arg3,arg4,&arg5,arg6,&arg7,};
    printf("CGO: godot_call_void_rid_rect2_rid_rect2_color_bool_rid_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_rect2_rid_rect2_color_bool_rid_bool ...godot_method_bind_ptrcall returned\n");

}
godot_variant *godot_call_variant(godot_object *instance, godot_method_bind *mb) {
    godot_variant *ret = new_godot_variant();
    const void *c_args[] = {};
    printf("CGO: godot_call_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_variant ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_vector3_vector3(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0,godot_vector3 *arg1) {
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_void_vector3_vector3 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_vector3_vector3 ...godot_method_bind_ptrcall returned\n");

}
bool godot_call_bool(godot_object *instance, godot_method_bind *mb) {
    godot_bool ret;
    const void *c_args[] = {};
    printf("CGO: godot_call_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_poolvector2array_poolcolorarray_poolvector2array_object_float_object(godot_object *instance, godot_method_bind *mb,godot_pool_vector2_array *arg0,godot_pool_color_array *arg1,godot_pool_vector2_array *arg2,godot_object *arg3,float arg4,godot_object *arg5) {
    const void *c_args[] = {arg0,arg1,arg2,arg3,&arg4,arg5,};
    printf("CGO: godot_call_void_poolvector2array_poolcolorarray_poolvector2array_object_float_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_poolvector2array_poolcolorarray_poolvector2array_object_float_object ...godot_method_bind_ptrcall returned\n");

}
int godot_call_int_vector3(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0) {
    godot_int ret;
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_int_vector3 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_vector3 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_poolvector2array(godot_object *instance, godot_method_bind *mb,godot_pool_vector2_array *arg0) {
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_void_poolvector2array calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_poolvector2array ...godot_method_bind_ptrcall returned\n");

}
godot_array *godot_call_array_object_int(godot_object *instance, godot_method_bind *mb,godot_object *arg0,int arg1) {
    godot_array *ret = new_godot_array();
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_array_object_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_array_object_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
const char *godot_call_string_variant(godot_object *instance, godot_method_bind *mb,godot_variant *arg0) {
    godot_string *ret = new_godot_string();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_string_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_string_variant ...godot_method_bind_ptrcall returned\n");
    godot_char_string gcsRet = godot_string_utf8(ret);
    const char *cRet = api->godot_char_string_get_data(&gcsRet);
    api->godot_char_string_destroy(&gcsRet);
    return cRet;

}
bool godot_call_bool_vector2_rect2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_rect2 *arg1) {
    godot_bool ret;
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_bool_vector2_rect2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_vector2_rect2 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
int godot_call_int_int_float_bool(godot_object *instance, godot_method_bind *mb,int arg0,float arg1,bool arg2) {
    godot_int ret;
    const void *c_args[] = {&arg0,&arg1,&arg2,};
    printf("CGO: godot_call_int_int_float_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_int_float_bool ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_object_object_int_bool(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_object *arg1,int arg2,bool arg3) {
    const void *c_args[] = {arg0,arg1,&arg2,&arg3,};
    printf("CGO: godot_call_void_object_object_int_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_object_int_bool ...godot_method_bind_ptrcall returned\n");

}
godot_rid *godot_call_rid_int_int_float(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,float arg2) {
    godot_rid *ret = new_godot_rid();
    const void *c_args[] = {&arg0,&arg1,&arg2,};
    printf("CGO: godot_call_rid_int_int_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_rid_int_int_float ...godot_method_bind_ptrcall returned\n");
    return ret;
}
int godot_call_int(godot_object *instance, godot_method_bind *mb) {
    godot_int ret;
    const void *c_args[] = {};
    printf("CGO: godot_call_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_string_array(godot_object *instance, godot_method_bind *mb,char *arg0,godot_array *arg1) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,arg1,};
    printf("CGO: godot_call_void_string_array calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_array ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
void godot_call_void_rid_vector2(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_vector2 *arg1) {
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_void_rid_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_vector2 ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_object_string_bool(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1,bool arg2) {
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {arg0,&gArg1,&arg2,};
    printf("CGO: godot_call_void_object_string_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_string_bool ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);

}
godot_object *godot_call_object_string_bool_string(godot_object *instance, godot_method_bind *mb,char *arg0,bool arg1,char *arg2) {
    godot_object *ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg2;
    api->godot_string_new(&gArg2);
    api->godot_string_parse_utf8(&gArg2, arg2);
    const void *c_args[] = {&gArg0,&arg1,&gArg2,};
    printf("CGO: godot_call_object_string_bool_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_object_string_bool_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg2);
    return ret;
}
godot_variant *godot_call_variant_int_string(godot_object *instance, godot_method_bind *mb,int arg0,char *arg1) {
    godot_variant *ret = new_godot_variant();
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&arg0,&gArg1,};
    printf("CGO: godot_call_variant_int_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_variant_int_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);
    return ret;
}
godot_rid *godot_call_rid_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1) {
    godot_rid *ret = new_godot_rid();
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_rid_rid_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_rid_rid_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_rid_vector3(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_vector3 *arg1) {
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_void_rid_vector3 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_vector3 ...godot_method_bind_ptrcall returned\n");

}
godot_variant *godot_call_variant_string_string(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1) {
    godot_variant *ret = new_godot_variant();
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&gArg0,&gArg1,};
    printf("CGO: godot_call_variant_string_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_variant_string_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg1);
    return ret;
}
godot_variant *godot_call_variant_int_string_string(godot_object *instance, godot_method_bind *mb,int arg0,char *arg1,char *arg2) {
    godot_variant *ret = new_godot_variant();
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    godot_string gArg2;
    api->godot_string_new(&gArg2);
    api->godot_string_parse_utf8(&gArg2, arg2);
    const void *c_args[] = {&arg0,&gArg1,&gArg2,};
    printf("CGO: godot_call_variant_int_string_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_variant_int_string_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);
    godot_string_destroy(&gArg2);
    return ret;
}
void godot_call_void_transform(godot_object *instance, godot_method_bind *mb,godot_transform *arg0) {
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_void_transform calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_transform ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rect2_color_bool(godot_object *instance, godot_method_bind *mb,godot_rect2 *arg0,godot_color *arg1,bool arg2) {
    const void *c_args[] = {arg0,arg1,&arg2,};
    printf("CGO: godot_call_void_rect2_color_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rect2_color_bool ...godot_method_bind_ptrcall returned\n");

}
godot_rid *godot_call_rid_int(godot_object *instance, godot_method_bind *mb,int arg0) {
    godot_rid *ret = new_godot_rid();
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_rid_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_rid_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
bool godot_call_bool_rid_int_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,int arg2) {
    godot_bool ret;
    const void *c_args[] = {arg0,&arg1,&arg2,};
    printf("CGO: godot_call_bool_rid_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_rid_int_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_rid_rid_transform(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rid *arg1,godot_transform *arg2) {
    const void *c_args[] = {arg0,arg1,arg2,};
    printf("CGO: godot_call_void_rid_rid_transform calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_rid_transform ...godot_method_bind_ptrcall returned\n");

}
godot_object *godot_call_object_float_bool(godot_object *instance, godot_method_bind *mb,float arg0,bool arg1) {
    godot_object *ret;
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_object_float_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_object_float_bool ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_bool(godot_object *instance, godot_method_bind *mb,int arg0,bool arg1) {
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_void_int_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_bool ...godot_method_bind_ptrcall returned\n");

}
bool godot_call_bool_string_dictionary(godot_object *instance, godot_method_bind *mb,char *arg0,godot_dictionary *arg1) {
    godot_bool ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,arg1,};
    printf("CGO: godot_call_bool_string_dictionary calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_string_dictionary ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
void godot_call_void_rid_int_int_bool(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,int arg2,bool arg3) {
    const void *c_args[] = {arg0,&arg1,&arg2,&arg3,};
    printf("CGO: godot_call_void_rid_int_int_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_int_int_bool ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_nodepath(godot_object *instance, godot_method_bind *mb,godot_node_path *arg0) {
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_void_nodepath calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_nodepath ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_variant_object(godot_object *instance, godot_method_bind *mb,godot_variant *arg0,godot_object *arg1) {
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_void_variant_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_variant_object ...godot_method_bind_ptrcall returned\n");

}
int godot_call_int_object_transform2d_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_transform2d *arg1,godot_object *arg2) {
    godot_int ret;
    const void *c_args[] = {arg0,arg1,arg2,};
    printf("CGO: godot_call_int_object_transform2d_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_object_transform2d_object ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_array *godot_call_array_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0) {
    godot_array *ret = new_godot_array();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_array_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_array_object ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_object *godot_call_object_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0) {
    godot_object *ret;
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_object_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_object_rid ...godot_method_bind_ptrcall returned\n");
    return ret;
}
const char *godot_call_string_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1) {
    godot_string *ret = new_godot_string();
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,};
    printf("CGO: godot_call_string_string_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_string_string_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_char_string gcsRet = godot_string_utf8(ret);
    const char *cRet = api->godot_char_string_get_data(&gcsRet);
    api->godot_char_string_destroy(&gcsRet);
    return cRet;

}
void godot_call_void_poolstringarray_int(godot_object *instance, godot_method_bind *mb,godot_pool_string_array *arg0,int arg1) {
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_void_poolstringarray_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_poolstringarray_int ...godot_method_bind_ptrcall returned\n");

}
float godot_call_float_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0) {
    godot_real ret;
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_float_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_float_rid ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_string_int(godot_object *instance, godot_method_bind *mb,int arg0,char *arg1,int arg2) {
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&arg0,&gArg1,&arg2,};
    printf("CGO: godot_call_void_int_string_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_string_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);

}
void godot_call_void_bool_bool(godot_object *instance, godot_method_bind *mb,bool arg0,bool arg1) {
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_void_bool_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_bool_bool ...godot_method_bind_ptrcall returned\n");

}
godot_object *godot_call_object_string(godot_object *instance, godot_method_bind *mb,char *arg0) {
    godot_object *ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,};
    printf("CGO: godot_call_object_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_object_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
godot_vector2 *godot_call_vector2_float_bool(godot_object *instance, godot_method_bind *mb,float arg0,bool arg1) {
    godot_vector2 *ret = new_godot_vector2();
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_vector2_float_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector2_float_bool ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_array *godot_call_array_vector2_int_array_int(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,int arg1,godot_array *arg2,int arg3) {
    godot_array *ret = new_godot_array();
    const void *c_args[] = {arg0,&arg1,arg2,&arg3,};
    printf("CGO: godot_call_array_vector2_int_array_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_array_vector2_int_array_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_rid_variant(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_variant *arg1) {
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_void_rid_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_variant ...godot_method_bind_ptrcall returned\n");

}
godot_basis *godot_call_basis(godot_object *instance, godot_method_bind *mb) {
    godot_basis *ret = new_godot_basis();
    const void *c_args[] = {};
    printf("CGO: godot_call_basis calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_basis ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_rid *godot_call_rid_vector2_vector2_rid_rid(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_vector2 *arg1,godot_rid *arg2,godot_rid *arg3) {
    godot_rid *ret = new_godot_rid();
    const void *c_args[] = {arg0,arg1,arg2,arg3,};
    printf("CGO: godot_call_rid_vector2_vector2_rid_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_rid_vector2_vector2_rid_rid ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_rid_poolvector2array_poolcolorarray_float_bool(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_pool_vector2_array *arg1,godot_pool_color_array *arg2,float arg3,bool arg4) {
    const void *c_args[] = {arg0,arg1,arg2,&arg3,&arg4,};
    printf("CGO: godot_call_void_rid_poolvector2array_poolcolorarray_float_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_poolvector2array_poolcolorarray_float_bool ...godot_method_bind_ptrcall returned\n");

}
int godot_call_int_vector2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0) {
    godot_int ret;
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_int_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_vector2 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_float_float_float(godot_object *instance, godot_method_bind *mb,int arg0,float arg1,float arg2,float arg3) {
    const void *c_args[] = {&arg0,&arg1,&arg2,&arg3,};
    printf("CGO: godot_call_void_int_float_float_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_float_float_float ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rid_poolvector2array(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_pool_vector2_array *arg1) {
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_void_rid_poolvector2array calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_poolvector2array ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_object_color_bool(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_color *arg1,bool arg2) {
    const void *c_args[] = {arg0,arg1,&arg2,};
    printf("CGO: godot_call_void_object_color_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_color_bool ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_aabb(godot_object *instance, godot_method_bind *mb,godot_aabb *arg0) {
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_void_aabb calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_aabb ...godot_method_bind_ptrcall returned\n");

}
bool godot_call_bool_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1) {
    godot_bool ret;
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_bool_rid_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_rid_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_rid_rid_vector2(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rid *arg1,godot_vector2 *arg2) {
    const void *c_args[] = {arg0,arg1,arg2,};
    printf("CGO: godot_call_void_rid_rid_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_rid_vector2 ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_poolvector2array_poolcolorarray_float_bool(godot_object *instance, godot_method_bind *mb,godot_pool_vector2_array *arg0,godot_pool_color_array *arg1,float arg2,bool arg3) {
    const void *c_args[] = {arg0,arg1,&arg2,&arg3,};
    printf("CGO: godot_call_void_poolvector2array_poolcolorarray_float_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_poolvector2array_poolcolorarray_float_bool ...godot_method_bind_ptrcall returned\n");

}
godot_pool_vector3_array *godot_call_poolvector3array_int_float(godot_object *instance, godot_method_bind *mb,int arg0,float arg1) {
    godot_pool_vector3_array *ret = new_godot_pool_vector3_array();
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_poolvector3array_int_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_poolvector3array_int_float ...godot_method_bind_ptrcall returned\n");
    return ret;
}
const char *godot_call_string_string_object(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1) {
    godot_string *ret = new_godot_string();
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,arg1,};
    printf("CGO: godot_call_string_string_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_string_string_object ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_char_string gcsRet = godot_string_utf8(ret);
    const char *cRet = api->godot_char_string_get_data(&gcsRet);
    api->godot_char_string_destroy(&gcsRet);
    return cRet;

}
bool godot_call_bool_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0) {
    godot_bool ret;
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_bool_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_rid ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_rid_poolvector2array_bool(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_pool_vector2_array *arg1,bool arg2) {
    const void *c_args[] = {arg0,arg1,&arg2,};
    printf("CGO: godot_call_void_rid_poolvector2array_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_poolvector2array_bool ...godot_method_bind_ptrcall returned\n");

}
int godot_call_int_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1) {
    godot_int ret;
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_int_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_int_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_pool_string_array *godot_call_poolstringarray_string(godot_object *instance, godot_method_bind *mb,char *arg0) {
    godot_pool_string_array *ret = new_godot_pool_string_array();
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,};
    printf("CGO: godot_call_poolstringarray_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_poolstringarray_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
godot_dictionary *godot_call_dictionary_bool(godot_object *instance, godot_method_bind *mb,bool arg0) {
    godot_dictionary *ret = new_godot_dictionary();
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_dictionary_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_dictionary_bool ...godot_method_bind_ptrcall returned\n");
    return ret;
}
int godot_call_int_vector2_bool(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,bool arg1) {
    godot_int ret;
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_int_vector2_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_vector2_bool ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_string_array_bool(godot_object *instance, godot_method_bind *mb,char *arg0,godot_array *arg1,bool arg2) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,arg1,&arg2,};
    printf("CGO: godot_call_void_string_array_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_array_bool ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
void godot_call_void_rid_int_bool(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,bool arg2) {
    const void *c_args[] = {arg0,&arg1,&arg2,};
    printf("CGO: godot_call_void_rid_int_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_int_bool ...godot_method_bind_ptrcall returned\n");

}
bool godot_call_bool_string(godot_object *instance, godot_method_bind *mb,char *arg0) {
    godot_bool ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,};
    printf("CGO: godot_call_bool_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
godot_pool_real_array *godot_call_poolrealarray_int(godot_object *instance, godot_method_bind *mb,int arg0) {
    godot_pool_real_array *ret = new_godot_pool_real_array();
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_poolrealarray_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_poolrealarray_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_poolvector2array(godot_object *instance, godot_method_bind *mb,int arg0,godot_pool_vector2_array *arg1) {
    const void *c_args[] = {&arg0,arg1,};
    printf("CGO: godot_call_void_int_poolvector2array calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_poolvector2array ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_int_int_vector2_float(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,godot_vector2 *arg2,float arg3) {
    const void *c_args[] = {&arg0,&arg1,arg2,&arg3,};
    printf("CGO: godot_call_void_int_int_vector2_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_int_vector2_float ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_vector3_vector3_vector3_int(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0,godot_vector3 *arg1,godot_vector3 *arg2,int arg3) {
    const void *c_args[] = {arg0,arg1,arg2,&arg3,};
    printf("CGO: godot_call_void_vector3_vector3_vector3_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_vector3_vector3_vector3_int ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_object_bool(godot_object *instance, godot_method_bind *mb,godot_object *arg0,bool arg1) {
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_void_object_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_bool ...godot_method_bind_ptrcall returned\n");

}
godot_aabb *godot_call_aabb_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0) {
    godot_aabb *ret = new_godot_aabb();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_aabb_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_aabb_rid ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_rid_int_int_int_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,int arg2,int arg3,int arg4) {
    const void *c_args[] = {arg0,&arg1,&arg2,&arg3,&arg4,};
    printf("CGO: godot_call_void_rid_int_int_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_int_int_int_int ...godot_method_bind_ptrcall returned\n");

}
int godot_call_int_variant(godot_object *instance, godot_method_bind *mb,godot_variant *arg0) {
    godot_int ret;
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_int_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_variant ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_variant *godot_call_variant_array_array_int_array(godot_object *instance, godot_method_bind *mb,godot_array *arg0,godot_array *arg1,int arg2,godot_array *arg3) {
    godot_variant *ret = new_godot_variant();
    const void *c_args[] = {arg0,arg1,&arg2,arg3,};
    printf("CGO: godot_call_variant_array_array_int_array calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_variant_array_array_int_array ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_pool_vector3_array *godot_call_poolvector3array_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1) {
    godot_pool_vector3_array *ret = new_godot_pool_vector3_array();
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_poolvector3array_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_poolvector3array_int_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_dictionary *godot_call_dictionary(godot_object *instance, godot_method_bind *mb) {
    godot_dictionary *ret = new_godot_dictionary();
    const void *c_args[] = {};
    printf("CGO: godot_call_dictionary calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_dictionary ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_transform2d *godot_call_transform2d(godot_object *instance, godot_method_bind *mb) {
    godot_transform2d *ret = new_godot_transform2d();
    const void *c_args[] = {};
    printf("CGO: godot_call_transform2d calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_transform2d ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_variant *godot_call_variant_string_string_variant(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,godot_variant *arg2) {
    godot_variant *ret = new_godot_variant();
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&gArg0,&gArg1,arg2,};
    printf("CGO: godot_call_variant_string_string_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_variant_string_string_variant ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg1);
    return ret;
}
int godot_call_int_string_int_int_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,int arg2,int arg3) {
    godot_int ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,&arg2,&arg3,};
    printf("CGO: godot_call_int_string_int_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_string_int_int_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
godot_object *godot_call_object_bool(godot_object *instance, godot_method_bind *mb,bool arg0) {
    godot_object *ret;
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_object_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_object_bool ...godot_method_bind_ptrcall returned\n");
    return ret;
}
const char *godot_call_string_string(godot_object *instance, godot_method_bind *mb,char *arg0) {
    godot_string *ret = new_godot_string();
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,};
    printf("CGO: godot_call_string_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_string_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_char_string gcsRet = godot_string_utf8(ret);
    const char *cRet = api->godot_char_string_get_data(&gcsRet);
    api->godot_char_string_destroy(&gcsRet);
    return cRet;

}
bool godot_call_bool_transform2d_vector2_object_transform2d_vector2(godot_object *instance, godot_method_bind *mb,godot_transform2d *arg0,godot_vector2 *arg1,godot_object *arg2,godot_transform2d *arg3,godot_vector2 *arg4) {
    godot_bool ret;
    const void *c_args[] = {arg0,arg1,arg2,arg3,arg4,};
    printf("CGO: godot_call_bool_transform2d_vector2_object_transform2d_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_transform2d_vector2_object_transform2d_vector2 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_dictionary(godot_object *instance, godot_method_bind *mb,godot_dictionary *arg0) {
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_void_dictionary calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_dictionary ...godot_method_bind_ptrcall returned\n");

}
godot_dictionary *godot_call_dictionary_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0) {
    godot_dictionary *ret = new_godot_dictionary();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_dictionary_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_dictionary_object ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_bool_bool(godot_object *instance, godot_method_bind *mb,int arg0,bool arg1,bool arg2) {
    const void *c_args[] = {&arg0,&arg1,&arg2,};
    printf("CGO: godot_call_void_int_bool_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_bool_bool ...godot_method_bind_ptrcall returned\n");

}
const char *godot_call_string_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0) {
    godot_string *ret = new_godot_string();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_string_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_string_object ...godot_method_bind_ptrcall returned\n");
    godot_char_string gcsRet = godot_string_utf8(ret);
    const char *cRet = api->godot_char_string_get_data(&gcsRet);
    api->godot_char_string_destroy(&gcsRet);
    return cRet;

}
int godot_call_int_string_string(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1) {
    godot_int ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&gArg0,&gArg1,};
    printf("CGO: godot_call_int_string_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_string_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg1);
    return ret;
}
bool godot_call_bool_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1) {
    godot_bool ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,};
    printf("CGO: godot_call_bool_string_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_string_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
void godot_call_void_object_rect2_bool_color_bool_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_rect2 *arg1,bool arg2,godot_color *arg3,bool arg4,godot_object *arg5) {
    const void *c_args[] = {arg0,arg1,&arg2,arg3,&arg4,arg5,};
    printf("CGO: godot_call_void_object_rect2_bool_color_bool_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_rect2_bool_color_bool_object ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_int_int_bool_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,bool arg2,int arg3) {
    const void *c_args[] = {&arg0,&arg1,&arg2,&arg3,};
    printf("CGO: godot_call_void_int_int_bool_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_int_bool_int ...godot_method_bind_ptrcall returned\n");

}
godot_transform2d *godot_call_transform2d_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1) {
    godot_transform2d *ret = new_godot_transform2d();
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_transform2d_rid_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_transform2d_rid_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_color(godot_object *instance, godot_method_bind *mb,godot_color *arg0) {
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_void_color calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_color ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_poolvector3array_bool_bool(godot_object *instance, godot_method_bind *mb,godot_pool_vector3_array *arg0,bool arg1,bool arg2) {
    const void *c_args[] = {arg0,&arg1,&arg2,};
    printf("CGO: godot_call_void_poolvector3array_bool_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_poolvector3array_bool_bool ...godot_method_bind_ptrcall returned\n");

}
godot_variant *godot_call_variant_string_bool(godot_object *instance, godot_method_bind *mb,char *arg0,bool arg1) {
    godot_variant *ret = new_godot_variant();
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,};
    printf("CGO: godot_call_variant_string_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_variant_string_bool ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
const char *godot_call_string_string_string_bool_int_int(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,bool arg2,int arg3,int arg4) {
    godot_string *ret = new_godot_string();
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&gArg0,&gArg1,&arg2,&arg3,&arg4,};
    printf("CGO: godot_call_string_string_string_bool_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_string_string_string_bool_int_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg1);
    godot_char_string gcsRet = godot_string_utf8(ret);
    const char *cRet = api->godot_char_string_get_data(&gcsRet);
    api->godot_char_string_destroy(&gcsRet);
    return cRet;

}
void godot_call_void_rid_rid_rid_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rid *arg1,godot_rid *arg2,godot_rid *arg3) {
    const void *c_args[] = {arg0,arg1,arg2,arg3,};
    printf("CGO: godot_call_void_rid_rid_rid_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_rid_rid_rid ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_transform2d(godot_object *instance, godot_method_bind *mb,godot_transform2d *arg0) {
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_void_transform2d calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_transform2d ...godot_method_bind_ptrcall returned\n");

}
float godot_call_float_float(godot_object *instance, godot_method_bind *mb,float arg0) {
    godot_real ret;
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_float_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_float_float ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_color *godot_call_color_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1) {
    godot_color *ret = new_godot_color();
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_color_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_color_int_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_rect2(godot_object *instance, godot_method_bind *mb,int arg0,godot_rect2 *arg1) {
    const void *c_args[] = {&arg0,arg1,};
    printf("CGO: godot_call_void_int_rect2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_rect2 ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_int_plane(godot_object *instance, godot_method_bind *mb,int arg0,godot_plane *arg1) {
    const void *c_args[] = {&arg0,arg1,};
    printf("CGO: godot_call_void_int_plane calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_plane ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_string_int_vector2(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,godot_vector2 *arg2) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,arg2,};
    printf("CGO: godot_call_void_string_int_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_int_vector2 ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
void godot_call_void_rid_int_transform(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,godot_transform *arg2) {
    const void *c_args[] = {arg0,&arg1,arg2,};
    printf("CGO: godot_call_void_rid_int_transform calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_int_transform ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_string_int_string(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,char *arg2) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg2;
    api->godot_string_new(&gArg2);
    api->godot_string_parse_utf8(&gArg2, arg2);
    const void *c_args[] = {&gArg0,&arg1,&gArg2,};
    printf("CGO: godot_call_void_string_int_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_int_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg2);

}
void godot_call_void_rid_int_array_array_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,godot_array *arg2,godot_array *arg3,int arg4) {
    const void *c_args[] = {arg0,&arg1,arg2,arg3,&arg4,};
    printf("CGO: godot_call_void_rid_int_array_array_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_int_array_array_int ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rid_object_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_object *arg1,int arg2) {
    const void *c_args[] = {arg0,arg1,&arg2,};
    printf("CGO: godot_call_void_rid_object_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_object_int ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_int_nodepath(godot_object *instance, godot_method_bind *mb,int arg0,godot_node_path *arg1) {
    const void *c_args[] = {&arg0,arg1,};
    printf("CGO: godot_call_void_int_nodepath calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_nodepath ...godot_method_bind_ptrcall returned\n");

}
godot_rect2 *godot_call_rect2(godot_object *instance, godot_method_bind *mb) {
    godot_rect2 *ret = new_godot_rect2();
    const void *c_args[] = {};
    printf("CGO: godot_call_rect2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_rect2 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_pool_vector2_array *godot_call_poolvector2array_int_float(godot_object *instance, godot_method_bind *mb,int arg0,float arg1) {
    godot_pool_vector2_array *ret = new_godot_pool_vector2_array();
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_poolvector2array_int_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_poolvector2array_int_float ...godot_method_bind_ptrcall returned\n");
    return ret;
}
bool godot_call_bool_object_string_variant_object_string_float_int_int_float(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1,godot_variant *arg2,godot_object *arg3,char *arg4,float arg5,int arg6,int arg7,float arg8) {
    godot_bool ret;
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    godot_string gArg4;
    api->godot_string_new(&gArg4);
    api->godot_string_parse_utf8(&gArg4, arg4);
    const void *c_args[] = {arg0,&gArg1,arg2,arg3,&gArg4,&arg5,&arg6,&arg7,&arg8,};
    printf("CGO: godot_call_bool_object_string_variant_object_string_float_int_int_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_object_string_variant_object_string_float_int_int_float ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);
    godot_string_destroy(&gArg4);
    return ret;
}
bool godot_call_bool_object_float_string_variant_variant_variant_variant_variant(godot_object *instance, godot_method_bind *mb,godot_object *arg0,float arg1,char *arg2,godot_variant *arg3,godot_variant *arg4,godot_variant *arg5,godot_variant *arg6,godot_variant *arg7) {
    godot_bool ret;
    godot_string gArg2;
    api->godot_string_new(&gArg2);
    api->godot_string_parse_utf8(&gArg2, arg2);
    const void *c_args[] = {arg0,&arg1,&gArg2,arg3,arg4,arg5,arg6,arg7,};
    printf("CGO: godot_call_bool_object_float_string_variant_variant_variant_variant_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_object_float_string_variant_variant_variant_variant_variant ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg2);
    return ret;
}
void godot_call_void_nodepath_object_int(godot_object *instance, godot_method_bind *mb,godot_node_path *arg0,godot_object *arg1,int arg2) {
    const void *c_args[] = {arg0,arg1,&arg2,};
    printf("CGO: godot_call_void_nodepath_object_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_nodepath_object_int ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_object_string(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1) {
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {arg0,&gArg1,};
    printf("CGO: godot_call_void_object_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);

}
void godot_call_void_vector3_float(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0,float arg1) {
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_void_vector3_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_vector3_float ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rid_vector2_vector2_color_float_bool(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_vector2 *arg1,godot_vector2 *arg2,godot_color *arg3,float arg4,bool arg5) {
    const void *c_args[] = {arg0,arg1,arg2,arg3,&arg4,&arg5,};
    printf("CGO: godot_call_void_rid_vector2_vector2_color_float_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_vector2_vector2_color_float_bool ...godot_method_bind_ptrcall returned\n");

}
godot_plane *godot_call_plane(godot_object *instance, godot_method_bind *mb) {
    godot_plane *ret = new_godot_plane();
    const void *c_args[] = {};
    printf("CGO: godot_call_plane calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_plane ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_rid_int_int_float(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,int arg2,float arg3) {
    const void *c_args[] = {arg0,&arg1,&arg2,&arg3,};
    printf("CGO: godot_call_void_rid_int_int_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_int_int_float ...godot_method_bind_ptrcall returned\n");

}
godot_array *godot_call_array_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0) {
    godot_array *ret = new_godot_array();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_array_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_array_rid ...godot_method_bind_ptrcall returned\n");
    return ret;
}
bool godot_call_bool_int(godot_object *instance, godot_method_bind *mb,int arg0) {
    godot_bool ret;
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_bool_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_pool_int_array *godot_call_poolintarray_int_float_float(godot_object *instance, godot_method_bind *mb,int arg0,float arg1,float arg2) {
    godot_pool_int_array *ret = new_godot_pool_int_array();
    const void *c_args[] = {&arg0,&arg1,&arg2,};
    printf("CGO: godot_call_poolintarray_int_float_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_poolintarray_int_float_float ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_string_int_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,int arg2) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,&arg2,};
    printf("CGO: godot_call_void_string_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_int_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
godot_object *godot_call_object_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1) {
    godot_object *ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,};
    printf("CGO: godot_call_object_string_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_object_string_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
bool godot_call_bool_object_nodepath_variant_variant_float_int_int_float(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_node_path *arg1,godot_variant *arg2,godot_variant *arg3,float arg4,int arg5,int arg6,float arg7) {
    godot_bool ret;
    const void *c_args[] = {arg0,arg1,arg2,arg3,&arg4,&arg5,&arg6,&arg7,};
    printf("CGO: godot_call_bool_object_nodepath_variant_variant_float_int_int_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_object_nodepath_variant_variant_float_int_int_float ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_rid_poolvector2array_poolcolorarray_poolvector2array_rid_rid_bool(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_pool_vector2_array *arg1,godot_pool_color_array *arg2,godot_pool_vector2_array *arg3,godot_rid *arg4,godot_rid *arg5,bool arg6) {
    const void *c_args[] = {arg0,arg1,arg2,arg3,arg4,arg5,&arg6,};
    printf("CGO: godot_call_void_rid_poolvector2array_poolcolorarray_poolvector2array_rid_rid_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_poolvector2array_poolcolorarray_poolvector2array_rid_rid_bool ...godot_method_bind_ptrcall returned\n");

}
int godot_call_int_string_int_bool_bool(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,bool arg2,bool arg3) {
    godot_int ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,&arg2,&arg3,};
    printf("CGO: godot_call_int_string_int_bool_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_string_int_bool_bool ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
void godot_call_void_string_dictionary(godot_object *instance, godot_method_bind *mb,char *arg0,godot_dictionary *arg1) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,arg1,};
    printf("CGO: godot_call_void_string_dictionary calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_dictionary ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
void godot_call_void_rid_rect2_rid_bool_color_bool_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rect2 *arg1,godot_rid *arg2,bool arg3,godot_color *arg4,bool arg5,godot_rid *arg6) {
    const void *c_args[] = {arg0,arg1,arg2,&arg3,arg4,&arg5,arg6,};
    printf("CGO: godot_call_void_rid_rect2_rid_bool_color_bool_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_rect2_rid_bool_color_bool_rid ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_poolvector3array(godot_object *instance, godot_method_bind *mb,godot_pool_vector3_array *arg0) {
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_void_poolvector3array calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_poolvector3array ...godot_method_bind_ptrcall returned\n");

}
int godot_call_int_vector2_object(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_object *arg1) {
    godot_int ret;
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_int_vector2_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_vector2_object ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_vector3 *godot_call_vector3_vector3_vector3_float_int_float(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0,godot_vector3 *arg1,float arg2,int arg3,float arg4) {
    godot_vector3 *ret = new_godot_vector3();
    const void *c_args[] = {arg0,arg1,&arg2,&arg3,&arg4,};
    printf("CGO: godot_call_vector3_vector3_vector3_float_int_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector3_vector3_vector3_float_int_float ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_rid_vector2_float_color(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_vector2 *arg1,float arg2,godot_color *arg3) {
    const void *c_args[] = {arg0,arg1,&arg2,arg3,};
    printf("CGO: godot_call_void_rid_vector2_float_color calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_vector2_float_color ...godot_method_bind_ptrcall returned\n");

}
int godot_call_int_string_object(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1) {
    godot_int ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,arg1,};
    printf("CGO: godot_call_int_string_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_string_object ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
void godot_call_void_string_vector2(godot_object *instance, godot_method_bind *mb,char *arg0,godot_vector2 *arg1) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,arg1,};
    printf("CGO: godot_call_void_string_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_vector2 ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
void godot_call_void_object_object_vector3_vector3_int(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_object *arg1,godot_vector3 *arg2,godot_vector3 *arg3,int arg4) {
    const void *c_args[] = {arg0,arg1,arg2,arg3,&arg4,};
    printf("CGO: godot_call_void_object_object_vector3_vector3_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_object_vector3_vector3_int ...godot_method_bind_ptrcall returned\n");

}
int godot_call_int_poolbytearray(godot_object *instance, godot_method_bind *mb,godot_pool_byte_array *arg0) {
    godot_int ret;
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_int_poolbytearray calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_poolbytearray ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_int_color(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,godot_color *arg2) {
    const void *c_args[] = {&arg0,&arg1,arg2,};
    printf("CGO: godot_call_void_int_int_color calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_int_color ...godot_method_bind_ptrcall returned\n");

}
bool godot_call_bool_object_string(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1) {
    godot_bool ret;
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {arg0,&gArg1,};
    printf("CGO: godot_call_bool_object_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_object_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);
    return ret;
}
godot_array *godot_call_array_int_float(godot_object *instance, godot_method_bind *mb,int arg0,float arg1) {
    godot_array *ret = new_godot_array();
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_array_int_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_array_int_float ...godot_method_bind_ptrcall returned\n");
    return ret;
}
bool godot_call_bool_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0) {
    godot_bool ret;
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_bool_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_object ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_poolrealarray(godot_object *instance, godot_method_bind *mb,int arg0,godot_pool_real_array *arg1) {
    const void *c_args[] = {&arg0,arg1,};
    printf("CGO: godot_call_void_int_poolrealarray calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_poolrealarray ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rid_rid_transform2d(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rid *arg1,godot_transform2d *arg2) {
    const void *c_args[] = {arg0,arg1,arg2,};
    printf("CGO: godot_call_void_rid_rid_transform2d calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_rid_transform2d ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rid_bool(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,bool arg1) {
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_void_rid_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_bool ...godot_method_bind_ptrcall returned\n");

}
float godot_call_float_rid_int_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,int arg2) {
    godot_real ret;
    const void *c_args[] = {arg0,&arg1,&arg2,};
    printf("CGO: godot_call_float_rid_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_float_rid_int_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_object_string_poolstringarray(godot_object *instance, godot_method_bind *mb,godot_object *arg0,char *arg1,godot_pool_string_array *arg2) {
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {arg0,&gArg1,arg2,};
    printf("CGO: godot_call_void_object_string_poolstringarray calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_string_poolstringarray ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);

}
void godot_call_void_string_poolbytearray_bool(godot_object *instance, godot_method_bind *mb,char *arg0,godot_pool_byte_array *arg1,bool arg2) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,arg1,&arg2,};
    printf("CGO: godot_call_void_string_poolbytearray_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_poolbytearray_bool ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
void godot_call_void_string_poolstringarray(godot_object *instance, godot_method_bind *mb,char *arg0,godot_pool_string_array *arg1) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,arg1,};
    printf("CGO: godot_call_void_string_poolstringarray calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_poolstringarray ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
void godot_call_void_int_variant_bool(godot_object *instance, godot_method_bind *mb,int arg0,godot_variant *arg1,bool arg2) {
    const void *c_args[] = {&arg0,arg1,&arg2,};
    printf("CGO: godot_call_void_int_variant_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_variant_bool ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_bool_float(godot_object *instance, godot_method_bind *mb,bool arg0,float arg1) {
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_void_bool_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_bool_float ...godot_method_bind_ptrcall returned\n");

}
int godot_call_int_int_int_float(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,float arg2) {
    godot_int ret;
    const void *c_args[] = {&arg0,&arg1,&arg2,};
    printf("CGO: godot_call_int_int_int_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_int_int_float ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_object_object_bool(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_object *arg1,bool arg2) {
    const void *c_args[] = {arg0,arg1,&arg2,};
    printf("CGO: godot_call_void_object_object_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_object_bool ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_int_color_bool(godot_object *instance, godot_method_bind *mb,int arg0,godot_color *arg1,bool arg2) {
    const void *c_args[] = {&arg0,arg1,&arg2,};
    printf("CGO: godot_call_void_int_color_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_color_bool ...godot_method_bind_ptrcall returned\n");

}
float godot_call_float_int(godot_object *instance, godot_method_bind *mb,int arg0) {
    godot_real ret;
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_float_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_float_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_object_vector2(godot_object *instance, godot_method_bind *mb,int arg0,godot_object *arg1,godot_vector2 *arg2) {
    const void *c_args[] = {&arg0,arg1,arg2,};
    printf("CGO: godot_call_void_int_object_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_object_vector2 ...godot_method_bind_ptrcall returned\n");

}
godot_pool_vector3_array *godot_call_poolvector3array_vector3_vector3_bool(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0,godot_vector3 *arg1,bool arg2) {
    godot_pool_vector3_array *ret = new_godot_pool_vector3_array();
    const void *c_args[] = {arg0,arg1,&arg2,};
    printf("CGO: godot_call_poolvector3array_vector3_vector3_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_poolvector3array_vector3_vector3_bool ...godot_method_bind_ptrcall returned\n");
    return ret;
}
int godot_call_int_bool(godot_object *instance, godot_method_bind *mb,bool arg0) {
    godot_int ret;
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_int_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_bool ...godot_method_bind_ptrcall returned\n");
    return ret;
}
bool godot_call_bool_vector2_float_object(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,float arg1,godot_object *arg2) {
    godot_bool ret;
    const void *c_args[] = {arg0,&arg1,arg2,};
    printf("CGO: godot_call_bool_vector2_float_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_vector2_float_object ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void(godot_object *instance, godot_method_bind *mb) {
    const void *c_args[] = {};
    printf("CGO: godot_call_void calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_int_vector3_float(godot_object *instance, godot_method_bind *mb,int arg0,godot_vector3 *arg1,float arg2) {
    const void *c_args[] = {&arg0,arg1,&arg2,};
    printf("CGO: godot_call_void_int_vector3_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_vector3_float ...godot_method_bind_ptrcall returned\n");

}
godot_array *godot_call_array_int(godot_object *instance, godot_method_bind *mb,int arg0) {
    godot_array *ret = new_godot_array();
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_array_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_array_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_vector3_vector3_vector3(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0,godot_vector3 *arg1,godot_vector3 *arg2) {
    const void *c_args[] = {arg0,arg1,arg2,};
    printf("CGO: godot_call_void_vector3_vector3_vector3 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_vector3_vector3_vector3 ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rid_int_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1,godot_rid *arg2) {
    const void *c_args[] = {arg0,&arg1,arg2,};
    printf("CGO: godot_call_void_rid_int_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_int_rid ...godot_method_bind_ptrcall returned\n");

}
godot_rid *godot_call_rid_vector2_vector2_vector2_rid_rid(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_vector2 *arg1,godot_vector2 *arg2,godot_rid *arg3,godot_rid *arg4) {
    godot_rid *ret = new_godot_rid();
    const void *c_args[] = {arg0,arg1,arg2,arg3,arg4,};
    printf("CGO: godot_call_rid_vector2_vector2_vector2_rid_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_rid_vector2_vector2_vector2_rid_rid ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_string_string_variant(godot_object *instance, godot_method_bind *mb,int arg0,char *arg1,char *arg2,godot_variant *arg3) {
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    godot_string gArg2;
    api->godot_string_new(&gArg2);
    api->godot_string_parse_utf8(&gArg2, arg2);
    const void *c_args[] = {&arg0,&gArg1,&gArg2,arg3,};
    printf("CGO: godot_call_void_int_string_string_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_string_string_variant ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);
    godot_string_destroy(&gArg2);

}
void godot_call_void_object_vector2_color_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_vector2 *arg1,godot_color *arg2,godot_object *arg3) {
    const void *c_args[] = {arg0,arg1,arg2,arg3,};
    printf("CGO: godot_call_void_object_vector2_color_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_vector2_color_object ...godot_method_bind_ptrcall returned\n");

}
godot_array *godot_call_array_nodepath(godot_object *instance, godot_method_bind *mb,godot_node_path *arg0) {
    godot_array *ret = new_godot_array();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_array_nodepath calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_array_nodepath ...godot_method_bind_ptrcall returned\n");
    return ret;
}
float godot_call_float_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1) {
    godot_real ret;
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_float_rid_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_float_rid_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
int godot_call_int_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1) {
    godot_int ret;
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_int_rid_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_rid_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_array(godot_object *instance, godot_method_bind *mb,godot_array *arg0) {
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_void_array calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_array ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_int_int_rect2_vector2_float(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,godot_rect2 *arg2,godot_vector2 *arg3,float arg4) {
    const void *c_args[] = {&arg0,&arg1,arg2,arg3,&arg4,};
    printf("CGO: godot_call_void_int_int_rect2_vector2_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_int_rect2_vector2_float ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rid_rid_rid_rid_int_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rid *arg1,godot_rid *arg2,godot_rid *arg3,int arg4,int arg5) {
    const void *c_args[] = {arg0,arg1,arg2,arg3,&arg4,&arg5,};
    printf("CGO: godot_call_void_rid_rid_rid_rid_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_rid_rid_rid_int_int ...godot_method_bind_ptrcall returned\n");

}
godot_array *godot_call_array_string_int_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,int arg2) {
    godot_array *ret = new_godot_array();
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,&arg2,};
    printf("CGO: godot_call_array_string_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_array_string_int_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
bool godot_call_bool_vector3(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0) {
    godot_bool ret;
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_bool_vector3 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_vector3 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_object_rect2(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_rect2 *arg1) {
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_void_object_rect2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_rect2 ...godot_method_bind_ptrcall returned\n");

}
int godot_call_int_int_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,int arg2) {
    godot_int ret;
    const void *c_args[] = {&arg0,&arg1,&arg2,};
    printf("CGO: godot_call_int_int_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_int_int_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_object_rect2_vector2(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_rect2 *arg1,godot_vector2 *arg2) {
    const void *c_args[] = {arg0,arg1,arg2,};
    printf("CGO: godot_call_void_object_rect2_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_rect2_vector2 ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_int_string_variant(godot_object *instance, godot_method_bind *mb,int arg0,char *arg1,godot_variant *arg2) {
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&arg0,&gArg1,arg2,};
    printf("CGO: godot_call_void_int_string_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_string_variant ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);

}
godot_rid *godot_call_rid_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0) {
    godot_rid *ret = new_godot_rid();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_rid_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_rid_rid ...godot_method_bind_ptrcall returned\n");
    return ret;
}
bool godot_call_bool_string_object_string(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1,char *arg2) {
    godot_bool ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg2;
    api->godot_string_new(&gArg2);
    api->godot_string_parse_utf8(&gArg2, arg2);
    const void *c_args[] = {&gArg0,arg1,&gArg2,};
    printf("CGO: godot_call_bool_string_object_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_string_object_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg2);
    return ret;
}
godot_rid *godot_call_rid_int_bool(godot_object *instance, godot_method_bind *mb,int arg0,bool arg1) {
    godot_rid *ret = new_godot_rid();
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_rid_int_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_rid_int_bool ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_rid_string_variant(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,char *arg1,godot_variant *arg2) {
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {arg0,&gArg1,arg2,};
    printf("CGO: godot_call_void_rid_string_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_string_variant ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);

}
godot_vector2 *godot_call_vector2_string(godot_object *instance, godot_method_bind *mb,char *arg0) {
    godot_vector2 *ret = new_godot_vector2();
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,};
    printf("CGO: godot_call_vector2_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector2_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
int godot_call_int_int_string_poolstringarray_string(godot_object *instance, godot_method_bind *mb,int arg0,char *arg1,godot_pool_string_array *arg2,char *arg3) {
    godot_int ret;
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    godot_string gArg3;
    api->godot_string_new(&gArg3);
    api->godot_string_parse_utf8(&gArg3, arg3);
    const void *c_args[] = {&arg0,&gArg1,arg2,&gArg3,};
    printf("CGO: godot_call_int_int_string_poolstringarray_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_int_string_poolstringarray_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);
    godot_string_destroy(&gArg3);
    return ret;
}
godot_vector2 *godot_call_vector2_int_int_object_vector2(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,godot_object *arg2,godot_vector2 *arg3) {
    godot_vector2 *ret = new_godot_vector2();
    const void *c_args[] = {&arg0,&arg1,arg2,arg3,};
    printf("CGO: godot_call_vector2_int_int_object_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector2_int_int_object_vector2 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,int arg2) {
    const void *c_args[] = {&arg0,&arg1,&arg2,};
    printf("CGO: godot_call_void_int_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_int_int ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_plane(godot_object *instance, godot_method_bind *mb,godot_plane *arg0) {
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_void_plane calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_plane ...godot_method_bind_ptrcall returned\n");

}
bool godot_call_bool_vector2_variant_object(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_variant *arg1,godot_object *arg2) {
    godot_bool ret;
    const void *c_args[] = {arg0,arg1,arg2,};
    printf("CGO: godot_call_bool_vector2_variant_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_vector2_variant_object ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_object *godot_call_object_int(godot_object *instance, godot_method_bind *mb,int arg0) {
    godot_object *ret;
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_object_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_object_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
int godot_call_int_nodepath(godot_object *instance, godot_method_bind *mb,godot_node_path *arg0) {
    godot_int ret;
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_int_nodepath calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_nodepath ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_pool_vector2_array *godot_call_poolvector2array(godot_object *instance, godot_method_bind *mb) {
    godot_pool_vector2_array *ret = new_godot_pool_vector2_array();
    const void *c_args[] = {};
    printf("CGO: godot_call_poolvector2array calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_poolvector2array ...godot_method_bind_ptrcall returned\n");
    return ret;
}
bool godot_call_bool_string_string(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1) {
    godot_bool ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&gArg0,&gArg1,};
    printf("CGO: godot_call_bool_string_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_string_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg1);
    return ret;
}
void godot_call_void_string_string_object_object(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,godot_object *arg2,godot_object *arg3) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&gArg0,&gArg1,arg2,arg3,};
    printf("CGO: godot_call_void_string_string_object_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_string_object_object ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg1);

}
godot_vector2 *godot_call_vector2_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1) {
    godot_vector2 *ret = new_godot_vector2();
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,};
    printf("CGO: godot_call_vector2_string_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector2_string_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
int godot_call_int_string_string_dictionary_array_array(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,godot_dictionary *arg2,godot_array *arg3,godot_array *arg4) {
    godot_int ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&gArg0,&gArg1,arg2,arg3,arg4,};
    printf("CGO: godot_call_int_string_string_dictionary_array_array calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_string_string_dictionary_array_array ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg1);
    return ret;
}
void godot_call_void_string_object_string_variant(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1,char *arg2,godot_variant *arg3) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg2;
    api->godot_string_new(&gArg2);
    api->godot_string_parse_utf8(&gArg2, arg2);
    const void *c_args[] = {&gArg0,arg1,&gArg2,arg3,};
    printf("CGO: godot_call_void_string_object_string_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_object_string_variant ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg2);

}
godot_pool_vector2_array *godot_call_poolvector2array_vector2_vector2_bool(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_vector2 *arg1,bool arg2) {
    godot_pool_vector2_array *ret = new_godot_pool_vector2_array();
    const void *c_args[] = {arg0,arg1,&arg2,};
    printf("CGO: godot_call_poolvector2array_vector2_vector2_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_poolvector2array_vector2_vector2_bool ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_rid *godot_call_rid_rid_string(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,char *arg1) {
    godot_rid *ret = new_godot_rid();
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {arg0,&gArg1,};
    printf("CGO: godot_call_rid_rid_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_rid_rid_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);
    return ret;
}
int godot_call_int_int_int_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,int arg2,int arg3) {
    godot_int ret;
    const void *c_args[] = {&arg0,&arg1,&arg2,&arg3,};
    printf("CGO: godot_call_int_int_int_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_int_int_int_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_array *godot_call_array_object_vector3(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_vector3 *arg1) {
    godot_array *ret = new_godot_array();
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_array_object_vector3 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_array_object_vector3 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
bool godot_call_bool_transform2d_object_transform2d(godot_object *instance, godot_method_bind *mb,godot_transform2d *arg0,godot_object *arg1,godot_transform2d *arg2) {
    godot_bool ret;
    const void *c_args[] = {arg0,arg1,arg2,};
    printf("CGO: godot_call_bool_transform2d_object_transform2d calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_transform2d_object_transform2d ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_poolvector2array_color_float_bool(godot_object *instance, godot_method_bind *mb,godot_pool_vector2_array *arg0,godot_color *arg1,float arg2,bool arg3) {
    const void *c_args[] = {arg0,arg1,&arg2,&arg3,};
    printf("CGO: godot_call_void_poolvector2array_color_float_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_poolvector2array_color_float_bool ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_int_array(godot_object *instance, godot_method_bind *mb,int arg0,godot_array *arg1) {
    const void *c_args[] = {&arg0,arg1,};
    printf("CGO: godot_call_void_int_array calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_array ...godot_method_bind_ptrcall returned\n");

}
godot_node_path *godot_call_nodepath_int_bool(godot_object *instance, godot_method_bind *mb,int arg0,bool arg1) {
    godot_node_path *ret = new_godot_node_path();
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_nodepath_int_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_nodepath_int_bool ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_variant *godot_call_variant_vector2_object(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_object *arg1) {
    godot_variant *ret = new_godot_variant();
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_variant_vector2_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_variant_vector2_object ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_string_nodepath_bool(godot_object *instance, godot_method_bind *mb,char *arg0,godot_node_path *arg1,bool arg2) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,arg1,&arg2,};
    printf("CGO: godot_call_void_string_nodepath_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_nodepath_bool ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
godot_transform2d *godot_call_transform2d_int(godot_object *instance, godot_method_bind *mb,int arg0) {
    godot_transform2d *ret = new_godot_transform2d();
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_transform2d_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_transform2d_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_int_bool_int_poolbytearray(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,bool arg2,int arg3,godot_pool_byte_array *arg4) {
    const void *c_args[] = {&arg0,&arg1,&arg2,&arg3,arg4,};
    printf("CGO: godot_call_void_int_int_bool_int_poolbytearray calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_int_bool_int_poolbytearray ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_poolvector2array_poolintarray(godot_object *instance, godot_method_bind *mb,godot_pool_vector2_array *arg0,godot_pool_int_array *arg1) {
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_void_poolvector2array_poolintarray calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_poolvector2array_poolintarray ...godot_method_bind_ptrcall returned\n");

}
const char *godot_call_string_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0) {
    godot_string *ret = new_godot_string();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_string_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_string_rid ...godot_method_bind_ptrcall returned\n");
    godot_char_string gcsRet = godot_string_utf8(ret);
    const char *cRet = api->godot_char_string_get_data(&gcsRet);
    api->godot_char_string_destroy(&gcsRet);
    return cRet;

}
godot_dictionary *godot_call_dictionary_vector2_vector2_array_int(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_vector2 *arg1,godot_array *arg2,int arg3) {
    godot_dictionary *ret = new_godot_dictionary();
    const void *c_args[] = {arg0,arg1,arg2,&arg3,};
    printf("CGO: godot_call_dictionary_vector2_vector2_array_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_dictionary_vector2_vector2_array_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_int_variant(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,godot_variant *arg2) {
    const void *c_args[] = {&arg0,&arg1,arg2,};
    printf("CGO: godot_call_void_int_int_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_int_variant ...godot_method_bind_ptrcall returned\n");

}
godot_vector2 *godot_call_vector2_vector3(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0) {
    godot_vector2 *ret = new_godot_vector2();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_vector2_vector3 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector2_vector3 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_vector3 *godot_call_vector3_float(godot_object *instance, godot_method_bind *mb,float arg0) {
    godot_vector3 *ret = new_godot_vector3();
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_vector3_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector3_float ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_int_poolstringarray_poolbytearray(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,godot_pool_string_array *arg2,godot_pool_byte_array *arg3) {
    const void *c_args[] = {&arg0,&arg1,arg2,arg3,};
    printf("CGO: godot_call_void_int_int_poolstringarray_poolbytearray calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_int_poolstringarray_poolbytearray ...godot_method_bind_ptrcall returned\n");

}
int godot_call_int_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1) {
    godot_int ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,};
    printf("CGO: godot_call_int_string_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_string_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
void godot_call_void_object_object_rect2_vector2(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_object *arg1,godot_rect2 *arg2,godot_vector2 *arg3) {
    const void *c_args[] = {arg0,arg1,arg2,arg3,};
    printf("CGO: godot_call_void_object_object_rect2_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_object_rect2_vector2 ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_float(godot_object *instance, godot_method_bind *mb,float arg0) {
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_void_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_float ...godot_method_bind_ptrcall returned\n");

}
float godot_call_float_string_string(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1) {
    godot_real ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&gArg0,&gArg1,};
    printf("CGO: godot_call_float_string_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_float_string_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg1);
    return ret;
}
bool godot_call_bool_string_object(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1) {
    godot_bool ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,arg1,};
    printf("CGO: godot_call_bool_string_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_string_object ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
godot_dictionary *godot_call_dictionary_vector3_vector3_array_int(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0,godot_vector3 *arg1,godot_array *arg2,int arg3) {
    godot_dictionary *ret = new_godot_dictionary();
    const void *c_args[] = {arg0,arg1,arg2,&arg3,};
    printf("CGO: godot_call_dictionary_vector3_vector3_array_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_dictionary_vector3_vector3_array_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_int_transform2d(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,godot_transform2d *arg2) {
    const void *c_args[] = {&arg0,&arg1,arg2,};
    printf("CGO: godot_call_void_int_int_transform2d calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_int_transform2d ...godot_method_bind_ptrcall returned\n");

}
bool godot_call_bool_object_nodepath_object_nodepath_variant_float_int_int_float(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_node_path *arg1,godot_object *arg2,godot_node_path *arg3,godot_variant *arg4,float arg5,int arg6,int arg7,float arg8) {
    godot_bool ret;
    const void *c_args[] = {arg0,arg1,arg2,arg3,arg4,&arg5,&arg6,&arg7,&arg8,};
    printf("CGO: godot_call_bool_object_nodepath_object_nodepath_variant_float_int_int_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_object_nodepath_object_nodepath_variant_float_int_int_float ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_poolbytearray(godot_object *instance, godot_method_bind *mb,godot_pool_byte_array *arg0) {
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_void_poolbytearray calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_poolbytearray ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_object_object_int(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_object *arg1,int arg2) {
    const void *c_args[] = {arg0,arg1,&arg2,};
    printf("CGO: godot_call_void_object_object_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_object_int ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_vector2_variant(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_variant *arg1) {
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_void_vector2_variant calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_vector2_variant ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rid_transform(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_transform *arg1) {
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_void_rid_transform calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_transform ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rid_rect2_bool_color_bool_object(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rect2 *arg1,bool arg2,godot_color *arg3,bool arg4,godot_object *arg5) {
    const void *c_args[] = {arg0,arg1,&arg2,arg3,&arg4,arg5,};
    printf("CGO: godot_call_void_rid_rect2_bool_color_bool_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_rect2_bool_color_bool_object ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rid_string(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,char *arg1) {
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {arg0,&gArg1,};
    printf("CGO: godot_call_void_rid_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);

}
bool godot_call_bool_bool(godot_object *instance, godot_method_bind *mb,bool arg0) {
    godot_bool ret;
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_bool_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_bool ...godot_method_bind_ptrcall returned\n");
    return ret;
}
int godot_call_int_int_string_poolstringarray_poolbytearray(godot_object *instance, godot_method_bind *mb,int arg0,char *arg1,godot_pool_string_array *arg2,godot_pool_byte_array *arg3) {
    godot_int ret;
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&arg0,&gArg1,arg2,arg3,};
    printf("CGO: godot_call_int_int_string_poolstringarray_poolbytearray calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_int_string_poolstringarray_poolbytearray ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg1);
    return ret;
}
godot_object *godot_call_object_vector3(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0) {
    godot_object *ret;
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_object_vector3 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_object_vector3 ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_object *godot_call_object_string_bool_bool(godot_object *instance, godot_method_bind *mb,char *arg0,bool arg1,bool arg2) {
    godot_object *ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,&arg2,};
    printf("CGO: godot_call_object_string_bool_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_object_string_bool_bool ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
void godot_call_void_rid_transform2d(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_transform2d *arg1) {
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_void_rid_transform2d calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_transform2d ...godot_method_bind_ptrcall returned\n");

}
godot_array *godot_call_array_poolbytearray(godot_object *instance, godot_method_bind *mb,godot_pool_byte_array *arg0) {
    godot_array *ret = new_godot_array();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_array_poolbytearray calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_array_poolbytearray ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_vector3 *godot_call_vector3_int_float(godot_object *instance, godot_method_bind *mb,int arg0,float arg1) {
    godot_vector3 *ret = new_godot_vector3();
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_vector3_int_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector3_int_float ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_rid_rect2_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rect2 *arg1,int arg2) {
    const void *c_args[] = {arg0,arg1,&arg2,};
    printf("CGO: godot_call_void_rid_rect2_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_rect2_int ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_poolintarray(godot_object *instance, godot_method_bind *mb,godot_pool_int_array *arg0) {
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_void_poolintarray calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_poolintarray ...godot_method_bind_ptrcall returned\n");

}
float godot_call_float(godot_object *instance, godot_method_bind *mb) {
    godot_real ret;
    const void *c_args[] = {};
    printf("CGO: godot_call_float calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_float ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_string_bool(godot_object *instance, godot_method_bind *mb,char *arg0,bool arg1) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,};
    printf("CGO: godot_call_void_string_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_bool ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
void godot_call_void_int_bool_int(godot_object *instance, godot_method_bind *mb,int arg0,bool arg1,int arg2) {
    const void *c_args[] = {&arg0,&arg1,&arg2,};
    printf("CGO: godot_call_void_int_bool_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_bool_int ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rid_poolvector2array_poolcolorarray_poolvector2array_rid_float_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_pool_vector2_array *arg1,godot_pool_color_array *arg2,godot_pool_vector2_array *arg3,godot_rid *arg4,float arg5,godot_rid *arg6) {
    const void *c_args[] = {arg0,arg1,arg2,arg3,arg4,&arg5,arg6,};
    printf("CGO: godot_call_void_rid_poolvector2array_poolcolorarray_poolvector2array_rid_float_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_poolvector2array_poolcolorarray_poolvector2array_rid_float_rid ...godot_method_bind_ptrcall returned\n");

}
godot_aabb *godot_call_aabb(godot_object *instance, godot_method_bind *mb) {
    godot_aabb *ret = new_godot_aabb();
    const void *c_args[] = {};
    printf("CGO: godot_call_aabb calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_aabb ...godot_method_bind_ptrcall returned\n");
    return ret;
}
int godot_call_int_string(godot_object *instance, godot_method_bind *mb,char *arg0) {
    godot_int ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,};
    printf("CGO: godot_call_int_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
godot_pool_int_array *godot_call_poolintarray(godot_object *instance, godot_method_bind *mb) {
    godot_pool_int_array *ret = new_godot_pool_int_array();
    const void *c_args[] = {};
    printf("CGO: godot_call_poolintarray calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_poolintarray ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_transform *godot_call_transform_int(godot_object *instance, godot_method_bind *mb,int arg0) {
    godot_transform *ret = new_godot_transform();
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_transform_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_transform_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_float_float_bool(godot_object *instance, godot_method_bind *mb,int arg0,float arg1,float arg2,bool arg3) {
    const void *c_args[] = {&arg0,&arg1,&arg2,&arg3,};
    printf("CGO: godot_call_void_int_float_float_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_float_float_bool ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_bool_bool_int_int(godot_object *instance, godot_method_bind *mb,bool arg0,bool arg1,int arg2,int arg3) {
    const void *c_args[] = {&arg0,&arg1,&arg2,&arg3,};
    printf("CGO: godot_call_void_bool_bool_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_bool_bool_int_int ...godot_method_bind_ptrcall returned\n");

}
godot_dictionary *godot_call_dictionary_string(godot_object *instance, godot_method_bind *mb,char *arg0) {
    godot_dictionary *ret = new_godot_dictionary();
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,};
    printf("CGO: godot_call_dictionary_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_dictionary_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
void godot_call_void_object_bool_rid(godot_object *instance, godot_method_bind *mb,godot_object *arg0,bool arg1,godot_rid *arg2) {
    const void *c_args[] = {arg0,&arg1,arg2,};
    printf("CGO: godot_call_void_object_bool_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_bool_rid ...godot_method_bind_ptrcall returned\n");

}
godot_vector3 *godot_call_vector3_vector3_vector3_bool(godot_object *instance, godot_method_bind *mb,godot_vector3 *arg0,godot_vector3 *arg1,bool arg2) {
    godot_vector3 *ret = new_godot_vector3();
    const void *c_args[] = {arg0,arg1,&arg2,};
    printf("CGO: godot_call_vector3_vector3_vector3_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector3_vector3_vector3_bool ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_object_int_bool(godot_object *instance, godot_method_bind *mb,godot_object *arg0,int arg1,bool arg2) {
    const void *c_args[] = {arg0,&arg1,&arg2,};
    printf("CGO: godot_call_void_object_int_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_int_bool ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_string_int_int_int_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,int arg2,int arg3,int arg4) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,&arg2,&arg3,&arg4,};
    printf("CGO: godot_call_void_string_int_int_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_int_int_int_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
godot_vector3 *godot_call_vector3_int(godot_object *instance, godot_method_bind *mb,int arg0) {
    godot_vector3 *ret = new_godot_vector3();
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_vector3_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector3_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_object *godot_call_object(godot_object *instance, godot_method_bind *mb) {
    godot_object *ret;
    const void *c_args[] = {};
    printf("CGO: godot_call_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_object ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_transform2d *godot_call_transform2d_object(godot_object *instance, godot_method_bind *mb,godot_object *arg0) {
    godot_transform2d *ret = new_godot_transform2d();
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_transform2d_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_transform2d_object ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_string_object_string(godot_object *instance, godot_method_bind *mb,char *arg0,godot_object *arg1,char *arg2) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg2;
    api->godot_string_new(&gArg2);
    api->godot_string_parse_utf8(&gArg2, arg2);
    const void *c_args[] = {&gArg0,arg1,&gArg2,};
    printf("CGO: godot_call_void_string_object_string calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_object_string ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg2);

}
godot_transform *godot_call_transform_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1) {
    godot_transform *ret = new_godot_transform();
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_transform_rid_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_transform_rid_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_object_vector2_string_color_int(godot_object *instance, godot_method_bind *mb,godot_object *arg0,godot_vector2 *arg1,char *arg2,godot_color *arg3,int arg4) {
    godot_string gArg2;
    api->godot_string_new(&gArg2);
    api->godot_string_parse_utf8(&gArg2, arg2);
    const void *c_args[] = {arg0,arg1,&gArg2,arg3,&arg4,};
    printf("CGO: godot_call_void_object_vector2_string_color_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_object_vector2_string_color_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg2);

}
void godot_call_void_int_vector2(godot_object *instance, godot_method_bind *mb,int arg0,godot_vector2 *arg1) {
    const void *c_args[] = {&arg0,arg1,};
    printf("CGO: godot_call_void_int_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_vector2 ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_vector2_vector2(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_vector2 *arg1) {
    const void *c_args[] = {arg0,arg1,};
    printf("CGO: godot_call_void_vector2_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_vector2_vector2 ...godot_method_bind_ptrcall returned\n");

}
godot_object *godot_call_object_rid_int(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,int arg1) {
    godot_object *ret;
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_object_rid_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_object_rid_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_vector2_vector2_vector2_int(godot_object *instance, godot_method_bind *mb,godot_vector2 *arg0,godot_vector2 *arg1,godot_vector2 *arg2,int arg3) {
    const void *c_args[] = {arg0,arg1,arg2,&arg3,};
    printf("CGO: godot_call_void_vector2_vector2_vector2_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_vector2_vector2_vector2_int ...godot_method_bind_ptrcall returned\n");

}
int godot_call_int_string_int_string_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,char *arg2,int arg3) {
    godot_int ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg2;
    api->godot_string_new(&gArg2);
    api->godot_string_parse_utf8(&gArg2, arg2);
    const void *c_args[] = {&gArg0,&arg1,&gArg2,&arg3,};
    printf("CGO: godot_call_int_string_int_string_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_int_string_int_string_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg2);
    return ret;
}
bool godot_call_bool_string_int_int_int_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,int arg2,int arg3,int arg4) {
    godot_bool ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,&arg2,&arg3,&arg4,};
    printf("CGO: godot_call_bool_string_int_int_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_bool_string_int_int_int_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
godot_variant *godot_call_variant_string_string_array(godot_object *instance, godot_method_bind *mb,char *arg0,char *arg1,godot_array *arg2) {
    godot_variant *ret = new_godot_variant();
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    godot_string gArg1;
    api->godot_string_new(&gArg1);
    api->godot_string_parse_utf8(&gArg1, arg1);
    const void *c_args[] = {&gArg0,&gArg1,arg2,};
    printf("CGO: godot_call_variant_string_string_array calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_variant_string_string_array ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    godot_string_destroy(&gArg1);
    return ret;
}
godot_dictionary *godot_call_dictionary_int(godot_object *instance, godot_method_bind *mb,int arg0) {
    godot_dictionary *ret = new_godot_dictionary();
    const void *c_args[] = {&arg0,};
    printf("CGO: godot_call_dictionary_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_dictionary_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_int_object(godot_object *instance, godot_method_bind *mb,int arg0,godot_object *arg1) {
    const void *c_args[] = {&arg0,arg1,};
    printf("CGO: godot_call_void_int_object calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_int_object ...godot_method_bind_ptrcall returned\n");

}
godot_array *godot_call_array_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1) {
    godot_array *ret = new_godot_array();
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_array_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_array_int_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_float_bool(godot_object *instance, godot_method_bind *mb,float arg0,bool arg1) {
    const void *c_args[] = {&arg0,&arg1,};
    printf("CGO: godot_call_void_float_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_float_bool ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_string_int_bool(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,bool arg2) {
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,&arg2,};
    printf("CGO: godot_call_void_string_int_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_string_int_bool ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);

}
void godot_call_void_rect2(godot_object *instance, godot_method_bind *mb,godot_rect2 *arg0) {
    const void *c_args[] = {arg0,};
    printf("CGO: godot_call_void_rect2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rect2 ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rect2_bool(godot_object *instance, godot_method_bind *mb,godot_rect2 *arg0,bool arg1) {
    const void *c_args[] = {arg0,&arg1,};
    printf("CGO: godot_call_void_rect2_bool calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rect2_bool ...godot_method_bind_ptrcall returned\n");

}
void godot_call_void_rid_rect2_rect2_rid_vector2_vector2_int_int_bool_color_rid(godot_object *instance, godot_method_bind *mb,godot_rid *arg0,godot_rect2 *arg1,godot_rect2 *arg2,godot_rid *arg3,godot_vector2 *arg4,godot_vector2 *arg5,int arg6,int arg7,bool arg8,godot_color *arg9,godot_rid *arg10) {
    const void *c_args[] = {arg0,arg1,arg2,arg3,arg4,arg5,&arg6,&arg7,&arg8,arg9,arg10,};
    printf("CGO: godot_call_void_rid_rect2_rect2_rid_vector2_vector2_int_int_bool_color_rid calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_rid_rect2_rect2_rid_vector2_vector2_int_int_bool_color_rid ...godot_method_bind_ptrcall returned\n");

}
godot_transform *godot_call_transform(godot_object *instance, godot_method_bind *mb) {
    godot_transform *ret = new_godot_transform();
    const void *c_args[] = {};
    printf("CGO: godot_call_transform calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_transform ...godot_method_bind_ptrcall returned\n");
    return ret;
}
godot_object *godot_call_object_string_int_int(godot_object *instance, godot_method_bind *mb,char *arg0,int arg1,int arg2) {
    godot_object *ret;
    godot_string gArg0;
    api->godot_string_new(&gArg0);
    api->godot_string_parse_utf8(&gArg0, arg0);
    const void *c_args[] = {&gArg0,&arg1,&arg2,};
    printf("CGO: godot_call_object_string_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,&ret);
    printf("CGO: godot_call_object_string_int_int ...godot_method_bind_ptrcall returned\n");
    godot_string_destroy(&gArg0);
    return ret;
}
godot_vector3 *godot_call_vector3_int_int_int(godot_object *instance, godot_method_bind *mb,int arg0,int arg1,int arg2) {
    godot_vector3 *ret = new_godot_vector3();
    const void *c_args[] = {&arg0,&arg1,&arg2,};
    printf("CGO: godot_call_vector3_int_int_int calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args,ret);
    printf("CGO: godot_call_vector3_int_int_int ...godot_method_bind_ptrcall returned\n");
    return ret;
}
void godot_call_void_bool_vector2_vector2(godot_object *instance, godot_method_bind *mb,bool arg0,godot_vector2 *arg1,godot_vector2 *arg2) {
    const void *c_args[] = {&arg0,arg1,arg2,};
    printf("CGO: godot_call_void_bool_vector2_vector2 calling godot_method_bind_ptrcall...\n");
    api->godot_method_bind_ptrcall(mb, instance, c_args, NULL);
    printf("CGO: godot_call_void_bool_vector2_vector2 ...godot_method_bind_ptrcall returned\n");

}


*/
import "C"

import (
	"fmt"
	"log"

	"unsafe"
)

var methodBinds map[string]map[string]*C.godot_method_bind = make(map[string]map[string]*C.godot_method_bind)

func lookupGodotMethod(baseClass string, methodName string) *C.godot_method_bind {
	// Convert the base class and method names to C strings.
	classCString := C.CString(baseClass)
	defer C.free(unsafe.Pointer(classCString))
	methodCString := C.CString(methodName)
	defer C.free(unsafe.Pointer(methodCString))

	// Get the Godot method bind pointer
	var methodBind *C.godot_method_bind
	methodBind = C.godot_method_bind_get_method(classCString, methodCString)
	log.Println("  Found method bind pointer:", fmt.Sprintf("0x%x", uintptr(unsafe.Pointer(methodBind))), "for", baseClass, "::", methodName)
	return methodBind
}

func getGodotMethod(baseClass string, methodName string) *C.godot_method_bind {
	if cms, found := methodBinds[baseClass]; found {
		if mb, found := cms[methodName]; found {
			return mb
		}
		mb := lookupGodotMethod(baseClass, methodName)
		cms[methodName] = mb
		return mb
	}
	mb := lookupGodotMethod(baseClass, methodName)
	methodBinds[baseClass] = map[string]*C.godot_method_bind{methodName: mb}
	return mb
}

func godotCallVoidVector2Vector2Vector2Int(o Class, methodName string, arg0 *Vector2, arg1 *Vector2, arg2 *Vector2, arg3 int64) {

	cArg0 := arg0.vector2
	cArg1 := arg1.vector2
	cArg2 := arg2.vector2
	cArg3 := C.godot_int(arg3)

	log.Println("  Calling godot_call_void_vector2_vector2_vector2_int...")
	C.godot_call_void_vector2_vector2_vector2_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_vector2_vector2_vector2_int returned")

}
func godotCallIntStringIntStringInt(o Class, methodName string, arg0 string, arg1 int64, arg2 string, arg3 int64) int64 {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)
	cArg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(cArg2))
	cArg3 := C.godot_int(arg3)

	log.Println("  Calling godot_call_int_string_int_string_int...")
	cRet := C.godot_call_int_string_int_string_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_int_string_int_string_int returned")
	ret := cRet
	return int64(ret)

}
func godotCallBoolStringIntIntIntInt(o Class, methodName string, arg0 string, arg1 int64, arg2 int64, arg3 int64, arg4 int64) bool {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_int(arg2)
	cArg3 := C.godot_int(arg3)
	cArg4 := C.godot_int(arg4)

	log.Println("  Calling godot_call_bool_string_int_int_int_int...")
	cRet := C.godot_call_bool_string_int_int_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_bool_string_int_int_int_int returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoidIntObject(o Class, methodName string, arg0 int64, arg1 *Object) {

	cArg0 := C.godot_int(arg0)
	cArg1 := unsafe.Pointer(arg1.owner)

	log.Println("  Calling godot_call_void_int_object...")
	C.godot_call_void_int_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_int_object returned")

}
func godotCallArrayIntInt(o Class, methodName string, arg0 int64, arg1 int64) *Array {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_array_int_int...")
	cRet := C.godot_call_array_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_array_int_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_array)(cRet)
	return &Array{ret}

}
func godotCallVoidFloatBool(o Class, methodName string, arg0 float64, arg1 bool) {

	cArg0 := C.godot_real(arg0)
	cArg1 := C.godot_bool(arg1)

	log.Println("  Calling godot_call_void_float_bool...")
	C.godot_call_void_float_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_float_bool returned")

}
func godotCallVoidStringIntBool(o Class, methodName string, arg0 string, arg1 int64, arg2 bool) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_void_string_int_bool...")
	C.godot_call_void_string_int_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_string_int_bool returned")

}
func godotCallVoidRect2(o Class, methodName string, arg0 *Rect2) {

	cArg0 := arg0.rect2

	log.Println("  Calling godot_call_void_rect2...")
	C.godot_call_void_rect2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_rect2 returned")

}
func godotCallVoidRect2Bool(o Class, methodName string, arg0 *Rect2, arg1 bool) {

	cArg0 := arg0.rect2
	cArg1 := C.godot_bool(arg1)

	log.Println("  Calling godot_call_void_rect2_bool...")
	C.godot_call_void_rect2_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_rect2_bool returned")

}
func godotCallVariantStringStringArray(o Class, methodName string, arg0 string, arg1 string, arg2 *Array) *Variant {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := arg2.array

	log.Println("  Calling godot_call_variant_string_string_array...")
	cRet := C.godot_call_variant_string_string_array(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_variant_string_string_array returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_variant)(cRet)
	return &Variant{ret}

}
func godotCallDictionaryInt(o Class, methodName string, arg0 int64) *Dictionary {

	cArg0 := C.godot_int(arg0)

	log.Println("  Calling godot_call_dictionary_int...")
	cRet := C.godot_call_dictionary_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_dictionary_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_dictionary)(cRet)
	return &Dictionary{ret}

}
func godotCallVoidRidRect2Rect2RidVector2Vector2IntIntBoolColorRid(o Class, methodName string, arg0 *RID, arg1 *Rect2, arg2 *Rect2, arg3 *RID, arg4 *Vector2, arg5 *Vector2, arg6 int64, arg7 int64, arg8 bool, arg9 *Color, arg10 *RID) {

	cArg0 := arg0.rid
	cArg1 := arg1.rect2
	cArg2 := arg2.rect2
	cArg3 := arg3.rid
	cArg4 := arg4.vector2
	cArg5 := arg5.vector2
	cArg6 := C.godot_int(arg6)
	cArg7 := C.godot_int(arg7)
	cArg8 := C.godot_bool(arg8)
	cArg9 := arg9.color
	cArg10 := arg10.rid

	log.Println("  Calling godot_call_void_rid_rect2_rect2_rid_vector2_vector2_int_int_bool_color_rid...")
	C.godot_call_void_rid_rect2_rect2_rid_vector2_vector2_int_int_bool_color_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5, cArg6, cArg7, cArg8, cArg9, cArg10,
	)
	log.Println("  ...godot_call_void_rid_rect2_rect2_rid_vector2_vector2_int_int_bool_color_rid returned")

}
func godotCallTransform(o Class, methodName string) *Transform {

	log.Println("  Calling godot_call_transform...")
	cRet := C.godot_call_transform(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_transform returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_transform)(cRet)
	return &Transform{ret}

}
func godotCallObjectStringIntInt(o Class, methodName string, arg0 string, arg1 int64, arg2 int64) *Object {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_object_string_int_int...")
	cRet := C.godot_call_object_string_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_object_string_int_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_object)(cRet)
	return &Object{ret}

}
func godotCallVector3IntIntInt(o Class, methodName string, arg0 int64, arg1 int64, arg2 int64) *Vector3 {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_vector3_int_int_int...")
	cRet := C.godot_call_vector3_int_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_vector3_int_int_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector3)(cRet)
	return &Vector3{ret}

}
func godotCallVoidBoolVector2Vector2(o Class, methodName string, arg0 bool, arg1 *Vector2, arg2 *Vector2) {

	cArg0 := C.godot_bool(arg0)
	cArg1 := arg1.vector2
	cArg2 := arg2.vector2

	log.Println("  Calling godot_call_void_bool_vector2_vector2...")
	C.godot_call_void_bool_vector2_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_bool_vector2_vector2 returned")

}
func godotCallIntStringPoolStringArrayBoolIntString(o Class, methodName string, arg0 string, arg1 *PoolStringArray, arg2 bool, arg3 int64, arg4 string) int64 {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := arg1.array
	cArg2 := C.godot_bool(arg2)
	cArg3 := C.godot_int(arg3)
	cArg4 := C.CString(arg4)
	defer C.free(unsafe.Pointer(cArg4))

	log.Println("  Calling godot_call_int_string_poolstringarray_bool_int_string...")
	cRet := C.godot_call_int_string_poolstringarray_bool_int_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_int_string_poolstringarray_bool_int_string returned")
	ret := cRet
	return int64(ret)

}
func godotCallVoidVariant(o Class, methodName string, arg0 *Variant) {

	cArg0 := arg0.variant

	log.Println("  Calling godot_call_void_variant...")
	C.godot_call_void_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_variant returned")

}
func godotCallPoolByteArrayRidInt(o Class, methodName string, arg0 *RID, arg1 int64) *PoolByteArray {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_poolbytearray_rid_int...")
	cRet := C.godot_call_poolbytearray_rid_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_poolbytearray_rid_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_pool_byte_array)(cRet)
	return &PoolByteArray{ret}

}
func godotCallVoidVector3(o Class, methodName string, arg0 *Vector3) {

	cArg0 := arg0.vector3

	log.Println("  Calling godot_call_void_vector3...")
	C.godot_call_void_vector3(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_vector3 returned")

}
func godotCallPoolRealArray(o Class, methodName string) *PoolRealArray {

	log.Println("  Calling godot_call_poolrealarray...")
	cRet := C.godot_call_poolrealarray(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_poolrealarray returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_pool_real_array)(cRet)
	return &PoolRealArray{ret}

}
func godotCallVoidPoolStringArrayBoolStringInt(o Class, methodName string, arg0 *PoolStringArray, arg1 bool, arg2 string, arg3 int64) {

	cArg0 := arg0.array
	cArg1 := C.godot_bool(arg1)
	cArg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(cArg2))
	cArg3 := C.godot_int(arg3)

	log.Println("  Calling godot_call_void_poolstringarray_bool_string_int...")
	C.godot_call_void_poolstringarray_bool_string_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_poolstringarray_bool_string_int returned")

}
func godotCallVoidIntBoolIntColorBoolIntColorObjectObject(o Class, methodName string, arg0 int64, arg1 bool, arg2 int64, arg3 *Color, arg4 bool, arg5 int64, arg6 *Color, arg7 *Object, arg8 *Object) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_bool(arg1)
	cArg2 := C.godot_int(arg2)
	cArg3 := arg3.color
	cArg4 := C.godot_bool(arg4)
	cArg5 := C.godot_int(arg5)
	cArg6 := arg6.color
	cArg7 := unsafe.Pointer(arg7.owner)
	cArg8 := unsafe.Pointer(arg8.owner)

	log.Println("  Calling godot_call_void_int_bool_int_color_bool_int_color_object_object...")
	C.godot_call_void_int_bool_int_color_bool_int_color_object_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5, cArg6, cArg7, cArg8,
	)
	log.Println("  ...godot_call_void_int_bool_int_color_bool_int_color_object_object returned")

}
func godotCallIntIntString(o Class, methodName string, arg0 int64, arg1 string) int64 {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))

	log.Println("  Calling godot_call_int_int_string...")
	cRet := C.godot_call_int_int_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_int_int_string returned")
	ret := cRet
	return int64(ret)

}
func godotCallVariantIntInt(o Class, methodName string, arg0 int64, arg1 int64) *Variant {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_variant_int_int...")
	cRet := C.godot_call_variant_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_variant_int_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_variant)(cRet)
	return &Variant{ret}

}
func godotCallFloatString(o Class, methodName string, arg0 string) float64 {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))

	log.Println("  Calling godot_call_float_string...")
	cRet := C.godot_call_float_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_float_string returned")
	ret := cRet
	return float64(ret)

}
func godotCallBoolNodePath(o Class, methodName string, arg0 *NodePath) bool {

	cArg0 := arg0.nodePath

	log.Println("  Calling godot_call_bool_nodepath...")
	cRet := C.godot_call_bool_nodepath(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_bool_nodepath returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoidVector2VariantObject(o Class, methodName string, arg0 *Vector2, arg1 *Variant, arg2 *Object) {

	cArg0 := arg0.vector2
	cArg1 := arg1.variant
	cArg2 := unsafe.Pointer(arg2.owner)

	log.Println("  Calling godot_call_void_vector2_variant_object...")
	C.godot_call_void_vector2_variant_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_vector2_variant_object returned")

}
func godotCallVoidRidIntInt(o Class, methodName string, arg0 *RID, arg1 int64, arg2 int64) {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_void_rid_int_int...")
	C.godot_call_void_rid_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rid_int_int returned")

}
func godotCallVoidRidVector2StringColorInt(o Class, methodName string, arg0 *RID, arg1 *Vector2, arg2 string, arg3 *Color, arg4 int64) {

	cArg0 := arg0.rid
	cArg1 := arg1.vector2
	cArg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(cArg2))
	cArg3 := arg3.color
	cArg4 := C.godot_int(arg4)

	log.Println("  Calling godot_call_void_rid_vector2_string_color_int...")
	C.godot_call_void_rid_vector2_string_color_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_void_rid_vector2_string_color_int returned")

}
func godotCallVoidObjectIntVector2(o Class, methodName string, arg0 *Object, arg1 int64, arg2 *Vector2) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.godot_int(arg1)
	cArg2 := arg2.vector2

	log.Println("  Calling godot_call_void_object_int_vector2...")
	C.godot_call_void_object_int_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_object_int_vector2 returned")

}
func godotCallVoidTransform2DVector2(o Class, methodName string, arg0 *Transform2D, arg1 *Vector2) {

	cArg0 := arg0.transform2d
	cArg1 := arg1.vector2

	log.Println("  Calling godot_call_void_transform2d_vector2...")
	C.godot_call_void_transform2d_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_transform2d_vector2 returned")

}
func godotCallVariantRidInt(o Class, methodName string, arg0 *RID, arg1 int64) *Variant {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_variant_rid_int...")
	cRet := C.godot_call_variant_rid_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_variant_rid_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_variant)(cRet)
	return &Variant{ret}

}
func godotCallPoolStringArray(o Class, methodName string) *PoolStringArray {

	log.Println("  Calling godot_call_poolstringarray...")
	cRet := C.godot_call_poolstringarray(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_poolstringarray returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_pool_string_array)(cRet)
	return &PoolStringArray{ret}

}
func godotCallVoidIntFloatBoolBool(o Class, methodName string, arg0 int64, arg1 float64, arg2 bool, arg3 bool) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_real(arg1)
	cArg2 := C.godot_bool(arg2)
	cArg3 := C.godot_bool(arg3)

	log.Println("  Calling godot_call_void_int_float_bool_bool...")
	C.godot_call_void_int_float_bool_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_int_float_bool_bool returned")

}
func godotCallFloatRidVector2IntIntColor(o Class, methodName string, arg0 *RID, arg1 *Vector2, arg2 int64, arg3 int64, arg4 *Color) float64 {

	cArg0 := arg0.rid
	cArg1 := arg1.vector2
	cArg2 := C.godot_int(arg2)
	cArg3 := C.godot_int(arg3)
	cArg4 := arg4.color

	log.Println("  Calling godot_call_float_rid_vector2_int_int_color...")
	cRet := C.godot_call_float_rid_vector2_int_int_color(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_float_rid_vector2_int_int_color returned")
	ret := cRet
	return float64(ret)

}
func godotCallVoidStringIntObjectVector2(o Class, methodName string, arg0 string, arg1 int64, arg2 *Object, arg3 *Vector2) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)
	cArg2 := unsafe.Pointer(arg2.owner)
	cArg3 := arg3.vector2

	log.Println("  Calling godot_call_void_string_int_object_vector2...")
	C.godot_call_void_string_int_object_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_string_int_object_vector2 returned")

}
func godotCallVoidStringString(o Class, methodName string, arg0 string, arg1 string) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))

	log.Println("  Calling godot_call_void_string_string...")
	C.godot_call_void_string_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_string_string returned")

}
func godotCallStringVector2(o Class, methodName string, arg0 *Vector2) string {

	cArg0 := arg0.vector2

	log.Println("  Calling godot_call_string_vector2...")
	cRet := C.godot_call_string_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_string_vector2 returned")
	return C.GoString(cRet)

}
func godotCallVoidObjectAabb(o Class, methodName string, arg0 *Object, arg1 *AABB) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := arg1.aabb

	log.Println("  Calling godot_call_void_object_aabb...")
	C.godot_call_void_object_aabb(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_object_aabb returned")

}
func godotCallVoidStringIntStringInt(o Class, methodName string, arg0 string, arg1 int64, arg2 string, arg3 int64) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)
	cArg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(cArg2))
	cArg3 := C.godot_int(arg3)

	log.Println("  Calling godot_call_void_string_int_string_int...")
	C.godot_call_void_string_int_string_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_string_int_string_int returned")

}
func godotCallVoidRidRid(o Class, methodName string, arg0 *RID, arg1 *RID) {

	cArg0 := arg0.rid
	cArg1 := arg1.rid

	log.Println("  Calling godot_call_void_rid_rid...")
	C.godot_call_void_rid_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_rid_rid returned")

}
func godotCallTransformBool(o Class, methodName string, arg0 bool) *Transform {

	cArg0 := C.godot_bool(arg0)

	log.Println("  Calling godot_call_transform_bool...")
	cRet := C.godot_call_transform_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_transform_bool returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_transform)(cRet)
	return &Transform{ret}

}
func godotCallIntIntFloat(o Class, methodName string, arg0 int64, arg1 float64) int64 {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_real(arg1)

	log.Println("  Calling godot_call_int_int_float...")
	cRet := C.godot_call_int_int_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_int_int_float returned")
	ret := cRet
	return int64(ret)

}
func godotCallVariantStringArray(o Class, methodName string, arg0 string, arg1 *Array) *Variant {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := arg1.array

	log.Println("  Calling godot_call_variant_string_array...")
	cRet := C.godot_call_variant_string_array(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_variant_string_array returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_variant)(cRet)
	return &Variant{ret}

}
func godotCallVoidRidRect2Color(o Class, methodName string, arg0 *RID, arg1 *Rect2, arg2 *Color) {

	cArg0 := arg0.rid
	cArg1 := arg1.rect2
	cArg2 := arg2.color

	log.Println("  Calling godot_call_void_rid_rect2_color...")
	C.godot_call_void_rid_rect2_color(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rid_rect2_color returned")

}
func godotCallStringIntInt(o Class, methodName string, arg0 int64, arg1 int64) string {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_string_int_int...")
	cRet := C.godot_call_string_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_string_int_int returned")
	return C.GoString(cRet)

}
func godotCallObjectStringString(o Class, methodName string, arg0 string, arg1 string) *Object {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))

	log.Println("  Calling godot_call_object_string_string...")
	cRet := C.godot_call_object_string_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_object_string_string returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_object)(cRet)
	return &Object{ret}

}
func godotCallVariantString(o Class, methodName string, arg0 string) *Variant {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))

	log.Println("  Calling godot_call_variant_string...")
	cRet := C.godot_call_variant_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_variant_string returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_variant)(cRet)
	return &Variant{ret}

}
func godotCallVoidIntBoolStringString(o Class, methodName string, arg0 int64, arg1 bool, arg2 string, arg3 string) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_bool(arg1)
	cArg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(cArg2))
	cArg3 := C.CString(arg3)
	defer C.free(unsafe.Pointer(cArg3))

	log.Println("  Calling godot_call_void_int_bool_string_string...")
	C.godot_call_void_int_bool_string_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_int_bool_string_string returned")

}
func godotCallVoidNodePathVariant(o Class, methodName string, arg0 *NodePath, arg1 *Variant) {

	cArg0 := arg0.nodePath
	cArg1 := arg1.variant

	log.Println("  Calling godot_call_void_nodepath_variant...")
	C.godot_call_void_nodepath_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_nodepath_variant returned")

}
func godotCallVoidPoolVector3ArrayObjectBool(o Class, methodName string, arg0 *PoolVector3Array, arg1 *Object, arg2 bool) {

	cArg0 := arg0.array
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_void_poolvector3array_object_bool...")
	C.godot_call_void_poolvector3array_object_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_poolvector3array_object_bool returned")

}
func godotCallIntObjectTransformObject(o Class, methodName string, arg0 *Object, arg1 *Transform, arg2 *Object) int64 {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := arg1.transform
	cArg2 := unsafe.Pointer(arg2.owner)

	log.Println("  Calling godot_call_int_object_transform_object...")
	cRet := C.godot_call_int_object_transform_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_int_object_transform_object returned")
	ret := cRet
	return int64(ret)

}
func godotCallVoidRidInt(o Class, methodName string, arg0 *RID, arg1 int64) {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_void_rid_int...")
	C.godot_call_void_rid_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_rid_int returned")

}
func godotCallVoidRidStringRid(o Class, methodName string, arg0 *RID, arg1 string, arg2 *RID) {

	cArg0 := arg0.rid
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := arg2.rid

	log.Println("  Calling godot_call_void_rid_string_rid...")
	C.godot_call_void_rid_string_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rid_string_rid returned")

}
func godotCallVoidIntTransform(o Class, methodName string, arg0 int64, arg1 *Transform) {

	cArg0 := C.godot_int(arg0)
	cArg1 := arg1.transform

	log.Println("  Calling godot_call_void_int_transform...")
	C.godot_call_void_int_transform(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_int_transform returned")

}
func godotCallVoidIntIntFloatBool(o Class, methodName string, arg0 int64, arg1 int64, arg2 float64, arg3 bool) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_real(arg2)
	cArg3 := C.godot_bool(arg3)

	log.Println("  Calling godot_call_void_int_int_float_bool...")
	C.godot_call_void_int_int_float_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_int_int_float_bool returned")

}
func godotCallVector3Rid(o Class, methodName string, arg0 *RID) *Vector3 {

	cArg0 := arg0.rid

	log.Println("  Calling godot_call_vector3_rid...")
	cRet := C.godot_call_vector3_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_vector3_rid returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector3)(cRet)
	return &Vector3{ret}

}
func godotCallTransform2DIntInt(o Class, methodName string, arg0 int64, arg1 int64) *Transform2D {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_transform2d_int_int...")
	cRet := C.godot_call_transform2d_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_transform2d_int_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_transform2d)(cRet)
	return &Transform2D{ret}

}
func godotCallVoidRidBoolRect2(o Class, methodName string, arg0 *RID, arg1 bool, arg2 *Rect2) {

	cArg0 := arg0.rid
	cArg1 := C.godot_bool(arg1)
	cArg2 := arg2.rect2

	log.Println("  Calling godot_call_void_rid_bool_rect2...")
	C.godot_call_void_rid_bool_rect2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rid_bool_rect2 returned")

}
func godotCallVoidVector2FloatColor(o Class, methodName string, arg0 *Vector2, arg1 float64, arg2 *Color) {

	cArg0 := arg0.vector2
	cArg1 := C.godot_real(arg1)
	cArg2 := arg2.color

	log.Println("  Calling godot_call_void_vector2_float_color...")
	C.godot_call_void_vector2_float_color(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_vector2_float_color returned")

}
func godotCallPoolVector3Array(o Class, methodName string) *PoolVector3Array {

	log.Println("  Calling godot_call_poolvector3array...")
	cRet := C.godot_call_poolvector3array(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_poolvector3array returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_pool_vector3_array)(cRet)
	return &PoolVector3Array{ret}

}
func godotCallVoidStringColor(o Class, methodName string, arg0 string, arg1 *Color) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := arg1.color

	log.Println("  Calling godot_call_void_string_color...")
	C.godot_call_void_string_color(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_string_color returned")

}
func godotCallVoidObjectFloat(o Class, methodName string, arg0 *Object, arg1 float64) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.godot_real(arg1)

	log.Println("  Calling godot_call_void_object_float...")
	C.godot_call_void_object_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_object_float returned")

}
func godotCallAabbRidInt(o Class, methodName string, arg0 *RID, arg1 int64) *AABB {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_aabb_rid_int...")
	cRet := C.godot_call_aabb_rid_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_aabb_rid_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_aabb)(cRet)
	return &AABB{ret}

}
func godotCallObjectObject(o Class, methodName string, arg0 *Object) *Object {

	cArg0 := unsafe.Pointer(arg0.owner)

	log.Println("  Calling godot_call_object_object...")
	cRet := C.godot_call_object_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_object_object returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_object)(cRet)
	return &Object{ret}

}
func godotCallStringDictionary(o Class, methodName string, arg0 *Dictionary) string {

	cArg0 := arg0.dictionary

	log.Println("  Calling godot_call_string_dictionary...")
	cRet := C.godot_call_string_dictionary(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_string_dictionary returned")
	return C.GoString(cRet)

}
func godotCallVariantNodePath(o Class, methodName string, arg0 *NodePath) *Variant {

	cArg0 := arg0.nodePath

	log.Println("  Calling godot_call_variant_nodepath...")
	cRet := C.godot_call_variant_nodepath(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_variant_nodepath returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_variant)(cRet)
	return &Variant{ret}

}
func godotCallVoidRidIntVariant(o Class, methodName string, arg0 *RID, arg1 int64, arg2 *Variant) {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)
	cArg2 := arg2.variant

	log.Println("  Calling godot_call_void_rid_int_variant...")
	C.godot_call_void_rid_int_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rid_int_variant returned")

}
func godotCallVoidStringVariantBool(o Class, methodName string, arg0 string, arg1 *Variant, arg2 bool) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := arg1.variant
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_void_string_variant_bool...")
	C.godot_call_void_string_variant_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_string_variant_bool returned")

}
func godotCallVoidIntColor(o Class, methodName string, arg0 int64, arg1 *Color) {

	cArg0 := C.godot_int(arg0)
	cArg1 := arg1.color

	log.Println("  Calling godot_call_void_int_color...")
	C.godot_call_void_int_color(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_int_color returned")

}
func godotCallBoolObjectStringVariantVariantFloatIntIntFloat(o Class, methodName string, arg0 *Object, arg1 string, arg2 *Variant, arg3 *Variant, arg4 float64, arg5 int64, arg6 int64, arg7 float64) bool {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := arg2.variant
	cArg3 := arg3.variant
	cArg4 := C.godot_real(arg4)
	cArg5 := C.godot_int(arg5)
	cArg6 := C.godot_int(arg6)
	cArg7 := C.godot_real(arg7)

	log.Println("  Calling godot_call_bool_object_string_variant_variant_float_int_int_float...")
	cRet := C.godot_call_bool_object_string_variant_variant_float_int_int_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5, cArg6, cArg7,
	)
	log.Println("  ...godot_call_bool_object_string_variant_variant_float_int_int_float returned")
	ret := cRet
	return bool(ret)

}
func godotCallStringInt(o Class, methodName string, arg0 int64) string {

	cArg0 := C.godot_int(arg0)

	log.Println("  Calling godot_call_string_int...")
	cRet := C.godot_call_string_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_string_int returned")
	return C.GoString(cRet)

}
func godotCallVoidStringVariant(o Class, methodName string, arg0 string, arg1 *Variant) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := arg1.variant

	log.Println("  Calling godot_call_void_string_variant...")
	C.godot_call_void_string_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_string_variant returned")

}
func godotCallVoidRid(o Class, methodName string, arg0 *RID) {

	cArg0 := arg0.rid

	log.Println("  Calling godot_call_void_rid...")
	C.godot_call_void_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_rid returned")

}
func godotCallVector2IntInt(o Class, methodName string, arg0 int64, arg1 int64) *Vector2 {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_vector2_int_int...")
	cRet := C.godot_call_vector2_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_vector2_int_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector2)(cRet)
	return &Vector2{ret}

}
func godotCallArrayArrayInt(o Class, methodName string, arg0 *Array, arg1 int64) *Array {

	cArg0 := arg0.array
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_array_array_int...")
	cRet := C.godot_call_array_array_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_array_array_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_array)(cRet)
	return &Array{ret}

}
func godotCallVoidStringObjectBool(o Class, methodName string, arg0 string, arg1 *Object, arg2 bool) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_void_string_object_bool...")
	C.godot_call_void_string_object_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_string_object_bool returned")

}
func godotCallVoidIntObjectString(o Class, methodName string, arg0 int64, arg1 *Object, arg2 string) {

	cArg0 := C.godot_int(arg0)
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(cArg2))

	log.Println("  Calling godot_call_void_int_object_string...")
	C.godot_call_void_int_object_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_int_object_string returned")

}
func godotCallVoidRidPoolIntArrayPoolVector2ArrayPoolColorArrayPoolVector2ArrayRidIntRid(o Class, methodName string, arg0 *RID, arg1 *PoolIntArray, arg2 *PoolVector2Array, arg3 *PoolColorArray, arg4 *PoolVector2Array, arg5 *RID, arg6 int64, arg7 *RID) {

	cArg0 := arg0.rid
	cArg1 := arg1.array
	cArg2 := arg2.array
	cArg3 := arg3.array
	cArg4 := arg4.array
	cArg5 := arg5.rid
	cArg6 := C.godot_int(arg6)
	cArg7 := arg7.rid

	log.Println("  Calling godot_call_void_rid_poolintarray_poolvector2array_poolcolorarray_poolvector2array_rid_int_rid...")
	C.godot_call_void_rid_poolintarray_poolvector2array_poolcolorarray_poolvector2array_rid_int_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5, cArg6, cArg7,
	)
	log.Println("  ...godot_call_void_rid_poolintarray_poolvector2array_poolcolorarray_poolvector2array_rid_int_rid returned")

}
func godotCallVoidString(o Class, methodName string, arg0 string) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))

	log.Println("  Calling godot_call_void_string...")
	C.godot_call_void_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_string returned")

}
func godotCallNodePath(o Class, methodName string) *NodePath {

	log.Println("  Calling godot_call_nodepath...")
	cRet := C.godot_call_nodepath(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_nodepath returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_node_path)(cRet)
	return &NodePath{ret}

}
func godotCallTransform2DRid(o Class, methodName string, arg0 *RID) *Transform2D {

	cArg0 := arg0.rid

	log.Println("  Calling godot_call_transform2d_rid...")
	cRet := C.godot_call_transform2d_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_transform2d_rid returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_transform2d)(cRet)
	return &Transform2D{ret}

}
func godotCallVoidStringStringObject(o Class, methodName string, arg0 string, arg1 string, arg2 *Object) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := unsafe.Pointer(arg2.owner)

	log.Println("  Calling godot_call_void_string_string_object...")
	C.godot_call_void_string_string_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_string_string_object returned")

}
func godotCallNodePathInt(o Class, methodName string, arg0 int64) *NodePath {

	cArg0 := C.godot_int(arg0)

	log.Println("  Calling godot_call_nodepath_int...")
	cRet := C.godot_call_nodepath_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_nodepath_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_node_path)(cRet)
	return &NodePath{ret}

}
func godotCallVoidIntFloatVariantFloat(o Class, methodName string, arg0 int64, arg1 float64, arg2 *Variant, arg3 float64) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_real(arg1)
	cArg2 := arg2.variant
	cArg3 := C.godot_real(arg3)

	log.Println("  Calling godot_call_void_int_float_variant_float...")
	C.godot_call_void_int_float_variant_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_int_float_variant_float returned")

}
func godotCallVoidIntArrayArrayInt(o Class, methodName string, arg0 int64, arg1 *Array, arg2 *Array, arg3 int64) {

	cArg0 := C.godot_int(arg0)
	cArg1 := arg1.array
	cArg2 := arg2.array
	cArg3 := C.godot_int(arg3)

	log.Println("  Calling godot_call_void_int_array_array_int...")
	C.godot_call_void_int_array_array_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_int_array_array_int returned")

}
func godotCallVoidPoolVector2ArrayColorPoolVector2ArrayObjectObjectBool(o Class, methodName string, arg0 *PoolVector2Array, arg1 *Color, arg2 *PoolVector2Array, arg3 *Object, arg4 *Object, arg5 bool) {

	cArg0 := arg0.array
	cArg1 := arg1.color
	cArg2 := arg2.array
	cArg3 := unsafe.Pointer(arg3.owner)
	cArg4 := unsafe.Pointer(arg4.owner)
	cArg5 := C.godot_bool(arg5)

	log.Println("  Calling godot_call_void_poolvector2array_color_poolvector2array_object_object_bool...")
	C.godot_call_void_poolvector2array_color_poolvector2array_object_object_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5,
	)
	log.Println("  ...godot_call_void_poolvector2array_color_poolvector2array_object_object_bool returned")

}
func godotCallVoidStringStringVariant(o Class, methodName string, arg0 string, arg1 string, arg2 *Variant) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := arg2.variant

	log.Println("  Calling godot_call_void_string_string_variant...")
	C.godot_call_void_string_string_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_string_string_variant returned")

}
func godotCallBoolObjectObject(o Class, methodName string, arg0 *Object, arg1 *Object) bool {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := unsafe.Pointer(arg1.owner)

	log.Println("  Calling godot_call_bool_object_object...")
	cRet := C.godot_call_bool_object_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_bool_object_object returned")
	ret := cRet
	return bool(ret)

}
func godotCallIntObjectBoolString(o Class, methodName string, arg0 *Object, arg1 bool, arg2 string) int64 {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.godot_bool(arg1)
	cArg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(cArg2))

	log.Println("  Calling godot_call_int_object_bool_string...")
	cRet := C.godot_call_int_object_bool_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_int_object_bool_string returned")
	ret := cRet
	return int64(ret)

}
func godotCallVoidIntFloatFloatFloatBool(o Class, methodName string, arg0 int64, arg1 float64, arg2 float64, arg3 float64, arg4 bool) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_real(arg1)
	cArg2 := C.godot_real(arg2)
	cArg3 := C.godot_real(arg3)
	cArg4 := C.godot_bool(arg4)

	log.Println("  Calling godot_call_void_int_float_float_float_bool...")
	C.godot_call_void_int_float_float_float_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_void_int_float_float_float_bool returned")

}
func godotCallVoidIntObjectInt(o Class, methodName string, arg0 int64, arg1 *Object, arg2 int64) {

	cArg0 := C.godot_int(arg0)
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_void_int_object_int...")
	C.godot_call_void_int_object_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_int_object_int returned")

}
func godotCallColor(o Class, methodName string) *Color {

	log.Println("  Calling godot_call_color...")
	cRet := C.godot_call_color(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_color returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_color)(cRet)
	return &Color{ret}

}
func godotCallObjectRect2(o Class, methodName string, arg0 *Rect2) *Object {

	cArg0 := arg0.rect2

	log.Println("  Calling godot_call_object_rect2...")
	cRet := C.godot_call_object_rect2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_object_rect2 returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_object)(cRet)
	return &Object{ret}

}
func godotCallVoidObjectStringArray(o Class, methodName string, arg0 *Object, arg1 string, arg2 *Array) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := arg2.array

	log.Println("  Calling godot_call_void_object_string_array...")
	C.godot_call_void_object_string_array(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_object_string_array returned")

}
func godotCallVoidInt(o Class, methodName string, arg0 int64) {

	cArg0 := C.godot_int(arg0)

	log.Println("  Calling godot_call_void_int...")
	C.godot_call_void_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_int returned")

}
func godotCallColorStringString(o Class, methodName string, arg0 string, arg1 string) *Color {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))

	log.Println("  Calling godot_call_color_string_string...")
	cRet := C.godot_call_color_string_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_color_string_string returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_color)(cRet)
	return &Color{ret}

}
func godotCallVoidIntIntIntIntInt(o Class, methodName string, arg0 int64, arg1 int64, arg2 int64, arg3 int64, arg4 int64) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_int(arg2)
	cArg3 := C.godot_int(arg3)
	cArg4 := C.godot_int(arg4)

	log.Println("  Calling godot_call_void_int_int_int_int_int...")
	C.godot_call_void_int_int_int_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_void_int_int_int_int_int returned")

}
func godotCallVoidObjectIntTransform(o Class, methodName string, arg0 *Object, arg1 int64, arg2 *Transform) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.godot_int(arg1)
	cArg2 := arg2.transform

	log.Println("  Calling godot_call_void_object_int_transform...")
	C.godot_call_void_object_int_transform(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_object_int_transform returned")

}
func godotCallVoidIntObjectTransform2DBoolVector2(o Class, methodName string, arg0 int64, arg1 *Object, arg2 *Transform2D, arg3 bool, arg4 *Vector2) {

	cArg0 := C.godot_int(arg0)
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := arg2.transform2d
	cArg3 := C.godot_bool(arg3)
	cArg4 := arg4.vector2

	log.Println("  Calling godot_call_void_int_object_transform2d_bool_vector2...")
	C.godot_call_void_int_object_transform2d_bool_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_void_int_object_transform2d_bool_vector2 returned")

}
func godotCallBoolRidTransform2DVector2FloatObject(o Class, methodName string, arg0 *RID, arg1 *Transform2D, arg2 *Vector2, arg3 float64, arg4 *Object) bool {

	cArg0 := arg0.rid
	cArg1 := arg1.transform2d
	cArg2 := arg2.vector2
	cArg3 := C.godot_real(arg3)
	cArg4 := unsafe.Pointer(arg4.owner)

	log.Println("  Calling godot_call_bool_rid_transform2d_vector2_float_object...")
	cRet := C.godot_call_bool_rid_transform2d_vector2_float_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_bool_rid_transform2d_vector2_float_object returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoidIntObjectBool(o Class, methodName string, arg0 int64, arg1 *Object, arg2 bool) {

	cArg0 := C.godot_int(arg0)
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_void_int_object_bool...")
	C.godot_call_void_int_object_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_int_object_bool returned")

}
func godotCallRidObjectInt(o Class, methodName string, arg0 *Object, arg1 int64) *RID {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_rid_object_int...")
	cRet := C.godot_call_rid_object_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_rid_object_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_rid)(cRet)
	return &RID{ret}

}
func godotCallVector2Int(o Class, methodName string, arg0 int64) *Vector2 {

	cArg0 := C.godot_int(arg0)

	log.Println("  Calling godot_call_vector2_int...")
	cRet := C.godot_call_vector2_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_vector2_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector2)(cRet)
	return &Vector2{ret}

}
func godotCallIntObject(o Class, methodName string, arg0 *Object) int64 {

	cArg0 := unsafe.Pointer(arg0.owner)

	log.Println("  Calling godot_call_int_object...")
	cRet := C.godot_call_int_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_int_object returned")
	ret := cRet
	return int64(ret)

}
func godotCallObjectObjectString(o Class, methodName string, arg0 *Object, arg1 string) *Object {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))

	log.Println("  Calling godot_call_object_object_string...")
	cRet := C.godot_call_object_object_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_object_object_string returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_object)(cRet)
	return &Object{ret}

}
func godotCallBoolStringVariant(o Class, methodName string, arg0 string, arg1 *Variant) bool {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := arg1.variant

	log.Println("  Calling godot_call_bool_string_variant...")
	cRet := C.godot_call_bool_string_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_bool_string_variant returned")
	ret := cRet
	return bool(ret)

}
func godotCallVariantTransform2DVector2ObjectTransform2DVector2(o Class, methodName string, arg0 *Transform2D, arg1 *Vector2, arg2 *Object, arg3 *Transform2D, arg4 *Vector2) *Variant {

	cArg0 := arg0.transform2d
	cArg1 := arg1.vector2
	cArg2 := unsafe.Pointer(arg2.owner)
	cArg3 := arg3.transform2d
	cArg4 := arg4.vector2

	log.Println("  Calling godot_call_variant_transform2d_vector2_object_transform2d_vector2...")
	cRet := C.godot_call_variant_transform2d_vector2_object_transform2d_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_variant_transform2d_vector2_object_transform2d_vector2 returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_variant)(cRet)
	return &Variant{ret}

}
func godotCallFloatIntInt(o Class, methodName string, arg0 int64, arg1 int64) float64 {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_float_int_int...")
	cRet := C.godot_call_float_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_float_int_int returned")
	ret := cRet
	return float64(ret)

}
func godotCallVoidObjectObjectStringVariant(o Class, methodName string, arg0 *Object, arg1 *Object, arg2 string, arg3 *Variant) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(cArg2))
	cArg3 := arg3.variant

	log.Println("  Calling godot_call_void_object_object_string_variant...")
	C.godot_call_void_object_object_string_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_object_object_string_variant returned")

}
func godotCallVariantRid(o Class, methodName string, arg0 *RID) *Variant {

	cArg0 := arg0.rid

	log.Println("  Calling godot_call_variant_rid...")
	cRet := C.godot_call_variant_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_variant_rid returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_variant)(cRet)
	return &Variant{ret}

}
func godotCallVoidIntIntIntBoolBoolBoolVector2(o Class, methodName string, arg0 int64, arg1 int64, arg2 int64, arg3 bool, arg4 bool, arg5 bool, arg6 *Vector2) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_int(arg2)
	cArg3 := C.godot_bool(arg3)
	cArg4 := C.godot_bool(arg4)
	cArg5 := C.godot_bool(arg5)
	cArg6 := arg6.vector2

	log.Println("  Calling godot_call_void_int_int_int_bool_bool_bool_vector2...")
	C.godot_call_void_int_int_int_bool_bool_bool_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5, cArg6,
	)
	log.Println("  ...godot_call_void_int_int_int_bool_bool_bool_vector2 returned")

}
func godotCallVoidIntIntFloat(o Class, methodName string, arg0 int64, arg1 int64, arg2 float64) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_real(arg2)

	log.Println("  Calling godot_call_void_int_int_float...")
	C.godot_call_void_int_int_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_int_int_float returned")

}
func godotCallVoidStringObjectVariant(o Class, methodName string, arg0 string, arg1 *Object, arg2 *Variant) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := arg2.variant

	log.Println("  Calling godot_call_void_string_object_variant...")
	C.godot_call_void_string_object_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_string_object_variant returned")

}
func godotCallRect2ObjectInt(o Class, methodName string, arg0 *Object, arg1 int64) *Rect2 {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_rect2_object_int...")
	cRet := C.godot_call_rect2_object_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_rect2_object_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_rect2)(cRet)
	return &Rect2{ret}

}
func godotCallVariantRidString(o Class, methodName string, arg0 *RID, arg1 string) *Variant {

	cArg0 := arg0.rid
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))

	log.Println("  Calling godot_call_variant_rid_string...")
	cRet := C.godot_call_variant_rid_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_variant_rid_string returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_variant)(cRet)
	return &Variant{ret}

}
func godotCallVoidIntInt(o Class, methodName string, arg0 int64, arg1 int64) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_void_int_int...")
	C.godot_call_void_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_int_int returned")

}
func godotCallVoidStringInt(o Class, methodName string, arg0 string, arg1 int64) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_void_string_int...")
	C.godot_call_void_string_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_string_int returned")

}
func godotCallVoidObjectRect2Rect2ColorBoolObjectBool(o Class, methodName string, arg0 *Object, arg1 *Rect2, arg2 *Rect2, arg3 *Color, arg4 bool, arg5 *Object, arg6 bool) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := arg1.rect2
	cArg2 := arg2.rect2
	cArg3 := arg3.color
	cArg4 := C.godot_bool(arg4)
	cArg5 := unsafe.Pointer(arg5.owner)
	cArg6 := C.godot_bool(arg6)

	log.Println("  Calling godot_call_void_object_rect2_rect2_color_bool_object_bool...")
	C.godot_call_void_object_rect2_rect2_color_bool_object_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5, cArg6,
	)
	log.Println("  ...godot_call_void_object_rect2_rect2_color_bool_object_bool returned")

}
func godotCallVariantInt(o Class, methodName string, arg0 int64) *Variant {

	cArg0 := C.godot_int(arg0)

	log.Println("  Calling godot_call_variant_int...")
	cRet := C.godot_call_variant_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_variant_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_variant)(cRet)
	return &Variant{ret}

}
func godotCallVoidFloatColor(o Class, methodName string, arg0 float64, arg1 *Color) {

	cArg0 := C.godot_real(arg0)
	cArg1 := arg1.color

	log.Println("  Calling godot_call_void_float_color...")
	C.godot_call_void_float_color(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_float_color returned")

}
func godotCallRidVector2RidRid(o Class, methodName string, arg0 *Vector2, arg1 *RID, arg2 *RID) *RID {

	cArg0 := arg0.vector2
	cArg1 := arg1.rid
	cArg2 := arg2.rid

	log.Println("  Calling godot_call_rid_vector2_rid_rid...")
	cRet := C.godot_call_rid_vector2_rid_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_rid_vector2_rid_rid returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_rid)(cRet)
	return &RID{ret}

}
func godotCallVoidObjectStringVariant(o Class, methodName string, arg0 *Object, arg1 string, arg2 *Variant) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := arg2.variant

	log.Println("  Calling godot_call_void_object_string_variant...")
	C.godot_call_void_object_string_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_object_string_variant returned")

}
func godotCallVoidStringStringInt(o Class, methodName string, arg0 string, arg1 string, arg2 int64) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_void_string_string_int...")
	C.godot_call_void_string_string_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_string_string_int returned")

}
func godotCallVector2(o Class, methodName string) *Vector2 {

	log.Println("  Calling godot_call_vector2...")
	cRet := C.godot_call_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_vector2 returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector2)(cRet)
	return &Vector2{ret}

}
func godotCallVoidStringStringFloat(o Class, methodName string, arg0 string, arg1 string, arg2 float64) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := C.godot_real(arg2)

	log.Println("  Calling godot_call_void_string_string_float...")
	C.godot_call_void_string_string_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_string_string_float returned")

}
func godotCallVoidObjectInt(o Class, methodName string, arg0 *Object, arg1 int64) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_void_object_int...")
	C.godot_call_void_object_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_object_int returned")

}
func godotCallVariantArray(o Class, methodName string, arg0 *Array) *Variant {

	cArg0 := arg0.array

	log.Println("  Calling godot_call_variant_array...")
	cRet := C.godot_call_variant_array(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_variant_array returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_variant)(cRet)
	return &Variant{ret}

}
func godotCallVoidStringIntObject(o Class, methodName string, arg0 string, arg1 int64, arg2 *Object) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)
	cArg2 := unsafe.Pointer(arg2.owner)

	log.Println("  Calling godot_call_void_string_int_object...")
	C.godot_call_void_string_int_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_string_int_object returned")

}
func godotCallVoidVector2IntBoolBoolBool(o Class, methodName string, arg0 *Vector2, arg1 int64, arg2 bool, arg3 bool, arg4 bool) {

	cArg0 := arg0.vector2
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_bool(arg2)
	cArg3 := C.godot_bool(arg3)
	cArg4 := C.godot_bool(arg4)

	log.Println("  Calling godot_call_void_vector2_int_bool_bool_bool...")
	C.godot_call_void_vector2_int_bool_bool_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_void_vector2_int_bool_bool_bool returned")

}
func godotCallBoolObjectNodePathVariantObjectNodePathFloatIntIntFloat(o Class, methodName string, arg0 *Object, arg1 *NodePath, arg2 *Variant, arg3 *Object, arg4 *NodePath, arg5 float64, arg6 int64, arg7 int64, arg8 float64) bool {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := arg1.nodePath
	cArg2 := arg2.variant
	cArg3 := unsafe.Pointer(arg3.owner)
	cArg4 := arg4.nodePath
	cArg5 := C.godot_real(arg5)
	cArg6 := C.godot_int(arg6)
	cArg7 := C.godot_int(arg7)
	cArg8 := C.godot_real(arg8)

	log.Println("  Calling godot_call_bool_object_nodepath_variant_object_nodepath_float_int_int_float...")
	cRet := C.godot_call_bool_object_nodepath_variant_object_nodepath_float_int_int_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5, cArg6, cArg7, cArg8,
	)
	log.Println("  ...godot_call_bool_object_nodepath_variant_object_nodepath_float_int_int_float returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoidObjectStringInt(o Class, methodName string, arg0 *Object, arg1 string, arg2 int64) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_void_object_string_int...")
	C.godot_call_void_object_string_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_object_string_int returned")

}
func godotCallVoidRidVector3Vector3(o Class, methodName string, arg0 *RID, arg1 *Vector3, arg2 *Vector3) {

	cArg0 := arg0.rid
	cArg1 := arg1.vector3
	cArg2 := arg2.vector3

	log.Println("  Calling godot_call_void_rid_vector3_vector3...")
	C.godot_call_void_rid_vector3_vector3(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rid_vector3_vector3 returned")

}
func godotCallVoidFloatFloatFloatFloat(o Class, methodName string, arg0 float64, arg1 float64, arg2 float64, arg3 float64) {

	cArg0 := C.godot_real(arg0)
	cArg1 := C.godot_real(arg1)
	cArg2 := C.godot_real(arg2)
	cArg3 := C.godot_real(arg3)

	log.Println("  Calling godot_call_void_float_float_float_float...")
	C.godot_call_void_float_float_float_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_float_float_float_float returned")

}
func godotCallRid(o Class, methodName string) *RID {

	log.Println("  Calling godot_call_rid...")
	cRet := C.godot_call_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_rid returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_rid)(cRet)
	return &RID{ret}

}
func godotCallVector2IntFloat(o Class, methodName string, arg0 int64, arg1 float64) *Vector2 {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_real(arg1)

	log.Println("  Calling godot_call_vector2_int_float...")
	cRet := C.godot_call_vector2_int_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_vector2_int_float returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector2)(cRet)
	return &Vector2{ret}

}
func godotCallObjectTransform2DVector2(o Class, methodName string, arg0 *Transform2D, arg1 *Vector2) *Object {

	cArg0 := arg0.transform2d
	cArg1 := arg1.vector2

	log.Println("  Calling godot_call_object_transform2d_vector2...")
	cRet := C.godot_call_object_transform2d_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_object_transform2d_vector2 returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_object)(cRet)
	return &Object{ret}

}
func godotCallPoolIntArrayIntInt(o Class, methodName string, arg0 int64, arg1 int64) *PoolIntArray {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_poolintarray_int_int...")
	cRet := C.godot_call_poolintarray_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_poolintarray_int_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_pool_int_array)(cRet)
	return &PoolIntArray{ret}

}
func godotCallVoidPoolStringArray(o Class, methodName string, arg0 *PoolStringArray) {

	cArg0 := arg0.array

	log.Println("  Calling godot_call_void_poolstringarray...")
	C.godot_call_void_poolstringarray(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_poolstringarray returned")

}
func godotCallVoidRidVector2ColorBoolObject(o Class, methodName string, arg0 *RID, arg1 *Vector2, arg2 *Color, arg3 bool, arg4 *Object) {

	cArg0 := arg0.rid
	cArg1 := arg1.vector2
	cArg2 := arg2.color
	cArg3 := C.godot_bool(arg3)
	cArg4 := unsafe.Pointer(arg4.owner)

	log.Println("  Calling godot_call_void_rid_vector2_color_bool_object...")
	C.godot_call_void_rid_vector2_color_bool_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_void_rid_vector2_color_bool_object returned")

}
func godotCallVoidIntFloat(o Class, methodName string, arg0 int64, arg1 float64) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_real(arg1)

	log.Println("  Calling godot_call_void_int_float...")
	C.godot_call_void_int_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_int_float returned")

}
func godotCallVector3Vector2(o Class, methodName string, arg0 *Vector2) *Vector3 {

	cArg0 := arg0.vector2

	log.Println("  Calling godot_call_vector3_vector2...")
	cRet := C.godot_call_vector3_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_vector3_vector2 returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector3)(cRet)
	return &Vector3{ret}

}
func godotCallVoidPoolVector3ArrayPoolVector2ArrayPoolColorArrayPoolVector2ArrayPoolVector3ArrayArray(o Class, methodName string, arg0 *PoolVector3Array, arg1 *PoolVector2Array, arg2 *PoolColorArray, arg3 *PoolVector2Array, arg4 *PoolVector3Array, arg5 *Array) {

	cArg0 := arg0.array
	cArg1 := arg1.array
	cArg2 := arg2.array
	cArg3 := arg3.array
	cArg4 := arg4.array
	cArg5 := arg5.array

	log.Println("  Calling godot_call_void_poolvector3array_poolvector2array_poolcolorarray_poolvector2array_poolvector3array_array...")
	C.godot_call_void_poolvector3array_poolvector2array_poolcolorarray_poolvector2array_poolvector3array_array(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5,
	)
	log.Println("  ...godot_call_void_poolvector3array_poolvector2array_poolcolorarray_poolvector2array_poolvector3array_array returned")

}
func godotCallPoolIntArrayStringIntIntInt(o Class, methodName string, arg0 string, arg1 int64, arg2 int64, arg3 int64) *PoolIntArray {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_int(arg2)
	cArg3 := C.godot_int(arg3)

	log.Println("  Calling godot_call_poolintarray_string_int_int_int...")
	cRet := C.godot_call_poolintarray_string_int_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_poolintarray_string_int_int_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_pool_int_array)(cRet)
	return &PoolIntArray{ret}

}
func godotCallVoidIntIntObject(o Class, methodName string, arg0 int64, arg1 int64, arg2 *Object) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := unsafe.Pointer(arg2.owner)

	log.Println("  Calling godot_call_void_int_int_object...")
	C.godot_call_void_int_int_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_int_int_object returned")

}
func godotCallVoidRidRidInt(o Class, methodName string, arg0 *RID, arg1 *RID, arg2 int64) {

	cArg0 := arg0.rid
	cArg1 := arg1.rid
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_void_rid_rid_int...")
	C.godot_call_void_rid_rid_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rid_rid_int returned")

}
func godotCallVoidBasis(o Class, methodName string, arg0 *Basis) {

	cArg0 := arg0.basis

	log.Println("  Calling godot_call_void_basis...")
	C.godot_call_void_basis(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_basis returned")

}
func godotCallFloatVector2(o Class, methodName string, arg0 *Vector2) float64 {

	cArg0 := arg0.vector2

	log.Println("  Calling godot_call_float_vector2...")
	cRet := C.godot_call_float_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_float_vector2 returned")
	ret := cRet
	return float64(ret)

}
func godotCallVoidStringStringColor(o Class, methodName string, arg0 string, arg1 string, arg2 *Color) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := arg2.color

	log.Println("  Calling godot_call_void_string_string_color...")
	C.godot_call_void_string_string_color(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_string_string_color returned")

}
func godotCallVoidIntString(o Class, methodName string, arg0 int64, arg1 string) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))

	log.Println("  Calling godot_call_void_int_string...")
	C.godot_call_void_int_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_int_string returned")

}
func godotCallVoidStringObject(o Class, methodName string, arg0 string, arg1 *Object) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := unsafe.Pointer(arg1.owner)

	log.Println("  Calling godot_call_void_string_object...")
	C.godot_call_void_string_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_string_object returned")

}
func godotCallVoidStringObjectIntStringVariant(o Class, methodName string, arg0 string, arg1 *Object, arg2 int64, arg3 string, arg4 *Variant) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := C.godot_int(arg2)
	cArg3 := C.CString(arg3)
	defer C.free(unsafe.Pointer(cArg3))
	cArg4 := arg4.variant

	log.Println("  Calling godot_call_void_string_object_int_string_variant...")
	C.godot_call_void_string_object_int_string_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_void_string_object_int_string_variant returned")

}
func godotCallVector2Vector2Vector2FloatIntFloat(o Class, methodName string, arg0 *Vector2, arg1 *Vector2, arg2 float64, arg3 int64, arg4 float64) *Vector2 {

	cArg0 := arg0.vector2
	cArg1 := arg1.vector2
	cArg2 := C.godot_real(arg2)
	cArg3 := C.godot_int(arg3)
	cArg4 := C.godot_real(arg4)

	log.Println("  Calling godot_call_vector2_vector2_vector2_float_int_float...")
	cRet := C.godot_call_vector2_vector2_vector2_float_int_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_vector2_vector2_vector2_float_int_float returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector2)(cRet)
	return &Vector2{ret}

}
func godotCallVoidObject(o Class, methodName string, arg0 *Object) {

	cArg0 := unsafe.Pointer(arg0.owner)

	log.Println("  Calling godot_call_void_object...")
	C.godot_call_void_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_object returned")

}
func godotCallBoolIntInt(o Class, methodName string, arg0 int64, arg1 int64) bool {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_bool_int_int...")
	cRet := C.godot_call_bool_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_bool_int_int returned")
	ret := cRet
	return bool(ret)

}
func godotCallIntIntFloatVector3QuatVector3(o Class, methodName string, arg0 int64, arg1 float64, arg2 *Vector3, arg3 *Quat, arg4 *Vector3) int64 {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_real(arg1)
	cArg2 := arg2.vector3
	cArg3 := arg3.quat
	cArg4 := arg4.vector3

	log.Println("  Calling godot_call_int_int_float_vector3_quat_vector3...")
	cRet := C.godot_call_int_int_float_vector3_quat_vector3(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_int_int_float_vector3_quat_vector3 returned")
	ret := cRet
	return int64(ret)

}
func godotCallBoolTransform2DVector2(o Class, methodName string, arg0 *Transform2D, arg1 *Vector2) bool {

	cArg0 := arg0.transform2d
	cArg1 := arg1.vector2

	log.Println("  Calling godot_call_bool_transform2d_vector2...")
	cRet := C.godot_call_bool_transform2d_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_bool_transform2d_vector2 returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoidStringIntIntInt(o Class, methodName string, arg0 string, arg1 int64, arg2 int64, arg3 int64) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_int(arg2)
	cArg3 := C.godot_int(arg3)

	log.Println("  Calling godot_call_void_string_int_int_int...")
	C.godot_call_void_string_int_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_string_int_int_int returned")

}
func godotCallPoolByteArray(o Class, methodName string) *PoolByteArray {

	log.Println("  Calling godot_call_poolbytearray...")
	cRet := C.godot_call_poolbytearray(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_poolbytearray returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_pool_byte_array)(cRet)
	return &PoolByteArray{ret}

}
func godotCallVoidVector2Bool(o Class, methodName string, arg0 *Vector2, arg1 bool) {

	cArg0 := arg0.vector2
	cArg1 := C.godot_bool(arg1)

	log.Println("  Calling godot_call_void_vector2_bool...")
	C.godot_call_void_vector2_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_vector2_bool returned")

}
func godotCallVoidFloatFloatFloat(o Class, methodName string, arg0 float64, arg1 float64, arg2 float64) {

	cArg0 := C.godot_real(arg0)
	cArg1 := C.godot_real(arg1)
	cArg2 := C.godot_real(arg2)

	log.Println("  Calling godot_call_void_float_float_float...")
	C.godot_call_void_float_float_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_float_float_float returned")

}
func godotCallPlaneInt(o Class, methodName string, arg0 int64) *Plane {

	cArg0 := C.godot_int(arg0)

	log.Println("  Calling godot_call_plane_int...")
	cRet := C.godot_call_plane_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_plane_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_plane)(cRet)
	return &Plane{ret}

}
func godotCallTransformRid(o Class, methodName string, arg0 *RID) *Transform {

	cArg0 := arg0.rid

	log.Println("  Calling godot_call_transform_rid...")
	cRet := C.godot_call_transform_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_transform_rid returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_transform)(cRet)
	return &Transform{ret}

}
func godotCallBoolStringStringInt(o Class, methodName string, arg0 string, arg1 string, arg2 int64) bool {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_bool_string_string_int...")
	cRet := C.godot_call_bool_string_string_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_bool_string_string_int returned")
	ret := cRet
	return bool(ret)

}
func godotCallArrayRidInt(o Class, methodName string, arg0 *RID, arg1 int64) *Array {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_array_rid_int...")
	cRet := C.godot_call_array_rid_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_array_rid_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_array)(cRet)
	return &Array{ret}

}
func godotCallString(o Class, methodName string) string {

	log.Println("  Calling godot_call_string...")
	cRet := C.godot_call_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_string returned")
	return C.GoString(cRet)

}
func godotCallVoidPoolColorArray(o Class, methodName string, arg0 *PoolColorArray) {

	cArg0 := arg0.array

	log.Println("  Calling godot_call_void_poolcolorarray...")
	C.godot_call_void_poolcolorarray(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_poolcolorarray returned")

}
func godotCallVoidPoolVector2ArrayInt(o Class, methodName string, arg0 *PoolVector2Array, arg1 int64) {

	cArg0 := arg0.array
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_void_poolvector2array_int...")
	C.godot_call_void_poolvector2array_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_poolvector2array_int returned")

}
func godotCallPoolStringArrayInt(o Class, methodName string, arg0 int64) *PoolStringArray {

	cArg0 := C.godot_int(arg0)

	log.Println("  Calling godot_call_poolstringarray_int...")
	cRet := C.godot_call_poolstringarray_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_poolstringarray_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_pool_string_array)(cRet)
	return &PoolStringArray{ret}

}
func godotCallVoidRidRect2Rect2ColorBoolObjectBool(o Class, methodName string, arg0 *RID, arg1 *Rect2, arg2 *Rect2, arg3 *Color, arg4 bool, arg5 *Object, arg6 bool) {

	cArg0 := arg0.rid
	cArg1 := arg1.rect2
	cArg2 := arg2.rect2
	cArg3 := arg3.color
	cArg4 := C.godot_bool(arg4)
	cArg5 := unsafe.Pointer(arg5.owner)
	cArg6 := C.godot_bool(arg6)

	log.Println("  Calling godot_call_void_rid_rect2_rect2_color_bool_object_bool...")
	C.godot_call_void_rid_rect2_rect2_color_bool_object_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5, cArg6,
	)
	log.Println("  ...godot_call_void_rid_rect2_rect2_color_bool_object_bool returned")

}
func godotCallVoidRidRidRid(o Class, methodName string, arg0 *RID, arg1 *RID, arg2 *RID) {

	cArg0 := arg0.rid
	cArg1 := arg1.rid
	cArg2 := arg2.rid

	log.Println("  Calling godot_call_void_rid_rid_rid...")
	C.godot_call_void_rid_rid_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rid_rid_rid returned")

}
func godotCallVector3(o Class, methodName string) *Vector3 {

	log.Println("  Calling godot_call_vector3...")
	cRet := C.godot_call_vector3(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_vector3 returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector3)(cRet)
	return &Vector3{ret}

}
func godotCallVoidBool(o Class, methodName string, arg0 bool) {

	cArg0 := C.godot_bool(arg0)

	log.Println("  Calling godot_call_void_bool...")
	C.godot_call_void_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_bool returned")

}
func godotCallVoidIntIntBool(o Class, methodName string, arg0 int64, arg1 int64, arg2 bool) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_void_int_int_bool...")
	C.godot_call_void_int_int_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_int_int_bool returned")

}
func godotCallIntObjectBool(o Class, methodName string, arg0 *Object, arg1 bool) int64 {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.godot_bool(arg1)

	log.Println("  Calling godot_call_int_object_bool...")
	cRet := C.godot_call_int_object_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_int_object_bool returned")
	ret := cRet
	return int64(ret)

}
func godotCallVoidIntIntIntInt(o Class, methodName string, arg0 int64, arg1 int64, arg2 int64, arg3 int64) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_int(arg2)
	cArg3 := C.godot_int(arg3)

	log.Println("  Calling godot_call_void_int_int_int_int...")
	C.godot_call_void_int_int_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_int_int_int_int returned")

}
func godotCallObjectFloat(o Class, methodName string, arg0 float64) *Object {

	cArg0 := C.godot_real(arg0)

	log.Println("  Calling godot_call_object_float...")
	cRet := C.godot_call_object_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_object_float returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_object)(cRet)
	return &Object{ret}

}
func godotCallNodePathObject(o Class, methodName string, arg0 *Object) *NodePath {

	cArg0 := unsafe.Pointer(arg0.owner)

	log.Println("  Calling godot_call_nodepath_object...")
	cRet := C.godot_call_nodepath_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_nodepath_object returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_node_path)(cRet)
	return &NodePath{ret}

}
func godotCallIntIntStringInt(o Class, methodName string, arg0 int64, arg1 string, arg2 int64) int64 {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_int_int_string_int...")
	cRet := C.godot_call_int_int_string_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_int_int_string_int returned")
	ret := cRet
	return int64(ret)

}
func godotCallVariantTransform2DObjectTransform2D(o Class, methodName string, arg0 *Transform2D, arg1 *Object, arg2 *Transform2D) *Variant {

	cArg0 := arg0.transform2d
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := arg2.transform2d

	log.Println("  Calling godot_call_variant_transform2d_object_transform2d...")
	cRet := C.godot_call_variant_transform2d_object_transform2d(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_variant_transform2d_object_transform2d returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_variant)(cRet)
	return &Variant{ret}

}
func godotCallBoolVector2Variant(o Class, methodName string, arg0 *Vector2, arg1 *Variant) bool {

	cArg0 := arg0.vector2
	cArg1 := arg1.variant

	log.Println("  Calling godot_call_bool_vector2_variant...")
	cRet := C.godot_call_bool_vector2_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_bool_vector2_variant returned")
	ret := cRet
	return bool(ret)

}
func godotCallObjectVector2(o Class, methodName string, arg0 *Vector2) *Object {

	cArg0 := arg0.vector2

	log.Println("  Calling godot_call_object_vector2...")
	cRet := C.godot_call_object_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_object_vector2 returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_object)(cRet)
	return &Object{ret}

}
func godotCallPoolVector2ArrayVector2Vector2(o Class, methodName string, arg0 *Vector2, arg1 *Vector2) *PoolVector2Array {

	cArg0 := arg0.vector2
	cArg1 := arg1.vector2

	log.Println("  Calling godot_call_poolvector2array_vector2_vector2...")
	cRet := C.godot_call_poolvector2array_vector2_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_poolvector2array_vector2_vector2 returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_pool_vector2_array)(cRet)
	return &PoolVector2Array{ret}

}
func godotCallVector2Vector2Bool(o Class, methodName string, arg0 *Vector2, arg1 bool) *Vector2 {

	cArg0 := arg0.vector2
	cArg1 := C.godot_bool(arg1)

	log.Println("  Calling godot_call_vector2_vector2_bool...")
	cRet := C.godot_call_vector2_vector2_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_vector2_vector2_bool returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector2)(cRet)
	return &Vector2{ret}

}
func godotCallVariantObjectString(o Class, methodName string, arg0 *Object, arg1 string) *Variant {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))

	log.Println("  Calling godot_call_variant_object_string...")
	cRet := C.godot_call_variant_object_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_variant_object_string returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_variant)(cRet)
	return &Variant{ret}

}
func godotCallBoolStringIntStringInt(o Class, methodName string, arg0 string, arg1 int64, arg2 string, arg3 int64) bool {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)
	cArg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(cArg2))
	cArg3 := C.godot_int(arg3)

	log.Println("  Calling godot_call_bool_string_int_string_int...")
	cRet := C.godot_call_bool_string_int_string_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_bool_string_int_string_int returned")
	ret := cRet
	return bool(ret)

}
func godotCallBoolFloat(o Class, methodName string, arg0 float64) bool {

	cArg0 := C.godot_real(arg0)

	log.Println("  Calling godot_call_bool_float...")
	cRet := C.godot_call_bool_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_bool_float returned")
	ret := cRet
	return bool(ret)

}
func godotCallObjectNodePath(o Class, methodName string, arg0 *NodePath) *Object {

	cArg0 := arg0.nodePath

	log.Println("  Calling godot_call_object_nodepath...")
	cRet := C.godot_call_object_nodepath(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_object_nodepath returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_object)(cRet)
	return &Object{ret}

}
func godotCallIntStringObjectStringArrayInt(o Class, methodName string, arg0 string, arg1 *Object, arg2 string, arg3 *Array, arg4 int64) int64 {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(cArg2))
	cArg3 := arg3.array
	cArg4 := C.godot_int(arg4)

	log.Println("  Calling godot_call_int_string_object_string_array_int...")
	cRet := C.godot_call_int_string_object_string_array_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_int_string_object_string_array_int returned")
	ret := cRet
	return int64(ret)

}
func godotCallIntRid(o Class, methodName string, arg0 *RID) int64 {

	cArg0 := arg0.rid

	log.Println("  Calling godot_call_int_rid...")
	cRet := C.godot_call_int_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_int_rid returned")
	ret := cRet
	return int64(ret)

}
func godotCallBoolObjectStringObjectStringVariantFloatIntIntFloat(o Class, methodName string, arg0 *Object, arg1 string, arg2 *Object, arg3 string, arg4 *Variant, arg5 float64, arg6 int64, arg7 int64, arg8 float64) bool {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := unsafe.Pointer(arg2.owner)
	cArg3 := C.CString(arg3)
	defer C.free(unsafe.Pointer(cArg3))
	cArg4 := arg4.variant
	cArg5 := C.godot_real(arg5)
	cArg6 := C.godot_int(arg6)
	cArg7 := C.godot_int(arg7)
	cArg8 := C.godot_real(arg8)

	log.Println("  Calling godot_call_bool_object_string_object_string_variant_float_int_int_float...")
	cRet := C.godot_call_bool_object_string_object_string_variant_float_int_int_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5, cArg6, cArg7, cArg8,
	)
	log.Println("  ...godot_call_bool_object_string_object_string_variant_float_int_int_float returned")
	ret := cRet
	return bool(ret)

}
func godotCallVector2Float(o Class, methodName string, arg0 float64) *Vector2 {

	cArg0 := C.godot_real(arg0)

	log.Println("  Calling godot_call_vector2_float...")
	cRet := C.godot_call_vector2_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_vector2_float returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector2)(cRet)
	return &Vector2{ret}

}
func godotCallVoidStringStringPoolStringArray(o Class, methodName string, arg0 string, arg1 string, arg2 *PoolStringArray) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := arg2.array

	log.Println("  Calling godot_call_void_string_string_poolstringarray...")
	C.godot_call_void_string_string_poolstringarray(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_string_string_poolstringarray returned")

}
func godotCallIntObjectInt(o Class, methodName string, arg0 *Object, arg1 int64) int64 {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_int_object_int...")
	cRet := C.godot_call_int_object_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_int_object_int returned")
	ret := cRet
	return int64(ret)

}
func godotCallVoidVector2(o Class, methodName string, arg0 *Vector2) {

	cArg0 := arg0.vector2

	log.Println("  Calling godot_call_void_vector2...")
	C.godot_call_void_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_vector2 returned")

}
func godotCallVoidStringFloat(o Class, methodName string, arg0 string, arg1 float64) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_real(arg1)

	log.Println("  Calling godot_call_void_string_float...")
	C.godot_call_void_string_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_string_float returned")

}
func godotCallVoidIntIntPoolByteArray(o Class, methodName string, arg0 int64, arg1 int64, arg2 *PoolByteArray) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := arg2.array

	log.Println("  Calling godot_call_void_int_int_poolbytearray...")
	C.godot_call_void_int_int_poolbytearray(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_int_int_poolbytearray returned")

}
func godotCallFloatObjectVector2StringStringColor(o Class, methodName string, arg0 *Object, arg1 *Vector2, arg2 string, arg3 string, arg4 *Color) float64 {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := arg1.vector2
	cArg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(cArg2))
	cArg3 := C.CString(arg3)
	defer C.free(unsafe.Pointer(cArg3))
	cArg4 := arg4.color

	log.Println("  Calling godot_call_float_object_vector2_string_string_color...")
	cRet := C.godot_call_float_object_vector2_string_string_color(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_float_object_vector2_string_string_color returned")
	ret := cRet
	return float64(ret)

}
func godotCallVoidRidColor(o Class, methodName string, arg0 *RID, arg1 *Color) {

	cArg0 := arg0.rid
	cArg1 := arg1.color

	log.Println("  Calling godot_call_void_rid_color...")
	C.godot_call_void_rid_color(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_rid_color returned")

}
func godotCallArray(o Class, methodName string) *Array {

	log.Println("  Calling godot_call_array...")
	cRet := C.godot_call_array(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_array returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_array)(cRet)
	return &Array{ret}

}
func godotCallIntStringStringInt(o Class, methodName string, arg0 string, arg1 string, arg2 int64) int64 {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_int_string_string_int...")
	cRet := C.godot_call_int_string_string_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_int_string_string_int returned")
	ret := cRet
	return int64(ret)

}
func godotCallRect2Int(o Class, methodName string, arg0 int64) *Rect2 {

	cArg0 := C.godot_int(arg0)

	log.Println("  Calling godot_call_rect2_int...")
	cRet := C.godot_call_rect2_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_rect2_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_rect2)(cRet)
	return &Rect2{ret}

}
func godotCallBoolTransformVector3(o Class, methodName string, arg0 *Transform, arg1 *Vector3) bool {

	cArg0 := arg0.transform
	cArg1 := arg1.vector3

	log.Println("  Calling godot_call_bool_transform_vector3...")
	cRet := C.godot_call_bool_transform_vector3(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_bool_transform_vector3 returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoidRidObjectString(o Class, methodName string, arg0 *RID, arg1 *Object, arg2 string) {

	cArg0 := arg0.rid
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(cArg2))

	log.Println("  Calling godot_call_void_rid_object_string...")
	C.godot_call_void_rid_object_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rid_object_string returned")

}
func godotCallPoolColorArray(o Class, methodName string) *PoolColorArray {

	log.Println("  Calling godot_call_poolcolorarray...")
	cRet := C.godot_call_poolcolorarray(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_poolcolorarray returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_pool_color_array)(cRet)
	return &PoolColorArray{ret}

}
func godotCallVoidVector2Vector2ColorFloatBool(o Class, methodName string, arg0 *Vector2, arg1 *Vector2, arg2 *Color, arg3 float64, arg4 bool) {

	cArg0 := arg0.vector2
	cArg1 := arg1.vector2
	cArg2 := arg2.color
	cArg3 := C.godot_real(arg3)
	cArg4 := C.godot_bool(arg4)

	log.Println("  Calling godot_call_void_vector2_vector2_color_float_bool...")
	C.godot_call_void_vector2_vector2_color_float_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_void_vector2_vector2_color_float_bool returned")

}
func godotCallVector2Vector2(o Class, methodName string, arg0 *Vector2) *Vector2 {

	cArg0 := arg0.vector2

	log.Println("  Calling godot_call_vector2_vector2...")
	cRet := C.godot_call_vector2_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_vector2_vector2 returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector2)(cRet)
	return &Vector2{ret}

}
func godotCallIntVector2FloatFloatIntInt(o Class, methodName string, arg0 *Vector2, arg1 float64, arg2 float64, arg3 int64, arg4 int64) int64 {

	cArg0 := arg0.vector2
	cArg1 := C.godot_real(arg1)
	cArg2 := C.godot_real(arg2)
	cArg3 := C.godot_int(arg3)
	cArg4 := C.godot_int(arg4)

	log.Println("  Calling godot_call_int_vector2_float_float_int_int...")
	cRet := C.godot_call_int_vector2_float_float_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_int_vector2_float_float_int_int returned")
	ret := cRet
	return int64(ret)

}
func godotCallVoidPoolRealArray(o Class, methodName string, arg0 *PoolRealArray) {

	cArg0 := arg0.array

	log.Println("  Calling godot_call_void_poolrealarray...")
	C.godot_call_void_poolrealarray(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_poolrealarray returned")

}
func godotCallVoidObjectStringIntInt(o Class, methodName string, arg0 *Object, arg1 string, arg2 int64, arg3 int64) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := C.godot_int(arg2)
	cArg3 := C.godot_int(arg3)

	log.Println("  Calling godot_call_void_object_string_int_int...")
	C.godot_call_void_object_string_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_object_string_int_int returned")

}
func godotCallVoidStringFloatFloatBool(o Class, methodName string, arg0 string, arg1 float64, arg2 float64, arg3 bool) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_real(arg1)
	cArg2 := C.godot_real(arg2)
	cArg3 := C.godot_bool(arg3)

	log.Println("  Calling godot_call_void_string_float_float_bool...")
	C.godot_call_void_string_float_float_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_string_float_float_bool returned")

}
func godotCallObjectIntInt(o Class, methodName string, arg0 int64, arg1 int64) *Object {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_object_int_int...")
	cRet := C.godot_call_object_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_object_int_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_object)(cRet)
	return &Object{ret}

}
func godotCallVector3FloatBool(o Class, methodName string, arg0 float64, arg1 bool) *Vector3 {

	cArg0 := C.godot_real(arg0)
	cArg1 := C.godot_bool(arg1)

	log.Println("  Calling godot_call_vector3_float_bool...")
	cRet := C.godot_call_vector3_float_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_vector3_float_bool returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector3)(cRet)
	return &Vector3{ret}

}
func godotCallColorFloat(o Class, methodName string, arg0 float64) *Color {

	cArg0 := C.godot_real(arg0)

	log.Println("  Calling godot_call_color_float...")
	cRet := C.godot_call_color_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_color_float returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_color)(cRet)
	return &Color{ret}

}
func godotCallVoidRidObjectStringVariant(o Class, methodName string, arg0 *RID, arg1 *Object, arg2 string, arg3 *Variant) {

	cArg0 := arg0.rid
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(cArg2))
	cArg3 := arg3.variant

	log.Println("  Calling godot_call_void_rid_object_string_variant...")
	C.godot_call_void_rid_object_string_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_rid_object_string_variant returned")

}
func godotCallBoolStringIntIntInt(o Class, methodName string, arg0 string, arg1 int64, arg2 int64, arg3 int64) bool {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_int(arg2)
	cArg3 := C.godot_int(arg3)

	log.Println("  Calling godot_call_bool_string_int_int_int...")
	cRet := C.godot_call_bool_string_int_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_bool_string_int_int_int returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoidRidAabb(o Class, methodName string, arg0 *RID, arg1 *AABB) {

	cArg0 := arg0.rid
	cArg1 := arg1.aabb

	log.Println("  Calling godot_call_void_rid_aabb...")
	C.godot_call_void_rid_aabb(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_rid_aabb returned")

}
func godotCallIntInt(o Class, methodName string, arg0 int64) int64 {

	cArg0 := C.godot_int(arg0)

	log.Println("  Calling godot_call_int_int...")
	cRet := C.godot_call_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_int_int returned")
	ret := cRet
	return int64(ret)

}
func godotCallVoidIntRidIntIntInt(o Class, methodName string, arg0 int64, arg1 *RID, arg2 int64, arg3 int64, arg4 int64) {

	cArg0 := C.godot_int(arg0)
	cArg1 := arg1.rid
	cArg2 := C.godot_int(arg2)
	cArg3 := C.godot_int(arg3)
	cArg4 := C.godot_int(arg4)

	log.Println("  Calling godot_call_void_int_rid_int_int_int...")
	C.godot_call_void_int_rid_int_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_void_int_rid_int_int_int returned")

}
func godotCallVoidRidVector2Vector2(o Class, methodName string, arg0 *RID, arg1 *Vector2, arg2 *Vector2) {

	cArg0 := arg0.rid
	cArg1 := arg1.vector2
	cArg2 := arg2.vector2

	log.Println("  Calling godot_call_void_rid_vector2_vector2...")
	C.godot_call_void_rid_vector2_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rid_vector2_vector2 returned")

}
func godotCallVoidRidIntFloat(o Class, methodName string, arg0 *RID, arg1 int64, arg2 float64) {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_real(arg2)

	log.Println("  Calling godot_call_void_rid_int_float...")
	C.godot_call_void_rid_int_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rid_int_float returned")

}
func godotCallRidRidTransformRidTransform(o Class, methodName string, arg0 *RID, arg1 *Transform, arg2 *RID, arg3 *Transform) *RID {

	cArg0 := arg0.rid
	cArg1 := arg1.transform
	cArg2 := arg2.rid
	cArg3 := arg3.transform

	log.Println("  Calling godot_call_rid_rid_transform_rid_transform...")
	cRet := C.godot_call_rid_rid_transform_rid_transform(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_rid_rid_transform_rid_transform returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_rid)(cRet)
	return &RID{ret}

}
func godotCallVoidStringObjectInt(o Class, methodName string, arg0 string, arg1 *Object, arg2 int64) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_void_string_object_int...")
	C.godot_call_void_string_object_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_string_object_int returned")

}
func godotCallObjectObjectInt(o Class, methodName string, arg0 *Object, arg1 int64) *Object {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_object_object_int...")
	cRet := C.godot_call_object_object_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_object_object_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_object)(cRet)
	return &Object{ret}

}
func godotCallVoidStringStringColorBool(o Class, methodName string, arg0 string, arg1 string, arg2 *Color, arg3 bool) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := arg2.color
	cArg3 := C.godot_bool(arg3)

	log.Println("  Calling godot_call_void_string_string_color_bool...")
	C.godot_call_void_string_string_color_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_string_string_color_bool returned")

}
func godotCallVoidIntVector3(o Class, methodName string, arg0 int64, arg1 *Vector3) {

	cArg0 := C.godot_int(arg0)
	cArg1 := arg1.vector3

	log.Println("  Calling godot_call_void_int_vector3...")
	C.godot_call_void_int_vector3(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_int_vector3 returned")

}
func godotCallVoidRidFloat(o Class, methodName string, arg0 *RID, arg1 float64) {

	cArg0 := arg0.rid
	cArg1 := C.godot_real(arg1)

	log.Println("  Calling godot_call_void_rid_float...")
	C.godot_call_void_rid_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_rid_float returned")

}
func godotCallRidRidVector3RidVector3(o Class, methodName string, arg0 *RID, arg1 *Vector3, arg2 *RID, arg3 *Vector3) *RID {

	cArg0 := arg0.rid
	cArg1 := arg1.vector3
	cArg2 := arg2.rid
	cArg3 := arg3.vector3

	log.Println("  Calling godot_call_rid_rid_vector3_rid_vector3...")
	cRet := C.godot_call_rid_rid_vector3_rid_vector3(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_rid_rid_vector3_rid_vector3 returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_rid)(cRet)
	return &RID{ret}

}
func godotCallVoidRidRect2(o Class, methodName string, arg0 *RID, arg1 *Rect2) {

	cArg0 := arg0.rid
	cArg1 := arg1.rect2

	log.Println("  Calling godot_call_void_rid_rect2...")
	C.godot_call_void_rid_rect2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_rid_rect2 returned")

}
func godotCallVoidIntObjectIntBoolString(o Class, methodName string, arg0 int64, arg1 *Object, arg2 int64, arg3 bool, arg4 string) {

	cArg0 := C.godot_int(arg0)
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := C.godot_int(arg2)
	cArg3 := C.godot_bool(arg3)
	cArg4 := C.CString(arg4)
	defer C.free(unsafe.Pointer(cArg4))

	log.Println("  Calling godot_call_void_int_object_int_bool_string...")
	C.godot_call_void_int_object_int_bool_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_void_int_object_int_bool_string returned")

}
func godotCallVector3Vector3(o Class, methodName string, arg0 *Vector3) *Vector3 {

	cArg0 := arg0.vector3

	log.Println("  Calling godot_call_vector3_vector3...")
	cRet := C.godot_call_vector3_vector3(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_vector3_vector3 returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector3)(cRet)
	return &Vector3{ret}

}
func godotCallBoolVector2(o Class, methodName string, arg0 *Vector2) bool {

	cArg0 := arg0.vector2

	log.Println("  Calling godot_call_bool_vector2...")
	cRet := C.godot_call_bool_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_bool_vector2 returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoidPoolVector2ArrayPoolColorArrayPoolVector2ArrayObjectObjectBool(o Class, methodName string, arg0 *PoolVector2Array, arg1 *PoolColorArray, arg2 *PoolVector2Array, arg3 *Object, arg4 *Object, arg5 bool) {

	cArg0 := arg0.array
	cArg1 := arg1.array
	cArg2 := arg2.array
	cArg3 := unsafe.Pointer(arg3.owner)
	cArg4 := unsafe.Pointer(arg4.owner)
	cArg5 := C.godot_bool(arg5)

	log.Println("  Calling godot_call_void_poolvector2array_poolcolorarray_poolvector2array_object_object_bool...")
	C.godot_call_void_poolvector2array_poolcolorarray_poolvector2array_object_object_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5,
	)
	log.Println("  ...godot_call_void_poolvector2array_poolcolorarray_poolvector2array_object_object_bool returned")

}
func godotCallArrayString(o Class, methodName string, arg0 string) *Array {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))

	log.Println("  Calling godot_call_array_string...")
	cRet := C.godot_call_array_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_array_string returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_array)(cRet)
	return &Array{ret}

}
func godotCallVoidIntPoolIntArray(o Class, methodName string, arg0 int64, arg1 *PoolIntArray) {

	cArg0 := C.godot_int(arg0)
	cArg1 := arg1.array

	log.Println("  Calling godot_call_void_int_poolintarray...")
	C.godot_call_void_int_poolintarray(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_int_poolintarray returned")

}
func godotCallPoolVector2ArrayInt(o Class, methodName string, arg0 int64) *PoolVector2Array {

	cArg0 := C.godot_int(arg0)

	log.Println("  Calling godot_call_poolvector2array_int...")
	cRet := C.godot_call_poolvector2array_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_poolvector2array_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_pool_vector2_array)(cRet)
	return &PoolVector2Array{ret}

}
func godotCallPoolIntArrayInt(o Class, methodName string, arg0 int64) *PoolIntArray {

	cArg0 := C.godot_int(arg0)

	log.Println("  Calling godot_call_poolintarray_int...")
	cRet := C.godot_call_poolintarray_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_poolintarray_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_pool_int_array)(cRet)
	return &PoolIntArray{ret}

}
func godotCallVoidVector2FloatVector2(o Class, methodName string, arg0 *Vector2, arg1 float64, arg2 *Vector2) {

	cArg0 := arg0.vector2
	cArg1 := C.godot_real(arg1)
	cArg2 := arg2.vector2

	log.Println("  Calling godot_call_void_vector2_float_vector2...")
	C.godot_call_void_vector2_float_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_vector2_float_vector2 returned")

}
func godotCallVoidIntTransform2D(o Class, methodName string, arg0 int64, arg1 *Transform2D) {

	cArg0 := C.godot_int(arg0)
	cArg1 := arg1.transform2d

	log.Println("  Calling godot_call_void_int_transform2d...")
	C.godot_call_void_int_transform2d(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_int_transform2d returned")

}
func godotCallColorInt(o Class, methodName string, arg0 int64) *Color {

	cArg0 := C.godot_int(arg0)

	log.Println("  Calling godot_call_color_int...")
	cRet := C.godot_call_color_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_color_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_color)(cRet)
	return &Color{ret}

}
func godotCallVoidRidIntTransform2D(o Class, methodName string, arg0 *RID, arg1 int64, arg2 *Transform2D) {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)
	cArg2 := arg2.transform2d

	log.Println("  Calling godot_call_void_rid_int_transform2d...")
	C.godot_call_void_rid_int_transform2d(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rid_int_transform2d returned")

}
func godotCallVariantVariant(o Class, methodName string, arg0 *Variant) *Variant {

	cArg0 := arg0.variant

	log.Println("  Calling godot_call_variant_variant...")
	cRet := C.godot_call_variant_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_variant_variant returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_variant)(cRet)
	return &Variant{ret}

}
func godotCallVoidIntVariant(o Class, methodName string, arg0 int64, arg1 *Variant) {

	cArg0 := C.godot_int(arg0)
	cArg1 := arg1.variant

	log.Println("  Calling godot_call_void_int_variant...")
	C.godot_call_void_int_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_int_variant returned")

}
func godotCallVoidRidRect2RidRect2ColorBoolRidBool(o Class, methodName string, arg0 *RID, arg1 *Rect2, arg2 *RID, arg3 *Rect2, arg4 *Color, arg5 bool, arg6 *RID, arg7 bool) {

	cArg0 := arg0.rid
	cArg1 := arg1.rect2
	cArg2 := arg2.rid
	cArg3 := arg3.rect2
	cArg4 := arg4.color
	cArg5 := C.godot_bool(arg5)
	cArg6 := arg6.rid
	cArg7 := C.godot_bool(arg7)

	log.Println("  Calling godot_call_void_rid_rect2_rid_rect2_color_bool_rid_bool...")
	C.godot_call_void_rid_rect2_rid_rect2_color_bool_rid_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5, cArg6, cArg7,
	)
	log.Println("  ...godot_call_void_rid_rect2_rid_rect2_color_bool_rid_bool returned")

}
func godotCallVariant(o Class, methodName string) *Variant {

	log.Println("  Calling godot_call_variant...")
	cRet := C.godot_call_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_variant returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_variant)(cRet)
	return &Variant{ret}

}
func godotCallVoidVector3Vector3(o Class, methodName string, arg0 *Vector3, arg1 *Vector3) {

	cArg0 := arg0.vector3
	cArg1 := arg1.vector3

	log.Println("  Calling godot_call_void_vector3_vector3...")
	C.godot_call_void_vector3_vector3(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_vector3_vector3 returned")

}
func godotCallBool(o Class, methodName string) bool {

	log.Println("  Calling godot_call_bool...")
	cRet := C.godot_call_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_bool returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoidPoolVector2ArrayPoolColorArrayPoolVector2ArrayObjectFloatObject(o Class, methodName string, arg0 *PoolVector2Array, arg1 *PoolColorArray, arg2 *PoolVector2Array, arg3 *Object, arg4 float64, arg5 *Object) {

	cArg0 := arg0.array
	cArg1 := arg1.array
	cArg2 := arg2.array
	cArg3 := unsafe.Pointer(arg3.owner)
	cArg4 := C.godot_real(arg4)
	cArg5 := unsafe.Pointer(arg5.owner)

	log.Println("  Calling godot_call_void_poolvector2array_poolcolorarray_poolvector2array_object_float_object...")
	C.godot_call_void_poolvector2array_poolcolorarray_poolvector2array_object_float_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5,
	)
	log.Println("  ...godot_call_void_poolvector2array_poolcolorarray_poolvector2array_object_float_object returned")

}
func godotCallIntVector3(o Class, methodName string, arg0 *Vector3) int64 {

	cArg0 := arg0.vector3

	log.Println("  Calling godot_call_int_vector3...")
	cRet := C.godot_call_int_vector3(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_int_vector3 returned")
	ret := cRet
	return int64(ret)

}
func godotCallVoidPoolVector2Array(o Class, methodName string, arg0 *PoolVector2Array) {

	cArg0 := arg0.array

	log.Println("  Calling godot_call_void_poolvector2array...")
	C.godot_call_void_poolvector2array(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_poolvector2array returned")

}
func godotCallArrayObjectInt(o Class, methodName string, arg0 *Object, arg1 int64) *Array {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_array_object_int...")
	cRet := C.godot_call_array_object_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_array_object_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_array)(cRet)
	return &Array{ret}

}
func godotCallStringVariant(o Class, methodName string, arg0 *Variant) string {

	cArg0 := arg0.variant

	log.Println("  Calling godot_call_string_variant...")
	cRet := C.godot_call_string_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_string_variant returned")
	return C.GoString(cRet)

}
func godotCallBoolVector2Rect2(o Class, methodName string, arg0 *Vector2, arg1 *Rect2) bool {

	cArg0 := arg0.vector2
	cArg1 := arg1.rect2

	log.Println("  Calling godot_call_bool_vector2_rect2...")
	cRet := C.godot_call_bool_vector2_rect2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_bool_vector2_rect2 returned")
	ret := cRet
	return bool(ret)

}
func godotCallIntIntFloatBool(o Class, methodName string, arg0 int64, arg1 float64, arg2 bool) int64 {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_real(arg1)
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_int_int_float_bool...")
	cRet := C.godot_call_int_int_float_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_int_int_float_bool returned")
	ret := cRet
	return int64(ret)

}
func godotCallVoidObjectObjectIntBool(o Class, methodName string, arg0 *Object, arg1 *Object, arg2 int64, arg3 bool) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := C.godot_int(arg2)
	cArg3 := C.godot_bool(arg3)

	log.Println("  Calling godot_call_void_object_object_int_bool...")
	C.godot_call_void_object_object_int_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_object_object_int_bool returned")

}
func godotCallRidIntIntFloat(o Class, methodName string, arg0 int64, arg1 int64, arg2 float64) *RID {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_real(arg2)

	log.Println("  Calling godot_call_rid_int_int_float...")
	cRet := C.godot_call_rid_int_int_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_rid_int_int_float returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_rid)(cRet)
	return &RID{ret}

}
func godotCallInt(o Class, methodName string) int64 {

	log.Println("  Calling godot_call_int...")
	cRet := C.godot_call_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_int returned")
	ret := cRet
	return int64(ret)

}
func godotCallVoidStringArray(o Class, methodName string, arg0 string, arg1 *Array) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := arg1.array

	log.Println("  Calling godot_call_void_string_array...")
	C.godot_call_void_string_array(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_string_array returned")

}
func godotCallVoidRidVector2(o Class, methodName string, arg0 *RID, arg1 *Vector2) {

	cArg0 := arg0.rid
	cArg1 := arg1.vector2

	log.Println("  Calling godot_call_void_rid_vector2...")
	C.godot_call_void_rid_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_rid_vector2 returned")

}
func godotCallVoidObjectStringBool(o Class, methodName string, arg0 *Object, arg1 string, arg2 bool) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_void_object_string_bool...")
	C.godot_call_void_object_string_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_object_string_bool returned")

}
func godotCallObjectStringBoolString(o Class, methodName string, arg0 string, arg1 bool, arg2 string) *Object {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_bool(arg1)
	cArg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(cArg2))

	log.Println("  Calling godot_call_object_string_bool_string...")
	cRet := C.godot_call_object_string_bool_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_object_string_bool_string returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_object)(cRet)
	return &Object{ret}

}
func godotCallVariantIntString(o Class, methodName string, arg0 int64, arg1 string) *Variant {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))

	log.Println("  Calling godot_call_variant_int_string...")
	cRet := C.godot_call_variant_int_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_variant_int_string returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_variant)(cRet)
	return &Variant{ret}

}
func godotCallRidRidInt(o Class, methodName string, arg0 *RID, arg1 int64) *RID {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_rid_rid_int...")
	cRet := C.godot_call_rid_rid_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_rid_rid_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_rid)(cRet)
	return &RID{ret}

}
func godotCallVoidRidVector3(o Class, methodName string, arg0 *RID, arg1 *Vector3) {

	cArg0 := arg0.rid
	cArg1 := arg1.vector3

	log.Println("  Calling godot_call_void_rid_vector3...")
	C.godot_call_void_rid_vector3(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_rid_vector3 returned")

}
func godotCallVariantStringString(o Class, methodName string, arg0 string, arg1 string) *Variant {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))

	log.Println("  Calling godot_call_variant_string_string...")
	cRet := C.godot_call_variant_string_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_variant_string_string returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_variant)(cRet)
	return &Variant{ret}

}
func godotCallVariantIntStringString(o Class, methodName string, arg0 int64, arg1 string, arg2 string) *Variant {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(cArg2))

	log.Println("  Calling godot_call_variant_int_string_string...")
	cRet := C.godot_call_variant_int_string_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_variant_int_string_string returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_variant)(cRet)
	return &Variant{ret}

}
func godotCallVoidTransform(o Class, methodName string, arg0 *Transform) {

	cArg0 := arg0.transform

	log.Println("  Calling godot_call_void_transform...")
	C.godot_call_void_transform(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_transform returned")

}
func godotCallVoidRect2ColorBool(o Class, methodName string, arg0 *Rect2, arg1 *Color, arg2 bool) {

	cArg0 := arg0.rect2
	cArg1 := arg1.color
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_void_rect2_color_bool...")
	C.godot_call_void_rect2_color_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rect2_color_bool returned")

}
func godotCallRidInt(o Class, methodName string, arg0 int64) *RID {

	cArg0 := C.godot_int(arg0)

	log.Println("  Calling godot_call_rid_int...")
	cRet := C.godot_call_rid_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_rid_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_rid)(cRet)
	return &RID{ret}

}
func godotCallBoolRidIntInt(o Class, methodName string, arg0 *RID, arg1 int64, arg2 int64) bool {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_bool_rid_int_int...")
	cRet := C.godot_call_bool_rid_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_bool_rid_int_int returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoidRidRidTransform(o Class, methodName string, arg0 *RID, arg1 *RID, arg2 *Transform) {

	cArg0 := arg0.rid
	cArg1 := arg1.rid
	cArg2 := arg2.transform

	log.Println("  Calling godot_call_void_rid_rid_transform...")
	C.godot_call_void_rid_rid_transform(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rid_rid_transform returned")

}
func godotCallObjectFloatBool(o Class, methodName string, arg0 float64, arg1 bool) *Object {

	cArg0 := C.godot_real(arg0)
	cArg1 := C.godot_bool(arg1)

	log.Println("  Calling godot_call_object_float_bool...")
	cRet := C.godot_call_object_float_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_object_float_bool returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_object)(cRet)
	return &Object{ret}

}
func godotCallVoidIntBool(o Class, methodName string, arg0 int64, arg1 bool) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_bool(arg1)

	log.Println("  Calling godot_call_void_int_bool...")
	C.godot_call_void_int_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_int_bool returned")

}
func godotCallBoolStringDictionary(o Class, methodName string, arg0 string, arg1 *Dictionary) bool {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := arg1.dictionary

	log.Println("  Calling godot_call_bool_string_dictionary...")
	cRet := C.godot_call_bool_string_dictionary(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_bool_string_dictionary returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoidRidIntIntBool(o Class, methodName string, arg0 *RID, arg1 int64, arg2 int64, arg3 bool) {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_int(arg2)
	cArg3 := C.godot_bool(arg3)

	log.Println("  Calling godot_call_void_rid_int_int_bool...")
	C.godot_call_void_rid_int_int_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_rid_int_int_bool returned")

}
func godotCallVoidNodePath(o Class, methodName string, arg0 *NodePath) {

	cArg0 := arg0.nodePath

	log.Println("  Calling godot_call_void_nodepath...")
	C.godot_call_void_nodepath(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_nodepath returned")

}
func godotCallVoidVariantObject(o Class, methodName string, arg0 *Variant, arg1 *Object) {

	cArg0 := arg0.variant
	cArg1 := unsafe.Pointer(arg1.owner)

	log.Println("  Calling godot_call_void_variant_object...")
	C.godot_call_void_variant_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_variant_object returned")

}
func godotCallIntObjectTransform2DObject(o Class, methodName string, arg0 *Object, arg1 *Transform2D, arg2 *Object) int64 {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := arg1.transform2d
	cArg2 := unsafe.Pointer(arg2.owner)

	log.Println("  Calling godot_call_int_object_transform2d_object...")
	cRet := C.godot_call_int_object_transform2d_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_int_object_transform2d_object returned")
	ret := cRet
	return int64(ret)

}
func godotCallArrayObject(o Class, methodName string, arg0 *Object) *Array {

	cArg0 := unsafe.Pointer(arg0.owner)

	log.Println("  Calling godot_call_array_object...")
	cRet := C.godot_call_array_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_array_object returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_array)(cRet)
	return &Array{ret}

}
func godotCallObjectRid(o Class, methodName string, arg0 *RID) *Object {

	cArg0 := arg0.rid

	log.Println("  Calling godot_call_object_rid...")
	cRet := C.godot_call_object_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_object_rid returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_object)(cRet)
	return &Object{ret}

}
func godotCallStringStringInt(o Class, methodName string, arg0 string, arg1 int64) string {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_string_string_int...")
	cRet := C.godot_call_string_string_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_string_string_int returned")
	return C.GoString(cRet)

}
func godotCallVoidPoolStringArrayInt(o Class, methodName string, arg0 *PoolStringArray, arg1 int64) {

	cArg0 := arg0.array
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_void_poolstringarray_int...")
	C.godot_call_void_poolstringarray_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_poolstringarray_int returned")

}
func godotCallFloatRid(o Class, methodName string, arg0 *RID) float64 {

	cArg0 := arg0.rid

	log.Println("  Calling godot_call_float_rid...")
	cRet := C.godot_call_float_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_float_rid returned")
	ret := cRet
	return float64(ret)

}
func godotCallVoidIntStringInt(o Class, methodName string, arg0 int64, arg1 string, arg2 int64) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_void_int_string_int...")
	C.godot_call_void_int_string_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_int_string_int returned")

}
func godotCallVoidBoolBool(o Class, methodName string, arg0 bool, arg1 bool) {

	cArg0 := C.godot_bool(arg0)
	cArg1 := C.godot_bool(arg1)

	log.Println("  Calling godot_call_void_bool_bool...")
	C.godot_call_void_bool_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_bool_bool returned")

}
func godotCallObjectString(o Class, methodName string, arg0 string) *Object {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))

	log.Println("  Calling godot_call_object_string...")
	cRet := C.godot_call_object_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_object_string returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_object)(cRet)
	return &Object{ret}

}
func godotCallVector2FloatBool(o Class, methodName string, arg0 float64, arg1 bool) *Vector2 {

	cArg0 := C.godot_real(arg0)
	cArg1 := C.godot_bool(arg1)

	log.Println("  Calling godot_call_vector2_float_bool...")
	cRet := C.godot_call_vector2_float_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_vector2_float_bool returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector2)(cRet)
	return &Vector2{ret}

}
func godotCallArrayVector2IntArrayInt(o Class, methodName string, arg0 *Vector2, arg1 int64, arg2 *Array, arg3 int64) *Array {

	cArg0 := arg0.vector2
	cArg1 := C.godot_int(arg1)
	cArg2 := arg2.array
	cArg3 := C.godot_int(arg3)

	log.Println("  Calling godot_call_array_vector2_int_array_int...")
	cRet := C.godot_call_array_vector2_int_array_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_array_vector2_int_array_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_array)(cRet)
	return &Array{ret}

}
func godotCallVoidRidVariant(o Class, methodName string, arg0 *RID, arg1 *Variant) {

	cArg0 := arg0.rid
	cArg1 := arg1.variant

	log.Println("  Calling godot_call_void_rid_variant...")
	C.godot_call_void_rid_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_rid_variant returned")

}
func godotCallBasis(o Class, methodName string) *Basis {

	log.Println("  Calling godot_call_basis...")
	cRet := C.godot_call_basis(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_basis returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_basis)(cRet)
	return &Basis{ret}

}
func godotCallRidVector2Vector2RidRid(o Class, methodName string, arg0 *Vector2, arg1 *Vector2, arg2 *RID, arg3 *RID) *RID {

	cArg0 := arg0.vector2
	cArg1 := arg1.vector2
	cArg2 := arg2.rid
	cArg3 := arg3.rid

	log.Println("  Calling godot_call_rid_vector2_vector2_rid_rid...")
	cRet := C.godot_call_rid_vector2_vector2_rid_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_rid_vector2_vector2_rid_rid returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_rid)(cRet)
	return &RID{ret}

}
func godotCallVoidRidPoolVector2ArrayPoolColorArrayFloatBool(o Class, methodName string, arg0 *RID, arg1 *PoolVector2Array, arg2 *PoolColorArray, arg3 float64, arg4 bool) {

	cArg0 := arg0.rid
	cArg1 := arg1.array
	cArg2 := arg2.array
	cArg3 := C.godot_real(arg3)
	cArg4 := C.godot_bool(arg4)

	log.Println("  Calling godot_call_void_rid_poolvector2array_poolcolorarray_float_bool...")
	C.godot_call_void_rid_poolvector2array_poolcolorarray_float_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_void_rid_poolvector2array_poolcolorarray_float_bool returned")

}
func godotCallIntVector2(o Class, methodName string, arg0 *Vector2) int64 {

	cArg0 := arg0.vector2

	log.Println("  Calling godot_call_int_vector2...")
	cRet := C.godot_call_int_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_int_vector2 returned")
	ret := cRet
	return int64(ret)

}
func godotCallVoidIntFloatFloatFloat(o Class, methodName string, arg0 int64, arg1 float64, arg2 float64, arg3 float64) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_real(arg1)
	cArg2 := C.godot_real(arg2)
	cArg3 := C.godot_real(arg3)

	log.Println("  Calling godot_call_void_int_float_float_float...")
	C.godot_call_void_int_float_float_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_int_float_float_float returned")

}
func godotCallVoidRidPoolVector2Array(o Class, methodName string, arg0 *RID, arg1 *PoolVector2Array) {

	cArg0 := arg0.rid
	cArg1 := arg1.array

	log.Println("  Calling godot_call_void_rid_poolvector2array...")
	C.godot_call_void_rid_poolvector2array(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_rid_poolvector2array returned")

}
func godotCallVoidObjectColorBool(o Class, methodName string, arg0 *Object, arg1 *Color, arg2 bool) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := arg1.color
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_void_object_color_bool...")
	C.godot_call_void_object_color_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_object_color_bool returned")

}
func godotCallVoidAabb(o Class, methodName string, arg0 *AABB) {

	cArg0 := arg0.aabb

	log.Println("  Calling godot_call_void_aabb...")
	C.godot_call_void_aabb(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_aabb returned")

}
func godotCallBoolRidInt(o Class, methodName string, arg0 *RID, arg1 int64) bool {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_bool_rid_int...")
	cRet := C.godot_call_bool_rid_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_bool_rid_int returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoidRidRidVector2(o Class, methodName string, arg0 *RID, arg1 *RID, arg2 *Vector2) {

	cArg0 := arg0.rid
	cArg1 := arg1.rid
	cArg2 := arg2.vector2

	log.Println("  Calling godot_call_void_rid_rid_vector2...")
	C.godot_call_void_rid_rid_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rid_rid_vector2 returned")

}
func godotCallVoidPoolVector2ArrayPoolColorArrayFloatBool(o Class, methodName string, arg0 *PoolVector2Array, arg1 *PoolColorArray, arg2 float64, arg3 bool) {

	cArg0 := arg0.array
	cArg1 := arg1.array
	cArg2 := C.godot_real(arg2)
	cArg3 := C.godot_bool(arg3)

	log.Println("  Calling godot_call_void_poolvector2array_poolcolorarray_float_bool...")
	C.godot_call_void_poolvector2array_poolcolorarray_float_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_poolvector2array_poolcolorarray_float_bool returned")

}
func godotCallPoolVector3ArrayIntFloat(o Class, methodName string, arg0 int64, arg1 float64) *PoolVector3Array {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_real(arg1)

	log.Println("  Calling godot_call_poolvector3array_int_float...")
	cRet := C.godot_call_poolvector3array_int_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_poolvector3array_int_float returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_pool_vector3_array)(cRet)
	return &PoolVector3Array{ret}

}
func godotCallStringStringObject(o Class, methodName string, arg0 string, arg1 *Object) string {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := unsafe.Pointer(arg1.owner)

	log.Println("  Calling godot_call_string_string_object...")
	cRet := C.godot_call_string_string_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_string_string_object returned")
	return C.GoString(cRet)

}
func godotCallIntIntInt(o Class, methodName string, arg0 int64, arg1 int64) int64 {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_int_int_int...")
	cRet := C.godot_call_int_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_int_int_int returned")
	ret := cRet
	return int64(ret)

}
func godotCallPoolStringArrayString(o Class, methodName string, arg0 string) *PoolStringArray {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))

	log.Println("  Calling godot_call_poolstringarray_string...")
	cRet := C.godot_call_poolstringarray_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_poolstringarray_string returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_pool_string_array)(cRet)
	return &PoolStringArray{ret}

}
func godotCallDictionaryBool(o Class, methodName string, arg0 bool) *Dictionary {

	cArg0 := C.godot_bool(arg0)

	log.Println("  Calling godot_call_dictionary_bool...")
	cRet := C.godot_call_dictionary_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_dictionary_bool returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_dictionary)(cRet)
	return &Dictionary{ret}

}
func godotCallIntVector2Bool(o Class, methodName string, arg0 *Vector2, arg1 bool) int64 {

	cArg0 := arg0.vector2
	cArg1 := C.godot_bool(arg1)

	log.Println("  Calling godot_call_int_vector2_bool...")
	cRet := C.godot_call_int_vector2_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_int_vector2_bool returned")
	ret := cRet
	return int64(ret)

}
func godotCallVoidStringArrayBool(o Class, methodName string, arg0 string, arg1 *Array, arg2 bool) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := arg1.array
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_void_string_array_bool...")
	C.godot_call_void_string_array_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_string_array_bool returned")

}
func godotCallVoidRidIntBool(o Class, methodName string, arg0 *RID, arg1 int64, arg2 bool) {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_void_rid_int_bool...")
	C.godot_call_void_rid_int_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rid_int_bool returned")

}
func godotCallBoolRid(o Class, methodName string, arg0 *RID) bool {

	cArg0 := arg0.rid

	log.Println("  Calling godot_call_bool_rid...")
	cRet := C.godot_call_bool_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_bool_rid returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoidRidPoolVector2ArrayBool(o Class, methodName string, arg0 *RID, arg1 *PoolVector2Array, arg2 bool) {

	cArg0 := arg0.rid
	cArg1 := arg1.array
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_void_rid_poolvector2array_bool...")
	C.godot_call_void_rid_poolvector2array_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rid_poolvector2array_bool returned")

}
func godotCallBoolString(o Class, methodName string, arg0 string) bool {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))

	log.Println("  Calling godot_call_bool_string...")
	cRet := C.godot_call_bool_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_bool_string returned")
	ret := cRet
	return bool(ret)

}
func godotCallPoolRealArrayInt(o Class, methodName string, arg0 int64) *PoolRealArray {

	cArg0 := C.godot_int(arg0)

	log.Println("  Calling godot_call_poolrealarray_int...")
	cRet := C.godot_call_poolrealarray_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_poolrealarray_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_pool_real_array)(cRet)
	return &PoolRealArray{ret}

}
func godotCallVoidIntPoolVector2Array(o Class, methodName string, arg0 int64, arg1 *PoolVector2Array) {

	cArg0 := C.godot_int(arg0)
	cArg1 := arg1.array

	log.Println("  Calling godot_call_void_int_poolvector2array...")
	C.godot_call_void_int_poolvector2array(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_int_poolvector2array returned")

}
func godotCallVoidIntIntVector2Float(o Class, methodName string, arg0 int64, arg1 int64, arg2 *Vector2, arg3 float64) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := arg2.vector2
	cArg3 := C.godot_real(arg3)

	log.Println("  Calling godot_call_void_int_int_vector2_float...")
	C.godot_call_void_int_int_vector2_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_int_int_vector2_float returned")

}
func godotCallVoidVector3Vector3Vector3Int(o Class, methodName string, arg0 *Vector3, arg1 *Vector3, arg2 *Vector3, arg3 int64) {

	cArg0 := arg0.vector3
	cArg1 := arg1.vector3
	cArg2 := arg2.vector3
	cArg3 := C.godot_int(arg3)

	log.Println("  Calling godot_call_void_vector3_vector3_vector3_int...")
	C.godot_call_void_vector3_vector3_vector3_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_vector3_vector3_vector3_int returned")

}
func godotCallVoidObjectBool(o Class, methodName string, arg0 *Object, arg1 bool) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.godot_bool(arg1)

	log.Println("  Calling godot_call_void_object_bool...")
	C.godot_call_void_object_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_object_bool returned")

}
func godotCallAabbRid(o Class, methodName string, arg0 *RID) *AABB {

	cArg0 := arg0.rid

	log.Println("  Calling godot_call_aabb_rid...")
	cRet := C.godot_call_aabb_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_aabb_rid returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_aabb)(cRet)
	return &AABB{ret}

}
func godotCallVoidRidIntIntIntInt(o Class, methodName string, arg0 *RID, arg1 int64, arg2 int64, arg3 int64, arg4 int64) {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_int(arg2)
	cArg3 := C.godot_int(arg3)
	cArg4 := C.godot_int(arg4)

	log.Println("  Calling godot_call_void_rid_int_int_int_int...")
	C.godot_call_void_rid_int_int_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_void_rid_int_int_int_int returned")

}
func godotCallIntVariant(o Class, methodName string, arg0 *Variant) int64 {

	cArg0 := arg0.variant

	log.Println("  Calling godot_call_int_variant...")
	cRet := C.godot_call_int_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_int_variant returned")
	ret := cRet
	return int64(ret)

}
func godotCallPoolVector3ArrayIntInt(o Class, methodName string, arg0 int64, arg1 int64) *PoolVector3Array {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_poolvector3array_int_int...")
	cRet := C.godot_call_poolvector3array_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_poolvector3array_int_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_pool_vector3_array)(cRet)
	return &PoolVector3Array{ret}

}
func godotCallDictionary(o Class, methodName string) *Dictionary {

	log.Println("  Calling godot_call_dictionary...")
	cRet := C.godot_call_dictionary(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_dictionary returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_dictionary)(cRet)
	return &Dictionary{ret}

}
func godotCallTransform2D(o Class, methodName string) *Transform2D {

	log.Println("  Calling godot_call_transform2d...")
	cRet := C.godot_call_transform2d(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_transform2d returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_transform2d)(cRet)
	return &Transform2D{ret}

}
func godotCallVariantStringStringVariant(o Class, methodName string, arg0 string, arg1 string, arg2 *Variant) *Variant {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := arg2.variant

	log.Println("  Calling godot_call_variant_string_string_variant...")
	cRet := C.godot_call_variant_string_string_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_variant_string_string_variant returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_variant)(cRet)
	return &Variant{ret}

}
func godotCallIntStringIntIntInt(o Class, methodName string, arg0 string, arg1 int64, arg2 int64, arg3 int64) int64 {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_int(arg2)
	cArg3 := C.godot_int(arg3)

	log.Println("  Calling godot_call_int_string_int_int_int...")
	cRet := C.godot_call_int_string_int_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_int_string_int_int_int returned")
	ret := cRet
	return int64(ret)

}
func godotCallObjectBool(o Class, methodName string, arg0 bool) *Object {

	cArg0 := C.godot_bool(arg0)

	log.Println("  Calling godot_call_object_bool...")
	cRet := C.godot_call_object_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_object_bool returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_object)(cRet)
	return &Object{ret}

}
func godotCallVariantArrayArrayIntArray(o Class, methodName string, arg0 *Array, arg1 *Array, arg2 int64, arg3 *Array) *Variant {

	cArg0 := arg0.array
	cArg1 := arg1.array
	cArg2 := C.godot_int(arg2)
	cArg3 := arg3.array

	log.Println("  Calling godot_call_variant_array_array_int_array...")
	cRet := C.godot_call_variant_array_array_int_array(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_variant_array_array_int_array returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_variant)(cRet)
	return &Variant{ret}

}
func godotCallStringString(o Class, methodName string, arg0 string) string {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))

	log.Println("  Calling godot_call_string_string...")
	cRet := C.godot_call_string_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_string_string returned")
	return C.GoString(cRet)

}
func godotCallBoolTransform2DVector2ObjectTransform2DVector2(o Class, methodName string, arg0 *Transform2D, arg1 *Vector2, arg2 *Object, arg3 *Transform2D, arg4 *Vector2) bool {

	cArg0 := arg0.transform2d
	cArg1 := arg1.vector2
	cArg2 := unsafe.Pointer(arg2.owner)
	cArg3 := arg3.transform2d
	cArg4 := arg4.vector2

	log.Println("  Calling godot_call_bool_transform2d_vector2_object_transform2d_vector2...")
	cRet := C.godot_call_bool_transform2d_vector2_object_transform2d_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_bool_transform2d_vector2_object_transform2d_vector2 returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoidDictionary(o Class, methodName string, arg0 *Dictionary) {

	cArg0 := arg0.dictionary

	log.Println("  Calling godot_call_void_dictionary...")
	C.godot_call_void_dictionary(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_dictionary returned")

}
func godotCallDictionaryObject(o Class, methodName string, arg0 *Object) *Dictionary {

	cArg0 := unsafe.Pointer(arg0.owner)

	log.Println("  Calling godot_call_dictionary_object...")
	cRet := C.godot_call_dictionary_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_dictionary_object returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_dictionary)(cRet)
	return &Dictionary{ret}

}
func godotCallVoidIntBoolBool(o Class, methodName string, arg0 int64, arg1 bool, arg2 bool) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_bool(arg1)
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_void_int_bool_bool...")
	C.godot_call_void_int_bool_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_int_bool_bool returned")

}
func godotCallStringObject(o Class, methodName string, arg0 *Object) string {

	cArg0 := unsafe.Pointer(arg0.owner)

	log.Println("  Calling godot_call_string_object...")
	cRet := C.godot_call_string_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_string_object returned")
	return C.GoString(cRet)

}
func godotCallIntStringString(o Class, methodName string, arg0 string, arg1 string) int64 {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))

	log.Println("  Calling godot_call_int_string_string...")
	cRet := C.godot_call_int_string_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_int_string_string returned")
	ret := cRet
	return int64(ret)

}
func godotCallBoolStringInt(o Class, methodName string, arg0 string, arg1 int64) bool {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_bool_string_int...")
	cRet := C.godot_call_bool_string_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_bool_string_int returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoidObjectRect2BoolColorBoolObject(o Class, methodName string, arg0 *Object, arg1 *Rect2, arg2 bool, arg3 *Color, arg4 bool, arg5 *Object) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := arg1.rect2
	cArg2 := C.godot_bool(arg2)
	cArg3 := arg3.color
	cArg4 := C.godot_bool(arg4)
	cArg5 := unsafe.Pointer(arg5.owner)

	log.Println("  Calling godot_call_void_object_rect2_bool_color_bool_object...")
	C.godot_call_void_object_rect2_bool_color_bool_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5,
	)
	log.Println("  ...godot_call_void_object_rect2_bool_color_bool_object returned")

}
func godotCallVoidIntIntBoolInt(o Class, methodName string, arg0 int64, arg1 int64, arg2 bool, arg3 int64) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_bool(arg2)
	cArg3 := C.godot_int(arg3)

	log.Println("  Calling godot_call_void_int_int_bool_int...")
	C.godot_call_void_int_int_bool_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_int_int_bool_int returned")

}
func godotCallTransform2DRidInt(o Class, methodName string, arg0 *RID, arg1 int64) *Transform2D {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_transform2d_rid_int...")
	cRet := C.godot_call_transform2d_rid_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_transform2d_rid_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_transform2d)(cRet)
	return &Transform2D{ret}

}
func godotCallVoidColor(o Class, methodName string, arg0 *Color) {

	cArg0 := arg0.color

	log.Println("  Calling godot_call_void_color...")
	C.godot_call_void_color(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_color returned")

}
func godotCallVoidPoolVector3ArrayBoolBool(o Class, methodName string, arg0 *PoolVector3Array, arg1 bool, arg2 bool) {

	cArg0 := arg0.array
	cArg1 := C.godot_bool(arg1)
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_void_poolvector3array_bool_bool...")
	C.godot_call_void_poolvector3array_bool_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_poolvector3array_bool_bool returned")

}
func godotCallVariantStringBool(o Class, methodName string, arg0 string, arg1 bool) *Variant {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_bool(arg1)

	log.Println("  Calling godot_call_variant_string_bool...")
	cRet := C.godot_call_variant_string_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_variant_string_bool returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_variant)(cRet)
	return &Variant{ret}

}
func godotCallStringStringStringBoolIntInt(o Class, methodName string, arg0 string, arg1 string, arg2 bool, arg3 int64, arg4 int64) string {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := C.godot_bool(arg2)
	cArg3 := C.godot_int(arg3)
	cArg4 := C.godot_int(arg4)

	log.Println("  Calling godot_call_string_string_string_bool_int_int...")
	cRet := C.godot_call_string_string_string_bool_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_string_string_string_bool_int_int returned")
	return C.GoString(cRet)

}
func godotCallVoidRidRidRidRid(o Class, methodName string, arg0 *RID, arg1 *RID, arg2 *RID, arg3 *RID) {

	cArg0 := arg0.rid
	cArg1 := arg1.rid
	cArg2 := arg2.rid
	cArg3 := arg3.rid

	log.Println("  Calling godot_call_void_rid_rid_rid_rid...")
	C.godot_call_void_rid_rid_rid_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_rid_rid_rid_rid returned")

}
func godotCallVoidTransform2D(o Class, methodName string, arg0 *Transform2D) {

	cArg0 := arg0.transform2d

	log.Println("  Calling godot_call_void_transform2d...")
	C.godot_call_void_transform2d(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_transform2d returned")

}
func godotCallFloatFloat(o Class, methodName string, arg0 float64) float64 {

	cArg0 := C.godot_real(arg0)

	log.Println("  Calling godot_call_float_float...")
	cRet := C.godot_call_float_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_float_float returned")
	ret := cRet
	return float64(ret)

}
func godotCallColorIntInt(o Class, methodName string, arg0 int64, arg1 int64) *Color {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_color_int_int...")
	cRet := C.godot_call_color_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_color_int_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_color)(cRet)
	return &Color{ret}

}
func godotCallVoidIntRect2(o Class, methodName string, arg0 int64, arg1 *Rect2) {

	cArg0 := C.godot_int(arg0)
	cArg1 := arg1.rect2

	log.Println("  Calling godot_call_void_int_rect2...")
	C.godot_call_void_int_rect2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_int_rect2 returned")

}
func godotCallVoidIntPlane(o Class, methodName string, arg0 int64, arg1 *Plane) {

	cArg0 := C.godot_int(arg0)
	cArg1 := arg1.plane

	log.Println("  Calling godot_call_void_int_plane...")
	C.godot_call_void_int_plane(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_int_plane returned")

}
func godotCallVoidStringIntVector2(o Class, methodName string, arg0 string, arg1 int64, arg2 *Vector2) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)
	cArg2 := arg2.vector2

	log.Println("  Calling godot_call_void_string_int_vector2...")
	C.godot_call_void_string_int_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_string_int_vector2 returned")

}
func godotCallVoidRidIntTransform(o Class, methodName string, arg0 *RID, arg1 int64, arg2 *Transform) {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)
	cArg2 := arg2.transform

	log.Println("  Calling godot_call_void_rid_int_transform...")
	C.godot_call_void_rid_int_transform(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rid_int_transform returned")

}
func godotCallVoidStringIntString(o Class, methodName string, arg0 string, arg1 int64, arg2 string) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)
	cArg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(cArg2))

	log.Println("  Calling godot_call_void_string_int_string...")
	C.godot_call_void_string_int_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_string_int_string returned")

}
func godotCallVoidRidIntArrayArrayInt(o Class, methodName string, arg0 *RID, arg1 int64, arg2 *Array, arg3 *Array, arg4 int64) {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)
	cArg2 := arg2.array
	cArg3 := arg3.array
	cArg4 := C.godot_int(arg4)

	log.Println("  Calling godot_call_void_rid_int_array_array_int...")
	C.godot_call_void_rid_int_array_array_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_void_rid_int_array_array_int returned")

}
func godotCallVoidRidObjectInt(o Class, methodName string, arg0 *RID, arg1 *Object, arg2 int64) {

	cArg0 := arg0.rid
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_void_rid_object_int...")
	C.godot_call_void_rid_object_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rid_object_int returned")

}
func godotCallVoidIntNodePath(o Class, methodName string, arg0 int64, arg1 *NodePath) {

	cArg0 := C.godot_int(arg0)
	cArg1 := arg1.nodePath

	log.Println("  Calling godot_call_void_int_nodepath...")
	C.godot_call_void_int_nodepath(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_int_nodepath returned")

}
func godotCallRect2(o Class, methodName string) *Rect2 {

	log.Println("  Calling godot_call_rect2...")
	cRet := C.godot_call_rect2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_rect2 returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_rect2)(cRet)
	return &Rect2{ret}

}
func godotCallPoolVector2ArrayIntFloat(o Class, methodName string, arg0 int64, arg1 float64) *PoolVector2Array {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_real(arg1)

	log.Println("  Calling godot_call_poolvector2array_int_float...")
	cRet := C.godot_call_poolvector2array_int_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_poolvector2array_int_float returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_pool_vector2_array)(cRet)
	return &PoolVector2Array{ret}

}
func godotCallBoolObjectStringVariantObjectStringFloatIntIntFloat(o Class, methodName string, arg0 *Object, arg1 string, arg2 *Variant, arg3 *Object, arg4 string, arg5 float64, arg6 int64, arg7 int64, arg8 float64) bool {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := arg2.variant
	cArg3 := unsafe.Pointer(arg3.owner)
	cArg4 := C.CString(arg4)
	defer C.free(unsafe.Pointer(cArg4))
	cArg5 := C.godot_real(arg5)
	cArg6 := C.godot_int(arg6)
	cArg7 := C.godot_int(arg7)
	cArg8 := C.godot_real(arg8)

	log.Println("  Calling godot_call_bool_object_string_variant_object_string_float_int_int_float...")
	cRet := C.godot_call_bool_object_string_variant_object_string_float_int_int_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5, cArg6, cArg7, cArg8,
	)
	log.Println("  ...godot_call_bool_object_string_variant_object_string_float_int_int_float returned")
	ret := cRet
	return bool(ret)

}
func godotCallBoolObjectFloatStringVariantVariantVariantVariantVariant(o Class, methodName string, arg0 *Object, arg1 float64, arg2 string, arg3 *Variant, arg4 *Variant, arg5 *Variant, arg6 *Variant, arg7 *Variant) bool {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.godot_real(arg1)
	cArg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(cArg2))
	cArg3 := arg3.variant
	cArg4 := arg4.variant
	cArg5 := arg5.variant
	cArg6 := arg6.variant
	cArg7 := arg7.variant

	log.Println("  Calling godot_call_bool_object_float_string_variant_variant_variant_variant_variant...")
	cRet := C.godot_call_bool_object_float_string_variant_variant_variant_variant_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5, cArg6, cArg7,
	)
	log.Println("  ...godot_call_bool_object_float_string_variant_variant_variant_variant_variant returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoidNodePathObjectInt(o Class, methodName string, arg0 *NodePath, arg1 *Object, arg2 int64) {

	cArg0 := arg0.nodePath
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_void_nodepath_object_int...")
	C.godot_call_void_nodepath_object_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_nodepath_object_int returned")

}
func godotCallVoidObjectString(o Class, methodName string, arg0 *Object, arg1 string) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))

	log.Println("  Calling godot_call_void_object_string...")
	C.godot_call_void_object_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_object_string returned")

}
func godotCallVoidVector3Float(o Class, methodName string, arg0 *Vector3, arg1 float64) {

	cArg0 := arg0.vector3
	cArg1 := C.godot_real(arg1)

	log.Println("  Calling godot_call_void_vector3_float...")
	C.godot_call_void_vector3_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_vector3_float returned")

}
func godotCallVoidRidVector2Vector2ColorFloatBool(o Class, methodName string, arg0 *RID, arg1 *Vector2, arg2 *Vector2, arg3 *Color, arg4 float64, arg5 bool) {

	cArg0 := arg0.rid
	cArg1 := arg1.vector2
	cArg2 := arg2.vector2
	cArg3 := arg3.color
	cArg4 := C.godot_real(arg4)
	cArg5 := C.godot_bool(arg5)

	log.Println("  Calling godot_call_void_rid_vector2_vector2_color_float_bool...")
	C.godot_call_void_rid_vector2_vector2_color_float_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5,
	)
	log.Println("  ...godot_call_void_rid_vector2_vector2_color_float_bool returned")

}
func godotCallPlane(o Class, methodName string) *Plane {

	log.Println("  Calling godot_call_plane...")
	cRet := C.godot_call_plane(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_plane returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_plane)(cRet)
	return &Plane{ret}

}
func godotCallVoidRidIntIntFloat(o Class, methodName string, arg0 *RID, arg1 int64, arg2 int64, arg3 float64) {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_int(arg2)
	cArg3 := C.godot_real(arg3)

	log.Println("  Calling godot_call_void_rid_int_int_float...")
	C.godot_call_void_rid_int_int_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_rid_int_int_float returned")

}
func godotCallBoolInt(o Class, methodName string, arg0 int64) bool {

	cArg0 := C.godot_int(arg0)

	log.Println("  Calling godot_call_bool_int...")
	cRet := C.godot_call_bool_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_bool_int returned")
	ret := cRet
	return bool(ret)

}
func godotCallPoolIntArrayIntFloatFloat(o Class, methodName string, arg0 int64, arg1 float64, arg2 float64) *PoolIntArray {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_real(arg1)
	cArg2 := C.godot_real(arg2)

	log.Println("  Calling godot_call_poolintarray_int_float_float...")
	cRet := C.godot_call_poolintarray_int_float_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_poolintarray_int_float_float returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_pool_int_array)(cRet)
	return &PoolIntArray{ret}

}
func godotCallVoidStringIntInt(o Class, methodName string, arg0 string, arg1 int64, arg2 int64) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_void_string_int_int...")
	C.godot_call_void_string_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_string_int_int returned")

}
func godotCallObjectStringInt(o Class, methodName string, arg0 string, arg1 int64) *Object {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_object_string_int...")
	cRet := C.godot_call_object_string_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_object_string_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_object)(cRet)
	return &Object{ret}

}
func godotCallBoolObjectNodePathVariantVariantFloatIntIntFloat(o Class, methodName string, arg0 *Object, arg1 *NodePath, arg2 *Variant, arg3 *Variant, arg4 float64, arg5 int64, arg6 int64, arg7 float64) bool {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := arg1.nodePath
	cArg2 := arg2.variant
	cArg3 := arg3.variant
	cArg4 := C.godot_real(arg4)
	cArg5 := C.godot_int(arg5)
	cArg6 := C.godot_int(arg6)
	cArg7 := C.godot_real(arg7)

	log.Println("  Calling godot_call_bool_object_nodepath_variant_variant_float_int_int_float...")
	cRet := C.godot_call_bool_object_nodepath_variant_variant_float_int_int_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5, cArg6, cArg7,
	)
	log.Println("  ...godot_call_bool_object_nodepath_variant_variant_float_int_int_float returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoidRidPoolVector2ArrayPoolColorArrayPoolVector2ArrayRidRidBool(o Class, methodName string, arg0 *RID, arg1 *PoolVector2Array, arg2 *PoolColorArray, arg3 *PoolVector2Array, arg4 *RID, arg5 *RID, arg6 bool) {

	cArg0 := arg0.rid
	cArg1 := arg1.array
	cArg2 := arg2.array
	cArg3 := arg3.array
	cArg4 := arg4.rid
	cArg5 := arg5.rid
	cArg6 := C.godot_bool(arg6)

	log.Println("  Calling godot_call_void_rid_poolvector2array_poolcolorarray_poolvector2array_rid_rid_bool...")
	C.godot_call_void_rid_poolvector2array_poolcolorarray_poolvector2array_rid_rid_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5, cArg6,
	)
	log.Println("  ...godot_call_void_rid_poolvector2array_poolcolorarray_poolvector2array_rid_rid_bool returned")

}
func godotCallArrayRid(o Class, methodName string, arg0 *RID) *Array {

	cArg0 := arg0.rid

	log.Println("  Calling godot_call_array_rid...")
	cRet := C.godot_call_array_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_array_rid returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_array)(cRet)
	return &Array{ret}

}
func godotCallIntStringIntBoolBool(o Class, methodName string, arg0 string, arg1 int64, arg2 bool, arg3 bool) int64 {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_bool(arg2)
	cArg3 := C.godot_bool(arg3)

	log.Println("  Calling godot_call_int_string_int_bool_bool...")
	cRet := C.godot_call_int_string_int_bool_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_int_string_int_bool_bool returned")
	ret := cRet
	return int64(ret)

}
func godotCallVoidStringDictionary(o Class, methodName string, arg0 string, arg1 *Dictionary) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := arg1.dictionary

	log.Println("  Calling godot_call_void_string_dictionary...")
	C.godot_call_void_string_dictionary(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_string_dictionary returned")

}
func godotCallVoidRidRect2RidBoolColorBoolRid(o Class, methodName string, arg0 *RID, arg1 *Rect2, arg2 *RID, arg3 bool, arg4 *Color, arg5 bool, arg6 *RID) {

	cArg0 := arg0.rid
	cArg1 := arg1.rect2
	cArg2 := arg2.rid
	cArg3 := C.godot_bool(arg3)
	cArg4 := arg4.color
	cArg5 := C.godot_bool(arg5)
	cArg6 := arg6.rid

	log.Println("  Calling godot_call_void_rid_rect2_rid_bool_color_bool_rid...")
	C.godot_call_void_rid_rect2_rid_bool_color_bool_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5, cArg6,
	)
	log.Println("  ...godot_call_void_rid_rect2_rid_bool_color_bool_rid returned")

}
func godotCallVoidPoolVector3Array(o Class, methodName string, arg0 *PoolVector3Array) {

	cArg0 := arg0.array

	log.Println("  Calling godot_call_void_poolvector3array...")
	C.godot_call_void_poolvector3array(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_poolvector3array returned")

}
func godotCallIntVector2Object(o Class, methodName string, arg0 *Vector2, arg1 *Object) int64 {

	cArg0 := arg0.vector2
	cArg1 := unsafe.Pointer(arg1.owner)

	log.Println("  Calling godot_call_int_vector2_object...")
	cRet := C.godot_call_int_vector2_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_int_vector2_object returned")
	ret := cRet
	return int64(ret)

}
func godotCallIntStringObject(o Class, methodName string, arg0 string, arg1 *Object) int64 {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := unsafe.Pointer(arg1.owner)

	log.Println("  Calling godot_call_int_string_object...")
	cRet := C.godot_call_int_string_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_int_string_object returned")
	ret := cRet
	return int64(ret)

}
func godotCallVoidStringVector2(o Class, methodName string, arg0 string, arg1 *Vector2) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := arg1.vector2

	log.Println("  Calling godot_call_void_string_vector2...")
	C.godot_call_void_string_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_string_vector2 returned")

}
func godotCallVoidObjectObjectVector3Vector3Int(o Class, methodName string, arg0 *Object, arg1 *Object, arg2 *Vector3, arg3 *Vector3, arg4 int64) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := arg2.vector3
	cArg3 := arg3.vector3
	cArg4 := C.godot_int(arg4)

	log.Println("  Calling godot_call_void_object_object_vector3_vector3_int...")
	C.godot_call_void_object_object_vector3_vector3_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_void_object_object_vector3_vector3_int returned")

}
func godotCallIntPoolByteArray(o Class, methodName string, arg0 *PoolByteArray) int64 {

	cArg0 := arg0.array

	log.Println("  Calling godot_call_int_poolbytearray...")
	cRet := C.godot_call_int_poolbytearray(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_int_poolbytearray returned")
	ret := cRet
	return int64(ret)

}
func godotCallVoidIntIntColor(o Class, methodName string, arg0 int64, arg1 int64, arg2 *Color) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := arg2.color

	log.Println("  Calling godot_call_void_int_int_color...")
	C.godot_call_void_int_int_color(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_int_int_color returned")

}
func godotCallBoolObjectString(o Class, methodName string, arg0 *Object, arg1 string) bool {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))

	log.Println("  Calling godot_call_bool_object_string...")
	cRet := C.godot_call_bool_object_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_bool_object_string returned")
	ret := cRet
	return bool(ret)

}
func godotCallVector3Vector3Vector3FloatIntFloat(o Class, methodName string, arg0 *Vector3, arg1 *Vector3, arg2 float64, arg3 int64, arg4 float64) *Vector3 {

	cArg0 := arg0.vector3
	cArg1 := arg1.vector3
	cArg2 := C.godot_real(arg2)
	cArg3 := C.godot_int(arg3)
	cArg4 := C.godot_real(arg4)

	log.Println("  Calling godot_call_vector3_vector3_vector3_float_int_float...")
	cRet := C.godot_call_vector3_vector3_vector3_float_int_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_vector3_vector3_vector3_float_int_float returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector3)(cRet)
	return &Vector3{ret}

}
func godotCallVoidRidVector2FloatColor(o Class, methodName string, arg0 *RID, arg1 *Vector2, arg2 float64, arg3 *Color) {

	cArg0 := arg0.rid
	cArg1 := arg1.vector2
	cArg2 := C.godot_real(arg2)
	cArg3 := arg3.color

	log.Println("  Calling godot_call_void_rid_vector2_float_color...")
	C.godot_call_void_rid_vector2_float_color(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_rid_vector2_float_color returned")

}
func godotCallArrayIntFloat(o Class, methodName string, arg0 int64, arg1 float64) *Array {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_real(arg1)

	log.Println("  Calling godot_call_array_int_float...")
	cRet := C.godot_call_array_int_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_array_int_float returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_array)(cRet)
	return &Array{ret}

}
func godotCallBoolObject(o Class, methodName string, arg0 *Object) bool {

	cArg0 := unsafe.Pointer(arg0.owner)

	log.Println("  Calling godot_call_bool_object...")
	cRet := C.godot_call_bool_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_bool_object returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoidIntPoolRealArray(o Class, methodName string, arg0 int64, arg1 *PoolRealArray) {

	cArg0 := C.godot_int(arg0)
	cArg1 := arg1.array

	log.Println("  Calling godot_call_void_int_poolrealarray...")
	C.godot_call_void_int_poolrealarray(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_int_poolrealarray returned")

}
func godotCallVoidRidRidTransform2D(o Class, methodName string, arg0 *RID, arg1 *RID, arg2 *Transform2D) {

	cArg0 := arg0.rid
	cArg1 := arg1.rid
	cArg2 := arg2.transform2d

	log.Println("  Calling godot_call_void_rid_rid_transform2d...")
	C.godot_call_void_rid_rid_transform2d(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rid_rid_transform2d returned")

}
func godotCallVoidRidBool(o Class, methodName string, arg0 *RID, arg1 bool) {

	cArg0 := arg0.rid
	cArg1 := C.godot_bool(arg1)

	log.Println("  Calling godot_call_void_rid_bool...")
	C.godot_call_void_rid_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_rid_bool returned")

}
func godotCallVoidStringPoolByteArrayBool(o Class, methodName string, arg0 string, arg1 *PoolByteArray, arg2 bool) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := arg1.array
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_void_string_poolbytearray_bool...")
	C.godot_call_void_string_poolbytearray_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_string_poolbytearray_bool returned")

}
func godotCallVoidStringPoolStringArray(o Class, methodName string, arg0 string, arg1 *PoolStringArray) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := arg1.array

	log.Println("  Calling godot_call_void_string_poolstringarray...")
	C.godot_call_void_string_poolstringarray(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_string_poolstringarray returned")

}
func godotCallVoidIntVariantBool(o Class, methodName string, arg0 int64, arg1 *Variant, arg2 bool) {

	cArg0 := C.godot_int(arg0)
	cArg1 := arg1.variant
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_void_int_variant_bool...")
	C.godot_call_void_int_variant_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_int_variant_bool returned")

}
func godotCallVoidBoolFloat(o Class, methodName string, arg0 bool, arg1 float64) {

	cArg0 := C.godot_bool(arg0)
	cArg1 := C.godot_real(arg1)

	log.Println("  Calling godot_call_void_bool_float...")
	C.godot_call_void_bool_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_bool_float returned")

}
func godotCallIntIntIntFloat(o Class, methodName string, arg0 int64, arg1 int64, arg2 float64) int64 {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_real(arg2)

	log.Println("  Calling godot_call_int_int_int_float...")
	cRet := C.godot_call_int_int_int_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_int_int_int_float returned")
	ret := cRet
	return int64(ret)

}
func godotCallVoidObjectObjectBool(o Class, methodName string, arg0 *Object, arg1 *Object, arg2 bool) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_void_object_object_bool...")
	C.godot_call_void_object_object_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_object_object_bool returned")

}
func godotCallFloatRidIntInt(o Class, methodName string, arg0 *RID, arg1 int64, arg2 int64) float64 {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_float_rid_int_int...")
	cRet := C.godot_call_float_rid_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_float_rid_int_int returned")
	ret := cRet
	return float64(ret)

}
func godotCallVoidObjectStringPoolStringArray(o Class, methodName string, arg0 *Object, arg1 string, arg2 *PoolStringArray) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := arg2.array

	log.Println("  Calling godot_call_void_object_string_poolstringarray...")
	C.godot_call_void_object_string_poolstringarray(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_object_string_poolstringarray returned")

}
func godotCallVoidIntColorBool(o Class, methodName string, arg0 int64, arg1 *Color, arg2 bool) {

	cArg0 := C.godot_int(arg0)
	cArg1 := arg1.color
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_void_int_color_bool...")
	C.godot_call_void_int_color_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_int_color_bool returned")

}
func godotCallFloatInt(o Class, methodName string, arg0 int64) float64 {

	cArg0 := C.godot_int(arg0)

	log.Println("  Calling godot_call_float_int...")
	cRet := C.godot_call_float_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_float_int returned")
	ret := cRet
	return float64(ret)

}
func godotCallVoidIntObjectVector2(o Class, methodName string, arg0 int64, arg1 *Object, arg2 *Vector2) {

	cArg0 := C.godot_int(arg0)
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := arg2.vector2

	log.Println("  Calling godot_call_void_int_object_vector2...")
	C.godot_call_void_int_object_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_int_object_vector2 returned")

}
func godotCallPoolVector3ArrayVector3Vector3Bool(o Class, methodName string, arg0 *Vector3, arg1 *Vector3, arg2 bool) *PoolVector3Array {

	cArg0 := arg0.vector3
	cArg1 := arg1.vector3
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_poolvector3array_vector3_vector3_bool...")
	cRet := C.godot_call_poolvector3array_vector3_vector3_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_poolvector3array_vector3_vector3_bool returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_pool_vector3_array)(cRet)
	return &PoolVector3Array{ret}

}
func godotCallIntBool(o Class, methodName string, arg0 bool) int64 {

	cArg0 := C.godot_bool(arg0)

	log.Println("  Calling godot_call_int_bool...")
	cRet := C.godot_call_int_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_int_bool returned")
	ret := cRet
	return int64(ret)

}
func godotCallBoolVector2FloatObject(o Class, methodName string, arg0 *Vector2, arg1 float64, arg2 *Object) bool {

	cArg0 := arg0.vector2
	cArg1 := C.godot_real(arg1)
	cArg2 := unsafe.Pointer(arg2.owner)

	log.Println("  Calling godot_call_bool_vector2_float_object...")
	cRet := C.godot_call_bool_vector2_float_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_bool_vector2_float_object returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoid(o Class, methodName string) {

	log.Println("  Calling godot_call_void...")
	C.godot_call_void(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_void returned")

}
func godotCallVoidIntVector3Float(o Class, methodName string, arg0 int64, arg1 *Vector3, arg2 float64) {

	cArg0 := C.godot_int(arg0)
	cArg1 := arg1.vector3
	cArg2 := C.godot_real(arg2)

	log.Println("  Calling godot_call_void_int_vector3_float...")
	C.godot_call_void_int_vector3_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_int_vector3_float returned")

}
func godotCallArrayInt(o Class, methodName string, arg0 int64) *Array {

	cArg0 := C.godot_int(arg0)

	log.Println("  Calling godot_call_array_int...")
	cRet := C.godot_call_array_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_array_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_array)(cRet)
	return &Array{ret}

}
func godotCallVoidVector3Vector3Vector3(o Class, methodName string, arg0 *Vector3, arg1 *Vector3, arg2 *Vector3) {

	cArg0 := arg0.vector3
	cArg1 := arg1.vector3
	cArg2 := arg2.vector3

	log.Println("  Calling godot_call_void_vector3_vector3_vector3...")
	C.godot_call_void_vector3_vector3_vector3(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_vector3_vector3_vector3 returned")

}
func godotCallVoidRidIntRid(o Class, methodName string, arg0 *RID, arg1 int64, arg2 *RID) {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)
	cArg2 := arg2.rid

	log.Println("  Calling godot_call_void_rid_int_rid...")
	C.godot_call_void_rid_int_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rid_int_rid returned")

}
func godotCallRidVector2Vector2Vector2RidRid(o Class, methodName string, arg0 *Vector2, arg1 *Vector2, arg2 *Vector2, arg3 *RID, arg4 *RID) *RID {

	cArg0 := arg0.vector2
	cArg1 := arg1.vector2
	cArg2 := arg2.vector2
	cArg3 := arg3.rid
	cArg4 := arg4.rid

	log.Println("  Calling godot_call_rid_vector2_vector2_vector2_rid_rid...")
	cRet := C.godot_call_rid_vector2_vector2_vector2_rid_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_rid_vector2_vector2_vector2_rid_rid returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_rid)(cRet)
	return &RID{ret}

}
func godotCallVoidIntStringStringVariant(o Class, methodName string, arg0 int64, arg1 string, arg2 string, arg3 *Variant) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(cArg2))
	cArg3 := arg3.variant

	log.Println("  Calling godot_call_void_int_string_string_variant...")
	C.godot_call_void_int_string_string_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_int_string_string_variant returned")

}
func godotCallVoidObjectVector2ColorObject(o Class, methodName string, arg0 *Object, arg1 *Vector2, arg2 *Color, arg3 *Object) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := arg1.vector2
	cArg2 := arg2.color
	cArg3 := unsafe.Pointer(arg3.owner)

	log.Println("  Calling godot_call_void_object_vector2_color_object...")
	C.godot_call_void_object_vector2_color_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_object_vector2_color_object returned")

}
func godotCallArrayNodePath(o Class, methodName string, arg0 *NodePath) *Array {

	cArg0 := arg0.nodePath

	log.Println("  Calling godot_call_array_nodepath...")
	cRet := C.godot_call_array_nodepath(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_array_nodepath returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_array)(cRet)
	return &Array{ret}

}
func godotCallFloatRidInt(o Class, methodName string, arg0 *RID, arg1 int64) float64 {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_float_rid_int...")
	cRet := C.godot_call_float_rid_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_float_rid_int returned")
	ret := cRet
	return float64(ret)

}
func godotCallIntRidInt(o Class, methodName string, arg0 *RID, arg1 int64) int64 {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_int_rid_int...")
	cRet := C.godot_call_int_rid_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_int_rid_int returned")
	ret := cRet
	return int64(ret)

}
func godotCallVoidArray(o Class, methodName string, arg0 *Array) {

	cArg0 := arg0.array

	log.Println("  Calling godot_call_void_array...")
	C.godot_call_void_array(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_array returned")

}
func godotCallVoidIntIntRect2Vector2Float(o Class, methodName string, arg0 int64, arg1 int64, arg2 *Rect2, arg3 *Vector2, arg4 float64) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := arg2.rect2
	cArg3 := arg3.vector2
	cArg4 := C.godot_real(arg4)

	log.Println("  Calling godot_call_void_int_int_rect2_vector2_float...")
	C.godot_call_void_int_int_rect2_vector2_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_void_int_int_rect2_vector2_float returned")

}
func godotCallVoidRidRidRidRidIntInt(o Class, methodName string, arg0 *RID, arg1 *RID, arg2 *RID, arg3 *RID, arg4 int64, arg5 int64) {

	cArg0 := arg0.rid
	cArg1 := arg1.rid
	cArg2 := arg2.rid
	cArg3 := arg3.rid
	cArg4 := C.godot_int(arg4)
	cArg5 := C.godot_int(arg5)

	log.Println("  Calling godot_call_void_rid_rid_rid_rid_int_int...")
	C.godot_call_void_rid_rid_rid_rid_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5,
	)
	log.Println("  ...godot_call_void_rid_rid_rid_rid_int_int returned")

}
func godotCallBoolVector3(o Class, methodName string, arg0 *Vector3) bool {

	cArg0 := arg0.vector3

	log.Println("  Calling godot_call_bool_vector3...")
	cRet := C.godot_call_bool_vector3(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_bool_vector3 returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoidObjectRect2(o Class, methodName string, arg0 *Object, arg1 *Rect2) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := arg1.rect2

	log.Println("  Calling godot_call_void_object_rect2...")
	C.godot_call_void_object_rect2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_object_rect2 returned")

}
func godotCallIntIntIntInt(o Class, methodName string, arg0 int64, arg1 int64, arg2 int64) int64 {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_int_int_int_int...")
	cRet := C.godot_call_int_int_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_int_int_int_int returned")
	ret := cRet
	return int64(ret)

}
func godotCallVoidObjectRect2Vector2(o Class, methodName string, arg0 *Object, arg1 *Rect2, arg2 *Vector2) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := arg1.rect2
	cArg2 := arg2.vector2

	log.Println("  Calling godot_call_void_object_rect2_vector2...")
	C.godot_call_void_object_rect2_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_object_rect2_vector2 returned")

}
func godotCallVoidIntStringVariant(o Class, methodName string, arg0 int64, arg1 string, arg2 *Variant) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := arg2.variant

	log.Println("  Calling godot_call_void_int_string_variant...")
	C.godot_call_void_int_string_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_int_string_variant returned")

}
func godotCallRidRid(o Class, methodName string, arg0 *RID) *RID {

	cArg0 := arg0.rid

	log.Println("  Calling godot_call_rid_rid...")
	cRet := C.godot_call_rid_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_rid_rid returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_rid)(cRet)
	return &RID{ret}

}
func godotCallArrayStringIntInt(o Class, methodName string, arg0 string, arg1 int64, arg2 int64) *Array {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_array_string_int_int...")
	cRet := C.godot_call_array_string_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_array_string_int_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_array)(cRet)
	return &Array{ret}

}
func godotCallBoolStringObjectString(o Class, methodName string, arg0 string, arg1 *Object, arg2 string) bool {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(cArg2))

	log.Println("  Calling godot_call_bool_string_object_string...")
	cRet := C.godot_call_bool_string_object_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_bool_string_object_string returned")
	ret := cRet
	return bool(ret)

}
func godotCallRidIntBool(o Class, methodName string, arg0 int64, arg1 bool) *RID {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_bool(arg1)

	log.Println("  Calling godot_call_rid_int_bool...")
	cRet := C.godot_call_rid_int_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_rid_int_bool returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_rid)(cRet)
	return &RID{ret}

}
func godotCallVoidRidStringVariant(o Class, methodName string, arg0 *RID, arg1 string, arg2 *Variant) {

	cArg0 := arg0.rid
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := arg2.variant

	log.Println("  Calling godot_call_void_rid_string_variant...")
	C.godot_call_void_rid_string_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rid_string_variant returned")

}
func godotCallVector2String(o Class, methodName string, arg0 string) *Vector2 {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))

	log.Println("  Calling godot_call_vector2_string...")
	cRet := C.godot_call_vector2_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_vector2_string returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector2)(cRet)
	return &Vector2{ret}

}
func godotCallIntIntStringPoolStringArrayString(o Class, methodName string, arg0 int64, arg1 string, arg2 *PoolStringArray, arg3 string) int64 {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := arg2.array
	cArg3 := C.CString(arg3)
	defer C.free(unsafe.Pointer(cArg3))

	log.Println("  Calling godot_call_int_int_string_poolstringarray_string...")
	cRet := C.godot_call_int_int_string_poolstringarray_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_int_int_string_poolstringarray_string returned")
	ret := cRet
	return int64(ret)

}
func godotCallVector2IntIntObjectVector2(o Class, methodName string, arg0 int64, arg1 int64, arg2 *Object, arg3 *Vector2) *Vector2 {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := unsafe.Pointer(arg2.owner)
	cArg3 := arg3.vector2

	log.Println("  Calling godot_call_vector2_int_int_object_vector2...")
	cRet := C.godot_call_vector2_int_int_object_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_vector2_int_int_object_vector2 returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector2)(cRet)
	return &Vector2{ret}

}
func godotCallVoidIntIntInt(o Class, methodName string, arg0 int64, arg1 int64, arg2 int64) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_void_int_int_int...")
	C.godot_call_void_int_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_int_int_int returned")

}
func godotCallVoidPlane(o Class, methodName string, arg0 *Plane) {

	cArg0 := arg0.plane

	log.Println("  Calling godot_call_void_plane...")
	C.godot_call_void_plane(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_plane returned")

}
func godotCallBoolVector2VariantObject(o Class, methodName string, arg0 *Vector2, arg1 *Variant, arg2 *Object) bool {

	cArg0 := arg0.vector2
	cArg1 := arg1.variant
	cArg2 := unsafe.Pointer(arg2.owner)

	log.Println("  Calling godot_call_bool_vector2_variant_object...")
	cRet := C.godot_call_bool_vector2_variant_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_bool_vector2_variant_object returned")
	ret := cRet
	return bool(ret)

}
func godotCallObjectInt(o Class, methodName string, arg0 int64) *Object {

	cArg0 := C.godot_int(arg0)

	log.Println("  Calling godot_call_object_int...")
	cRet := C.godot_call_object_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_object_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_object)(cRet)
	return &Object{ret}

}
func godotCallIntNodePath(o Class, methodName string, arg0 *NodePath) int64 {

	cArg0 := arg0.nodePath

	log.Println("  Calling godot_call_int_nodepath...")
	cRet := C.godot_call_int_nodepath(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_int_nodepath returned")
	ret := cRet
	return int64(ret)

}
func godotCallPoolVector2Array(o Class, methodName string) *PoolVector2Array {

	log.Println("  Calling godot_call_poolvector2array...")
	cRet := C.godot_call_poolvector2array(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_poolvector2array returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_pool_vector2_array)(cRet)
	return &PoolVector2Array{ret}

}
func godotCallBoolStringString(o Class, methodName string, arg0 string, arg1 string) bool {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))

	log.Println("  Calling godot_call_bool_string_string...")
	cRet := C.godot_call_bool_string_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_bool_string_string returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoidStringStringObjectObject(o Class, methodName string, arg0 string, arg1 string, arg2 *Object, arg3 *Object) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := unsafe.Pointer(arg2.owner)
	cArg3 := unsafe.Pointer(arg3.owner)

	log.Println("  Calling godot_call_void_string_string_object_object...")
	C.godot_call_void_string_string_object_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_string_string_object_object returned")

}
func godotCallVector2StringInt(o Class, methodName string, arg0 string, arg1 int64) *Vector2 {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_vector2_string_int...")
	cRet := C.godot_call_vector2_string_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_vector2_string_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector2)(cRet)
	return &Vector2{ret}

}
func godotCallIntStringStringDictionaryArrayArray(o Class, methodName string, arg0 string, arg1 string, arg2 *Dictionary, arg3 *Array, arg4 *Array) int64 {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := arg2.dictionary
	cArg3 := arg3.array
	cArg4 := arg4.array

	log.Println("  Calling godot_call_int_string_string_dictionary_array_array...")
	cRet := C.godot_call_int_string_string_dictionary_array_array(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_int_string_string_dictionary_array_array returned")
	ret := cRet
	return int64(ret)

}
func godotCallVoidStringObjectStringVariant(o Class, methodName string, arg0 string, arg1 *Object, arg2 string, arg3 *Variant) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(cArg2))
	cArg3 := arg3.variant

	log.Println("  Calling godot_call_void_string_object_string_variant...")
	C.godot_call_void_string_object_string_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_string_object_string_variant returned")

}
func godotCallPoolVector2ArrayVector2Vector2Bool(o Class, methodName string, arg0 *Vector2, arg1 *Vector2, arg2 bool) *PoolVector2Array {

	cArg0 := arg0.vector2
	cArg1 := arg1.vector2
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_poolvector2array_vector2_vector2_bool...")
	cRet := C.godot_call_poolvector2array_vector2_vector2_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_poolvector2array_vector2_vector2_bool returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_pool_vector2_array)(cRet)
	return &PoolVector2Array{ret}

}
func godotCallRidRidString(o Class, methodName string, arg0 *RID, arg1 string) *RID {

	cArg0 := arg0.rid
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))

	log.Println("  Calling godot_call_rid_rid_string...")
	cRet := C.godot_call_rid_rid_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_rid_rid_string returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_rid)(cRet)
	return &RID{ret}

}
func godotCallIntIntIntIntInt(o Class, methodName string, arg0 int64, arg1 int64, arg2 int64, arg3 int64) int64 {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_int(arg2)
	cArg3 := C.godot_int(arg3)

	log.Println("  Calling godot_call_int_int_int_int_int...")
	cRet := C.godot_call_int_int_int_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_int_int_int_int_int returned")
	ret := cRet
	return int64(ret)

}
func godotCallArrayObjectVector3(o Class, methodName string, arg0 *Object, arg1 *Vector3) *Array {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := arg1.vector3

	log.Println("  Calling godot_call_array_object_vector3...")
	cRet := C.godot_call_array_object_vector3(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_array_object_vector3 returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_array)(cRet)
	return &Array{ret}

}
func godotCallBoolTransform2DObjectTransform2D(o Class, methodName string, arg0 *Transform2D, arg1 *Object, arg2 *Transform2D) bool {

	cArg0 := arg0.transform2d
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := arg2.transform2d

	log.Println("  Calling godot_call_bool_transform2d_object_transform2d...")
	cRet := C.godot_call_bool_transform2d_object_transform2d(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_bool_transform2d_object_transform2d returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoidPoolVector2ArrayColorFloatBool(o Class, methodName string, arg0 *PoolVector2Array, arg1 *Color, arg2 float64, arg3 bool) {

	cArg0 := arg0.array
	cArg1 := arg1.color
	cArg2 := C.godot_real(arg2)
	cArg3 := C.godot_bool(arg3)

	log.Println("  Calling godot_call_void_poolvector2array_color_float_bool...")
	C.godot_call_void_poolvector2array_color_float_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_poolvector2array_color_float_bool returned")

}
func godotCallVoidIntArray(o Class, methodName string, arg0 int64, arg1 *Array) {

	cArg0 := C.godot_int(arg0)
	cArg1 := arg1.array

	log.Println("  Calling godot_call_void_int_array...")
	C.godot_call_void_int_array(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_int_array returned")

}
func godotCallNodePathIntBool(o Class, methodName string, arg0 int64, arg1 bool) *NodePath {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_bool(arg1)

	log.Println("  Calling godot_call_nodepath_int_bool...")
	cRet := C.godot_call_nodepath_int_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_nodepath_int_bool returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_node_path)(cRet)
	return &NodePath{ret}

}
func godotCallVariantVector2Object(o Class, methodName string, arg0 *Vector2, arg1 *Object) *Variant {

	cArg0 := arg0.vector2
	cArg1 := unsafe.Pointer(arg1.owner)

	log.Println("  Calling godot_call_variant_vector2_object...")
	cRet := C.godot_call_variant_vector2_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_variant_vector2_object returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_variant)(cRet)
	return &Variant{ret}

}
func godotCallVoidStringNodePathBool(o Class, methodName string, arg0 string, arg1 *NodePath, arg2 bool) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := arg1.nodePath
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_void_string_nodepath_bool...")
	C.godot_call_void_string_nodepath_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_string_nodepath_bool returned")

}
func godotCallTransform2DInt(o Class, methodName string, arg0 int64) *Transform2D {

	cArg0 := C.godot_int(arg0)

	log.Println("  Calling godot_call_transform2d_int...")
	cRet := C.godot_call_transform2d_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_transform2d_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_transform2d)(cRet)
	return &Transform2D{ret}

}
func godotCallVoidIntIntBoolIntPoolByteArray(o Class, methodName string, arg0 int64, arg1 int64, arg2 bool, arg3 int64, arg4 *PoolByteArray) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_bool(arg2)
	cArg3 := C.godot_int(arg3)
	cArg4 := arg4.array

	log.Println("  Calling godot_call_void_int_int_bool_int_poolbytearray...")
	C.godot_call_void_int_int_bool_int_poolbytearray(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_void_int_int_bool_int_poolbytearray returned")

}
func godotCallVoidPoolVector2ArrayPoolIntArray(o Class, methodName string, arg0 *PoolVector2Array, arg1 *PoolIntArray) {

	cArg0 := arg0.array
	cArg1 := arg1.array

	log.Println("  Calling godot_call_void_poolvector2array_poolintarray...")
	C.godot_call_void_poolvector2array_poolintarray(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_poolvector2array_poolintarray returned")

}
func godotCallStringRid(o Class, methodName string, arg0 *RID) string {

	cArg0 := arg0.rid

	log.Println("  Calling godot_call_string_rid...")
	cRet := C.godot_call_string_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_string_rid returned")
	return C.GoString(cRet)

}
func godotCallDictionaryVector2Vector2ArrayInt(o Class, methodName string, arg0 *Vector2, arg1 *Vector2, arg2 *Array, arg3 int64) *Dictionary {

	cArg0 := arg0.vector2
	cArg1 := arg1.vector2
	cArg2 := arg2.array
	cArg3 := C.godot_int(arg3)

	log.Println("  Calling godot_call_dictionary_vector2_vector2_array_int...")
	cRet := C.godot_call_dictionary_vector2_vector2_array_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_dictionary_vector2_vector2_array_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_dictionary)(cRet)
	return &Dictionary{ret}

}
func godotCallVoidIntIntVariant(o Class, methodName string, arg0 int64, arg1 int64, arg2 *Variant) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := arg2.variant

	log.Println("  Calling godot_call_void_int_int_variant...")
	C.godot_call_void_int_int_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_int_int_variant returned")

}
func godotCallVector2Vector3(o Class, methodName string, arg0 *Vector3) *Vector2 {

	cArg0 := arg0.vector3

	log.Println("  Calling godot_call_vector2_vector3...")
	cRet := C.godot_call_vector2_vector3(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_vector2_vector3 returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector2)(cRet)
	return &Vector2{ret}

}
func godotCallVector3Float(o Class, methodName string, arg0 float64) *Vector3 {

	cArg0 := C.godot_real(arg0)

	log.Println("  Calling godot_call_vector3_float...")
	cRet := C.godot_call_vector3_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_vector3_float returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector3)(cRet)
	return &Vector3{ret}

}
func godotCallVoidIntIntPoolStringArrayPoolByteArray(o Class, methodName string, arg0 int64, arg1 int64, arg2 *PoolStringArray, arg3 *PoolByteArray) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := arg2.array
	cArg3 := arg3.array

	log.Println("  Calling godot_call_void_int_int_poolstringarray_poolbytearray...")
	C.godot_call_void_int_int_poolstringarray_poolbytearray(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_int_int_poolstringarray_poolbytearray returned")

}
func godotCallIntStringInt(o Class, methodName string, arg0 string, arg1 int64) int64 {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_int_string_int...")
	cRet := C.godot_call_int_string_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_int_string_int returned")
	ret := cRet
	return int64(ret)

}
func godotCallVoidObjectObjectRect2Vector2(o Class, methodName string, arg0 *Object, arg1 *Object, arg2 *Rect2, arg3 *Vector2) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := arg2.rect2
	cArg3 := arg3.vector2

	log.Println("  Calling godot_call_void_object_object_rect2_vector2...")
	C.godot_call_void_object_object_rect2_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_object_object_rect2_vector2 returned")

}
func godotCallVoidFloat(o Class, methodName string, arg0 float64) {

	cArg0 := C.godot_real(arg0)

	log.Println("  Calling godot_call_void_float...")
	C.godot_call_void_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_float returned")

}
func godotCallFloatStringString(o Class, methodName string, arg0 string, arg1 string) float64 {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))

	log.Println("  Calling godot_call_float_string_string...")
	cRet := C.godot_call_float_string_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_float_string_string returned")
	ret := cRet
	return float64(ret)

}
func godotCallBoolStringObject(o Class, methodName string, arg0 string, arg1 *Object) bool {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := unsafe.Pointer(arg1.owner)

	log.Println("  Calling godot_call_bool_string_object...")
	cRet := C.godot_call_bool_string_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_bool_string_object returned")
	ret := cRet
	return bool(ret)

}
func godotCallDictionaryVector3Vector3ArrayInt(o Class, methodName string, arg0 *Vector3, arg1 *Vector3, arg2 *Array, arg3 int64) *Dictionary {

	cArg0 := arg0.vector3
	cArg1 := arg1.vector3
	cArg2 := arg2.array
	cArg3 := C.godot_int(arg3)

	log.Println("  Calling godot_call_dictionary_vector3_vector3_array_int...")
	cRet := C.godot_call_dictionary_vector3_vector3_array_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_dictionary_vector3_vector3_array_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_dictionary)(cRet)
	return &Dictionary{ret}

}
func godotCallVoidIntIntTransform2D(o Class, methodName string, arg0 int64, arg1 int64, arg2 *Transform2D) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_int(arg1)
	cArg2 := arg2.transform2d

	log.Println("  Calling godot_call_void_int_int_transform2d...")
	C.godot_call_void_int_int_transform2d(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_int_int_transform2d returned")

}
func godotCallBoolObjectNodePathObjectNodePathVariantFloatIntIntFloat(o Class, methodName string, arg0 *Object, arg1 *NodePath, arg2 *Object, arg3 *NodePath, arg4 *Variant, arg5 float64, arg6 int64, arg7 int64, arg8 float64) bool {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := arg1.nodePath
	cArg2 := unsafe.Pointer(arg2.owner)
	cArg3 := arg3.nodePath
	cArg4 := arg4.variant
	cArg5 := C.godot_real(arg5)
	cArg6 := C.godot_int(arg6)
	cArg7 := C.godot_int(arg7)
	cArg8 := C.godot_real(arg8)

	log.Println("  Calling godot_call_bool_object_nodepath_object_nodepath_variant_float_int_int_float...")
	cRet := C.godot_call_bool_object_nodepath_object_nodepath_variant_float_int_int_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5, cArg6, cArg7, cArg8,
	)
	log.Println("  ...godot_call_bool_object_nodepath_object_nodepath_variant_float_int_int_float returned")
	ret := cRet
	return bool(ret)

}
func godotCallVoidPoolByteArray(o Class, methodName string, arg0 *PoolByteArray) {

	cArg0 := arg0.array

	log.Println("  Calling godot_call_void_poolbytearray...")
	C.godot_call_void_poolbytearray(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_poolbytearray returned")

}
func godotCallVoidObjectObjectInt(o Class, methodName string, arg0 *Object, arg1 *Object, arg2 int64) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_void_object_object_int...")
	C.godot_call_void_object_object_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_object_object_int returned")

}
func godotCallVoidVector2Variant(o Class, methodName string, arg0 *Vector2, arg1 *Variant) {

	cArg0 := arg0.vector2
	cArg1 := arg1.variant

	log.Println("  Calling godot_call_void_vector2_variant...")
	C.godot_call_void_vector2_variant(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_vector2_variant returned")

}
func godotCallVoidRidTransform(o Class, methodName string, arg0 *RID, arg1 *Transform) {

	cArg0 := arg0.rid
	cArg1 := arg1.transform

	log.Println("  Calling godot_call_void_rid_transform...")
	C.godot_call_void_rid_transform(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_rid_transform returned")

}
func godotCallVoidRidRect2BoolColorBoolObject(o Class, methodName string, arg0 *RID, arg1 *Rect2, arg2 bool, arg3 *Color, arg4 bool, arg5 *Object) {

	cArg0 := arg0.rid
	cArg1 := arg1.rect2
	cArg2 := C.godot_bool(arg2)
	cArg3 := arg3.color
	cArg4 := C.godot_bool(arg4)
	cArg5 := unsafe.Pointer(arg5.owner)

	log.Println("  Calling godot_call_void_rid_rect2_bool_color_bool_object...")
	C.godot_call_void_rid_rect2_bool_color_bool_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5,
	)
	log.Println("  ...godot_call_void_rid_rect2_bool_color_bool_object returned")

}
func godotCallVoidRidString(o Class, methodName string, arg0 *RID, arg1 string) {

	cArg0 := arg0.rid
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))

	log.Println("  Calling godot_call_void_rid_string...")
	C.godot_call_void_rid_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_rid_string returned")

}
func godotCallBoolBool(o Class, methodName string, arg0 bool) bool {

	cArg0 := C.godot_bool(arg0)

	log.Println("  Calling godot_call_bool_bool...")
	cRet := C.godot_call_bool_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_bool_bool returned")
	ret := cRet
	return bool(ret)

}
func godotCallIntIntStringPoolStringArrayPoolByteArray(o Class, methodName string, arg0 int64, arg1 string, arg2 *PoolStringArray, arg3 *PoolByteArray) int64 {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(cArg1))
	cArg2 := arg2.array
	cArg3 := arg3.array

	log.Println("  Calling godot_call_int_int_string_poolstringarray_poolbytearray...")
	cRet := C.godot_call_int_int_string_poolstringarray_poolbytearray(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_int_int_string_poolstringarray_poolbytearray returned")
	ret := cRet
	return int64(ret)

}
func godotCallObjectVector3(o Class, methodName string, arg0 *Vector3) *Object {

	cArg0 := arg0.vector3

	log.Println("  Calling godot_call_object_vector3...")
	cRet := C.godot_call_object_vector3(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_object_vector3 returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_object)(cRet)
	return &Object{ret}

}
func godotCallObjectStringBoolBool(o Class, methodName string, arg0 string, arg1 bool, arg2 bool) *Object {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_bool(arg1)
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_object_string_bool_bool...")
	cRet := C.godot_call_object_string_bool_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_object_string_bool_bool returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_object)(cRet)
	return &Object{ret}

}
func godotCallVoidRidTransform2D(o Class, methodName string, arg0 *RID, arg1 *Transform2D) {

	cArg0 := arg0.rid
	cArg1 := arg1.transform2d

	log.Println("  Calling godot_call_void_rid_transform2d...")
	C.godot_call_void_rid_transform2d(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_rid_transform2d returned")

}
func godotCallArrayPoolByteArray(o Class, methodName string, arg0 *PoolByteArray) *Array {

	cArg0 := arg0.array

	log.Println("  Calling godot_call_array_poolbytearray...")
	cRet := C.godot_call_array_poolbytearray(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_array_poolbytearray returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_array)(cRet)
	return &Array{ret}

}
func godotCallVector3IntFloat(o Class, methodName string, arg0 int64, arg1 float64) *Vector3 {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_real(arg1)

	log.Println("  Calling godot_call_vector3_int_float...")
	cRet := C.godot_call_vector3_int_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_vector3_int_float returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector3)(cRet)
	return &Vector3{ret}

}
func godotCallVoidRidRect2Int(o Class, methodName string, arg0 *RID, arg1 *Rect2, arg2 int64) {

	cArg0 := arg0.rid
	cArg1 := arg1.rect2
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_void_rid_rect2_int...")
	C.godot_call_void_rid_rect2_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_rid_rect2_int returned")

}
func godotCallVoidPoolIntArray(o Class, methodName string, arg0 *PoolIntArray) {

	cArg0 := arg0.array

	log.Println("  Calling godot_call_void_poolintarray...")
	C.godot_call_void_poolintarray(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_void_poolintarray returned")

}
func godotCallFloat(o Class, methodName string) float64 {

	log.Println("  Calling godot_call_float...")
	cRet := C.godot_call_float(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_float returned")
	ret := cRet
	return float64(ret)

}
func godotCallVoidStringBool(o Class, methodName string, arg0 string, arg1 bool) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_bool(arg1)

	log.Println("  Calling godot_call_void_string_bool...")
	C.godot_call_void_string_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_string_bool returned")

}
func godotCallVoidIntBoolInt(o Class, methodName string, arg0 int64, arg1 bool, arg2 int64) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_bool(arg1)
	cArg2 := C.godot_int(arg2)

	log.Println("  Calling godot_call_void_int_bool_int...")
	C.godot_call_void_int_bool_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_int_bool_int returned")

}
func godotCallVoidRidPoolVector2ArrayPoolColorArrayPoolVector2ArrayRidFloatRid(o Class, methodName string, arg0 *RID, arg1 *PoolVector2Array, arg2 *PoolColorArray, arg3 *PoolVector2Array, arg4 *RID, arg5 float64, arg6 *RID) {

	cArg0 := arg0.rid
	cArg1 := arg1.array
	cArg2 := arg2.array
	cArg3 := arg3.array
	cArg4 := arg4.rid
	cArg5 := C.godot_real(arg5)
	cArg6 := arg6.rid

	log.Println("  Calling godot_call_void_rid_poolvector2array_poolcolorarray_poolvector2array_rid_float_rid...")
	C.godot_call_void_rid_poolvector2array_poolcolorarray_poolvector2array_rid_float_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4, cArg5, cArg6,
	)
	log.Println("  ...godot_call_void_rid_poolvector2array_poolcolorarray_poolvector2array_rid_float_rid returned")

}
func godotCallAabb(o Class, methodName string) *AABB {

	log.Println("  Calling godot_call_aabb...")
	cRet := C.godot_call_aabb(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_aabb returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_aabb)(cRet)
	return &AABB{ret}

}
func godotCallIntString(o Class, methodName string, arg0 string) int64 {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))

	log.Println("  Calling godot_call_int_string...")
	cRet := C.godot_call_int_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_int_string returned")
	ret := cRet
	return int64(ret)

}
func godotCallPoolIntArray(o Class, methodName string) *PoolIntArray {

	log.Println("  Calling godot_call_poolintarray...")
	cRet := C.godot_call_poolintarray(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_poolintarray returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_pool_int_array)(cRet)
	return &PoolIntArray{ret}

}
func godotCallTransformInt(o Class, methodName string, arg0 int64) *Transform {

	cArg0 := C.godot_int(arg0)

	log.Println("  Calling godot_call_transform_int...")
	cRet := C.godot_call_transform_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_transform_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_transform)(cRet)
	return &Transform{ret}

}
func godotCallVoidIntFloatFloatBool(o Class, methodName string, arg0 int64, arg1 float64, arg2 float64, arg3 bool) {

	cArg0 := C.godot_int(arg0)
	cArg1 := C.godot_real(arg1)
	cArg2 := C.godot_real(arg2)
	cArg3 := C.godot_bool(arg3)

	log.Println("  Calling godot_call_void_int_float_float_bool...")
	C.godot_call_void_int_float_float_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_int_float_float_bool returned")

}
func godotCallVoidBoolBoolIntInt(o Class, methodName string, arg0 bool, arg1 bool, arg2 int64, arg3 int64) {

	cArg0 := C.godot_bool(arg0)
	cArg1 := C.godot_bool(arg1)
	cArg2 := C.godot_int(arg2)
	cArg3 := C.godot_int(arg3)

	log.Println("  Calling godot_call_void_bool_bool_int_int...")
	C.godot_call_void_bool_bool_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3,
	)
	log.Println("  ...godot_call_void_bool_bool_int_int returned")

}
func godotCallDictionaryString(o Class, methodName string, arg0 string) *Dictionary {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))

	log.Println("  Calling godot_call_dictionary_string...")
	cRet := C.godot_call_dictionary_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_dictionary_string returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_dictionary)(cRet)
	return &Dictionary{ret}

}
func godotCallVoidObjectBoolRid(o Class, methodName string, arg0 *Object, arg1 bool, arg2 *RID) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.godot_bool(arg1)
	cArg2 := arg2.rid

	log.Println("  Calling godot_call_void_object_bool_rid...")
	C.godot_call_void_object_bool_rid(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_object_bool_rid returned")

}
func godotCallVector3Vector3Vector3Bool(o Class, methodName string, arg0 *Vector3, arg1 *Vector3, arg2 bool) *Vector3 {

	cArg0 := arg0.vector3
	cArg1 := arg1.vector3
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_vector3_vector3_vector3_bool...")
	cRet := C.godot_call_vector3_vector3_vector3_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_vector3_vector3_vector3_bool returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector3)(cRet)
	return &Vector3{ret}

}
func godotCallVoidObjectIntBool(o Class, methodName string, arg0 *Object, arg1 int64, arg2 bool) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_bool(arg2)

	log.Println("  Calling godot_call_void_object_int_bool...")
	C.godot_call_void_object_int_bool(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_object_int_bool returned")

}
func godotCallVoidStringIntIntIntInt(o Class, methodName string, arg0 string, arg1 int64, arg2 int64, arg3 int64, arg4 int64) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := C.godot_int(arg1)
	cArg2 := C.godot_int(arg2)
	cArg3 := C.godot_int(arg3)
	cArg4 := C.godot_int(arg4)

	log.Println("  Calling godot_call_void_string_int_int_int_int...")
	C.godot_call_void_string_int_int_int_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_void_string_int_int_int_int returned")

}
func godotCallVector3Int(o Class, methodName string, arg0 int64) *Vector3 {

	cArg0 := C.godot_int(arg0)

	log.Println("  Calling godot_call_vector3_int...")
	cRet := C.godot_call_vector3_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_vector3_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_vector3)(cRet)
	return &Vector3{ret}

}
func godotCallObject(o Class, methodName string) *Object {

	log.Println("  Calling godot_call_object...")
	cRet := C.godot_call_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
	)
	log.Println("  ...godot_call_object returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_object)(cRet)
	return &Object{ret}

}
func godotCallTransform2DObject(o Class, methodName string, arg0 *Object) *Transform2D {

	cArg0 := unsafe.Pointer(arg0.owner)

	log.Println("  Calling godot_call_transform2d_object...")
	cRet := C.godot_call_transform2d_object(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0,
	)
	log.Println("  ...godot_call_transform2d_object returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_transform2d)(cRet)
	return &Transform2D{ret}

}
func godotCallVoidStringObjectString(o Class, methodName string, arg0 string, arg1 *Object, arg2 string) {

	cArg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(cArg0))
	cArg1 := unsafe.Pointer(arg1.owner)
	cArg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(cArg2))

	log.Println("  Calling godot_call_void_string_object_string...")
	C.godot_call_void_string_object_string(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2,
	)
	log.Println("  ...godot_call_void_string_object_string returned")

}
func godotCallTransformRidInt(o Class, methodName string, arg0 *RID, arg1 int64) *Transform {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_transform_rid_int...")
	cRet := C.godot_call_transform_rid_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_transform_rid_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_transform)(cRet)
	return &Transform{ret}

}
func godotCallVoidObjectVector2StringColorInt(o Class, methodName string, arg0 *Object, arg1 *Vector2, arg2 string, arg3 *Color, arg4 int64) {

	cArg0 := unsafe.Pointer(arg0.owner)
	cArg1 := arg1.vector2
	cArg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(cArg2))
	cArg3 := arg3.color
	cArg4 := C.godot_int(arg4)

	log.Println("  Calling godot_call_void_object_vector2_string_color_int...")
	C.godot_call_void_object_vector2_string_color_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1, cArg2, cArg3, cArg4,
	)
	log.Println("  ...godot_call_void_object_vector2_string_color_int returned")

}
func godotCallVoidIntVector2(o Class, methodName string, arg0 int64, arg1 *Vector2) {

	cArg0 := C.godot_int(arg0)
	cArg1 := arg1.vector2

	log.Println("  Calling godot_call_void_int_vector2...")
	C.godot_call_void_int_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_int_vector2 returned")

}
func godotCallVoidVector2Vector2(o Class, methodName string, arg0 *Vector2, arg1 *Vector2) {

	cArg0 := arg0.vector2
	cArg1 := arg1.vector2

	log.Println("  Calling godot_call_void_vector2_vector2...")
	C.godot_call_void_vector2_vector2(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_void_vector2_vector2 returned")

}
func godotCallObjectRidInt(o Class, methodName string, arg0 *RID, arg1 int64) *Object {

	cArg0 := arg0.rid
	cArg1 := C.godot_int(arg1)

	log.Println("  Calling godot_call_object_rid_int...")
	cRet := C.godot_call_object_rid_int(
		unsafe.Pointer(o.getOwner()),
		getGodotMethod(o.baseClass(), methodName),
		cArg0, cArg1,
	)
	log.Println("  ...godot_call_object_rid_int returned")
	if cRet == nil {
		return nil
	}
	ret := (*C.godot_object)(cRet)
	return &Object{ret}

}
