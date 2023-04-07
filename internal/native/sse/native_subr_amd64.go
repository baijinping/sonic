// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__f32toa = 29504
    _entry__f64toa = 464
    _entry__format_significand = 32832
    _entry__format_integer = 3264
    _entry__fsm_exec = 18768
    _entry__advance_ns = 14240
    _entry__advance_string = 15104
    _entry__advance_string_default = 34224
    _entry__do_skip_number = 21120
    _entry__get_by_path = 26336
    _entry__skip_one_fast = 22784
    _entry__html_escape = 9216
    _entry__i64toa = 3696
    _entry__u64toa = 3824
    _entry__lspace = 80
    _entry__quote = 5136
    _entry__skip_array = 18736
    _entry__skip_number = 22416
    _entry__skip_object = 20768
    _entry__skip_one = 22560
    _entry__unquote = 6880
    _entry__validate_one = 22608
    _entry__validate_utf8 = 28272
    _entry__validate_utf8_fast = 28944
    _entry__value = 12656
    _entry__vnumber = 16496
    _entry__atof_eisel_lemire64 = 10496
    _entry__atof_native = 12048
    _entry__decimal_to_f64 = 10864
    _entry__right_shift = 33792
    _entry__left_shift = 33296
    _entry__vsigned = 18048
    _entry__vstring = 14928
    _entry__vunsigned = 18384
)

const (
    _stack__f32toa = 48
    _stack__f64toa = 88
    _stack__format_significand = 24
    _stack__format_integer = 16
    _stack__fsm_exec = 168
    _stack__advance_ns = 16
    _stack__advance_string = 64
    _stack__advance_string_default = 64
    _stack__do_skip_number = 48
    _stack__get_by_path = 320
    _stack__skip_one_fast = 176
    _stack__html_escape = 72
    _stack__i64toa = 24
    _stack__u64toa = 8
    _stack__lspace = 8
    _stack__quote = 64
    _stack__skip_array = 176
    _stack__skip_number = 96
    _stack__skip_object = 176
    _stack__skip_one = 176
    _stack__unquote = 88
    _stack__validate_one = 176
    _stack__validate_utf8 = 48
    _stack__validate_utf8_fast = 24
    _stack__value = 416
    _stack__vnumber = 312
    _stack__atof_eisel_lemire64 = 32
    _stack__atof_native = 184
    _stack__decimal_to_f64 = 120
    _stack__right_shift = 8
    _stack__left_shift = 24
    _stack__vsigned = 16
    _stack__vstring = 128
    _stack__vunsigned = 8
)

const (
    _size__f32toa = 3318
    _size__f64toa = 2792
    _size__format_significand = 457
    _size__format_integer = 426
    _size__fsm_exec = 1464
    _size__advance_ns = 686
    _size__advance_string = 1339
    _size__advance_string_default = 955
    _size__do_skip_number = 956
    _size__get_by_path = 1933
    _size__skip_one_fast = 3045
    _size__html_escape = 1271
    _size__i64toa = 37
    _size__u64toa = 1251
    _size__lspace = 362
    _size__quote = 1722
    _size__skip_array = 32
    _size__skip_number = 139
    _size__skip_object = 32
    _size__skip_one = 36
    _size__unquote = 2270
    _size__validate_one = 41
    _size__validate_utf8 = 666
    _size__validate_utf8_fast = 534
    _size__value = 1003
    _size__vnumber = 1551
    _size__atof_eisel_lemire64 = 362
    _size__atof_native = 593
    _size__decimal_to_f64 = 1169
    _size__right_shift = 398
    _size__left_shift = 487
    _size__vsigned = 335
    _size__vstring = 118
    _size__vunsigned = 322
)

var (
    _pcsp__f32toa = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {3287, 40},
        {3289, 32},
        {3291, 24},
        {3293, 16},
        {3295, 8},
        {3296, 0},
        {3297, 48},
    }
    _pcsp__f64toa = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {14, 56},
        {2744, 48},
        {2745, 40},
        {2747, 32},
        {2749, 24},
        {2751, 16},
        {2753, 8},
        {2754, 0},
        {2755, 56},
    }
    _pcsp__format_significand = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {7, 24},
        {453, 16},
        {455, 8},
        {456, 0},
    }
    _pcsp__format_integer = [][2]uint32{
        {0, 0},
        {1, 8},
        {5, 16},
        {413, 8},
        {414, 0},
        {415, 16},
        {424, 8},
        {425, 0},
    }
    _pcsp__fsm_exec = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {17, 88},
        {1198, 48},
        {1199, 40},
        {1201, 32},
        {1203, 24},
        {1205, 16},
        {1207, 8},
        {1208, 0},
        {1209, 88},
    }
    _pcsp__advance_ns = [][2]uint32{
        {0, 0},
        {1, 8},
        {5, 16},
        {658, 8},
        {659, 0},
        {660, 16},
        {684, 8},
        {685, 0},
    }
    _pcsp__advance_string = [][2]uint32{
        {0, 0},
        {15, 8},
        {20, 16},
        {22, 24},
        {24, 32},
        {26, 40},
        {27, 48},
        {28, 56},
        {618, 48},
        {619, 40},
        {621, 32},
        {623, 24},
        {625, 16},
        {627, 8},
        {628, 0},
        {629, 56},
    }
    _pcsp__advance_string_default = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {17, 64},
        {580, 48},
        {581, 40},
        {583, 32},
        {585, 24},
        {587, 16},
        {589, 8},
        {590, 0},
        {591, 64},
    }
    _pcsp__do_skip_number = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {882, 40},
        {884, 32},
        {886, 24},
        {888, 16},
        {890, 8},
        {891, 0},
        {892, 48},
    }
    _pcsp__get_by_path = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {17, 104},
        {1815, 48},
        {1816, 40},
        {1818, 32},
        {1820, 24},
        {1822, 16},
        {1824, 8},
        {1825, 0},
        {1826, 104},
    }
    _pcsp__skip_one_fast = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {17, 152},
        {317, 48},
        {318, 40},
        {320, 32},
        {322, 24},
        {324, 16},
        {326, 8},
        {327, 0},
        {328, 152},
    }
    _pcsp__html_escape = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {17, 72},
        {1260, 48},
        {1261, 40},
        {1263, 32},
        {1265, 24},
        {1267, 16},
        {1269, 8},
        {1270, 0},
    }
    _pcsp__i64toa = [][2]uint32{
        {0, 0},
        {15, 8},
        {35, 0},
    }
    _pcsp__u64toa = [][2]uint32{
        {0, 0},
        {1, 8},
        {162, 0},
        {163, 8},
        {458, 0},
        {459, 8},
        {773, 0},
        {774, 8},
        {1250, 0},
    }
    _pcsp__lspace = [][2]uint32{
        {0, 0},
        {1, 8},
        {345, 0},
        {346, 8},
        {353, 0},
        {354, 8},
        {361, 0},
    }
    _pcsp__quote = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {17, 64},
        {1685, 48},
        {1686, 40},
        {1688, 32},
        {1690, 24},
        {1692, 16},
        {1694, 8},
        {1695, 0},
        {1696, 64},
    }
    _pcsp__skip_array = [][2]uint32{
        {0, 0},
        {1, 8},
        {27, 0},
    }
    _pcsp__skip_number = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {11, 40},
        {101, 32},
        {103, 24},
        {105, 16},
        {107, 8},
        {108, 0},
        {109, 40},
    }
    _pcsp__skip_object = [][2]uint32{
        {0, 0},
        {1, 8},
        {27, 0},
    }
    _pcsp__skip_one = [][2]uint32{
        {0, 0},
        {1, 8},
        {31, 0},
    }
    _pcsp__unquote = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {17, 88},
        {1688, 48},
        {1689, 40},
        {1691, 32},
        {1693, 24},
        {1695, 16},
        {1697, 8},
        {1698, 0},
        {1699, 88},
    }
    _pcsp__validate_one = [][2]uint32{
        {0, 0},
        {1, 8},
        {36, 0},
    }
    _pcsp__validate_utf8 = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {11, 40},
        {12, 48},
        {627, 40},
        {628, 32},
        {630, 24},
        {632, 16},
        {634, 8},
        {635, 0},
        {636, 48},
    }
    _pcsp__validate_utf8_fast = [][2]uint32{
        {0, 0},
        {1, 8},
        {5, 16},
        {6, 24},
        {251, 16},
        {252, 8},
        {253, 0},
        {254, 24},
        {531, 16},
        {532, 8},
        {533, 0},
    }
    _pcsp__value = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {17, 88},
        {992, 48},
        {993, 40},
        {995, 32},
        {997, 24},
        {999, 16},
        {1001, 8},
        {1002, 0},
    }
    _pcsp__vnumber = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {17, 104},
        {807, 48},
        {808, 40},
        {810, 32},
        {812, 24},
        {814, 16},
        {816, 8},
        {817, 0},
        {818, 104},
    }
    _pcsp__atof_eisel_lemire64 = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {8, 24},
        {9, 32},
        {293, 24},
        {295, 16},
        {297, 8},
        {298, 0},
        {299, 32},
    }
    _pcsp__atof_native = [][2]uint32{
        {0, 0},
        {1, 8},
        {8, 56},
        {591, 8},
        {592, 0},
    }
    _pcsp__decimal_to_f64 = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {14, 56},
        {1148, 48},
        {1149, 40},
        {1151, 32},
        {1153, 24},
        {1155, 16},
        {1157, 8},
        {1158, 0},
        {1159, 56},
    }
    _pcsp__right_shift = [][2]uint32{
        {0, 0},
        {1, 8},
        {319, 0},
        {320, 8},
        {388, 0},
        {389, 8},
        {397, 0},
    }
    _pcsp__left_shift = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {7, 24},
        {364, 16},
        {366, 8},
        {367, 0},
        {368, 24},
        {471, 16},
        {473, 8},
        {474, 0},
        {475, 24},
    }
    _pcsp__vsigned = [][2]uint32{
        {0, 0},
        {1, 8},
        {5, 16},
        {120, 8},
        {121, 0},
        {122, 16},
        {133, 8},
        {134, 0},
        {135, 16},
        {277, 8},
        {278, 0},
        {279, 16},
        {283, 8},
        {284, 0},
        {285, 16},
        {323, 8},
        {324, 0},
        {325, 16},
        {333, 8},
        {334, 0},
    }
    _pcsp__vstring = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {11, 40},
        {15, 56},
        {109, 40},
        {110, 32},
        {112, 24},
        {114, 16},
        {116, 8},
        {117, 0},
    }
    _pcsp__vunsigned = [][2]uint32{
        {0, 0},
        {1, 8},
        {79, 0},
        {80, 8},
        {91, 0},
        {92, 8},
        {115, 0},
        {116, 8},
        {274, 0},
        {275, 8},
        {313, 0},
        {314, 8},
        {321, 0},
    }
)

var Funcs = []loader.CFunc{
    {"__native_entry__", 0, 67, 0, nil},
    {"_f32toa", _entry__f32toa, _size__f32toa, _stack__f32toa, _pcsp__f32toa},
    {"_f64toa", _entry__f64toa, _size__f64toa, _stack__f64toa, _pcsp__f64toa},
    {"_format_significand", _entry__format_significand, _size__format_significand, _stack__format_significand, _pcsp__format_significand},
    {"_format_integer", _entry__format_integer, _size__format_integer, _stack__format_integer, _pcsp__format_integer},
    {"_fsm_exec", _entry__fsm_exec, _size__fsm_exec, _stack__fsm_exec, _pcsp__fsm_exec},
    {"_advance_ns", _entry__advance_ns, _size__advance_ns, _stack__advance_ns, _pcsp__advance_ns},
    {"_advance_string", _entry__advance_string, _size__advance_string, _stack__advance_string, _pcsp__advance_string},
    {"_advance_string_default", _entry__advance_string_default, _size__advance_string_default, _stack__advance_string_default, _pcsp__advance_string_default},
    {"_do_skip_number", _entry__do_skip_number, _size__do_skip_number, _stack__do_skip_number, _pcsp__do_skip_number},
    {"_get_by_path", _entry__get_by_path, _size__get_by_path, _stack__get_by_path, _pcsp__get_by_path},
    {"_skip_one_fast", _entry__skip_one_fast, _size__skip_one_fast, _stack__skip_one_fast, _pcsp__skip_one_fast},
    {"_html_escape", _entry__html_escape, _size__html_escape, _stack__html_escape, _pcsp__html_escape},
    {"_i64toa", _entry__i64toa, _size__i64toa, _stack__i64toa, _pcsp__i64toa},
    {"_u64toa", _entry__u64toa, _size__u64toa, _stack__u64toa, _pcsp__u64toa},
    {"_lspace", _entry__lspace, _size__lspace, _stack__lspace, _pcsp__lspace},
    {"_quote", _entry__quote, _size__quote, _stack__quote, _pcsp__quote},
    {"_skip_array", _entry__skip_array, _size__skip_array, _stack__skip_array, _pcsp__skip_array},
    {"_skip_number", _entry__skip_number, _size__skip_number, _stack__skip_number, _pcsp__skip_number},
    {"_skip_object", _entry__skip_object, _size__skip_object, _stack__skip_object, _pcsp__skip_object},
    {"_skip_one", _entry__skip_one, _size__skip_one, _stack__skip_one, _pcsp__skip_one},
    {"_unquote", _entry__unquote, _size__unquote, _stack__unquote, _pcsp__unquote},
    {"_validate_one", _entry__validate_one, _size__validate_one, _stack__validate_one, _pcsp__validate_one},
    {"_validate_utf8", _entry__validate_utf8, _size__validate_utf8, _stack__validate_utf8, _pcsp__validate_utf8},
    {"_validate_utf8_fast", _entry__validate_utf8_fast, _size__validate_utf8_fast, _stack__validate_utf8_fast, _pcsp__validate_utf8_fast},
    {"_value", _entry__value, _size__value, _stack__value, _pcsp__value},
    {"_vnumber", _entry__vnumber, _size__vnumber, _stack__vnumber, _pcsp__vnumber},
    {"_atof_eisel_lemire64", _entry__atof_eisel_lemire64, _size__atof_eisel_lemire64, _stack__atof_eisel_lemire64, _pcsp__atof_eisel_lemire64},
    {"_atof_native", _entry__atof_native, _size__atof_native, _stack__atof_native, _pcsp__atof_native},
    {"_decimal_to_f64", _entry__decimal_to_f64, _size__decimal_to_f64, _stack__decimal_to_f64, _pcsp__decimal_to_f64},
    {"_right_shift", _entry__right_shift, _size__right_shift, _stack__right_shift, _pcsp__right_shift},
    {"_left_shift", _entry__left_shift, _size__left_shift, _stack__left_shift, _pcsp__left_shift},
    {"_vsigned", _entry__vsigned, _size__vsigned, _stack__vsigned, _pcsp__vsigned},
    {"_vstring", _entry__vstring, _size__vstring, _stack__vstring, _pcsp__vstring},
    {"_vunsigned", _entry__vunsigned, _size__vunsigned, _stack__vunsigned, _pcsp__vunsigned},
}
