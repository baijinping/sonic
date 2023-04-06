// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__f32toa = 43264
    _entry__f64toa = 800
    _entry__format_significand = 46656
    _entry__format_integer = 3648
    _entry__fsm_exec = 22144
    _entry__advance_string = 18528
    _entry__advance_string_default = 48096
    _entry__do_skip_number = 25520
    _entry__get_by_path = 31744
    _entry__skip_one_fast = 27904
    _entry__html_escape = 11040
    _entry__i64toa = 4080
    _entry__u64toa = 4192
    _entry__lspace = 224
    _entry__quote = 5584
    _entry__skip_array = 21888
    _entry__skip_number = 27216
    _entry__skip_object = 24928
    _entry__skip_one = 27360
    _entry__unquote = 8368
    _entry__validate_one = 27408
    _entry__validate_utf8 = 39664
    _entry__validate_utf8_fast = 40624
    _entry__value = 16064
    _entry__vnumber = 19664
    _entry__atof_eisel_lemire64 = 13104
    _entry__atof_native = 15248
    _entry__decimal_to_f64 = 13536
    _entry__right_shift = 47616
    _entry__left_shift = 47120
    _entry__vsigned = 21216
    _entry__vstring = 18304
    _entry__vunsigned = 21536
)

const (
    _stack__f32toa = 48
    _stack__f64toa = 80
    _stack__format_significand = 24
    _stack__format_integer = 16
    _stack__fsm_exec = 144
    _stack__advance_string = 56
    _stack__advance_string_default = 48
    _stack__do_skip_number = 48
    _stack__get_by_path = 280
    _stack__skip_one_fast = 176
    _stack__html_escape = 72
    _stack__i64toa = 16
    _stack__u64toa = 8
    _stack__lspace = 8
    _stack__quote = 56
    _stack__skip_array = 152
    _stack__skip_number = 88
    _stack__skip_object = 152
    _stack__skip_one = 152
    _stack__unquote = 72
    _stack__validate_one = 152
    _stack__validate_utf8 = 48
    _stack__validate_utf8_fast = 176
    _stack__value = 328
    _stack__vnumber = 240
    _stack__atof_eisel_lemire64 = 32
    _stack__atof_native = 136
    _stack__decimal_to_f64 = 80
    _stack__right_shift = 8
    _stack__left_shift = 24
    _stack__vsigned = 16
    _stack__vstring = 112
    _stack__vunsigned = 8
)

const (
    _size__f32toa = 3380
    _size__f64toa = 2838
    _size__format_significand = 457
    _size__format_integer = 426
    _size__fsm_exec = 2231
    _size__advance_string = 1073
    _size__advance_string_default = 752
    _size__do_skip_number = 1348
    _size__get_by_path = 7900
    _size__skip_one_fast = 3136
    _size__html_escape = 2063
    _size__i64toa = 36
    _size__u64toa = 1227
    _size__lspace = 514
    _size__quote = 2726
    _size__skip_array = 29
    _size__skip_number = 134
    _size__skip_object = 29
    _size__skip_one = 31
    _size__unquote = 2459
    _size__validate_one = 36
    _size__validate_utf8 = 661
    _size__validate_utf8_fast = 2589
    _size__value = 1728
    _size__vnumber = 1542
    _size__atof_eisel_lemire64 = 357
    _size__atof_native = 593
    _size__decimal_to_f64 = 1697
    _size__right_shift = 398
    _size__left_shift = 475
    _size__vsigned = 319
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
        {3351, 40},
        {3353, 32},
        {3355, 24},
        {3357, 16},
        {3359, 8},
        {3360, 0},
        {3364, 48},
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
        {2792, 48},
        {2793, 40},
        {2795, 32},
        {2797, 24},
        {2799, 16},
        {2801, 8},
        {2802, 0},
        {2806, 56},
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
        {1869, 48},
        {1870, 40},
        {1872, 32},
        {1874, 24},
        {1876, 16},
        {1878, 8},
        {1879, 0},
        {1883, 88},
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
        {437, 48},
        {438, 40},
        {440, 32},
        {442, 24},
        {444, 16},
        {446, 8},
        {447, 0},
        {451, 56},
    }
    _pcsp__advance_string_default = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {333, 40},
        {335, 32},
        {337, 24},
        {339, 16},
        {341, 8},
        {342, 0},
        {346, 48},
    }
    _pcsp__do_skip_number = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1275, 40},
        {1277, 32},
        {1279, 24},
        {1281, 16},
        {1283, 8},
        {1284, 0},
        {1288, 48},
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
        {7745, 48},
        {7746, 40},
        {7748, 32},
        {7750, 24},
        {7752, 16},
        {7754, 8},
        {7755, 0},
        {7759, 104},
    }
    _pcsp__skip_one_fast = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {20, 176},
        {1167, 48},
        {1168, 40},
        {1170, 32},
        {1172, 24},
        {1174, 16},
        {1176, 8},
        {1177, 0},
        {1181, 176},
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
        {2049, 48},
        {2050, 40},
        {2052, 32},
        {2054, 24},
        {2056, 16},
        {2058, 8},
        {2059, 0},
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
        {759, 0},
        {760, 8},
        {1226, 0},
    }
    _pcsp__lspace = [][2]uint32{
        {0, 0},
        {1, 8},
        {481, 0},
        {482, 8},
        {489, 0},
        {490, 8},
        {505, 0},
        {506, 8},
        {513, 0},
    }
    _pcsp__quote = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {14, 56},
        {2691, 48},
        {2692, 40},
        {2694, 32},
        {2696, 24},
        {2698, 16},
        {2700, 8},
        {2701, 0},
        {2705, 56},
    }
    _pcsp__skip_array = [][2]uint32{
        {0, 0},
        {1, 8},
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
    }
    _pcsp__skip_one = [][2]uint32{
        {0, 0},
        {1, 8},
    }
    _pcsp__unquote = [][2]uint32{
        {0, 0},
        {1, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {17, 72},
        {83, 48},
        {84, 40},
        {86, 32},
        {88, 24},
        {90, 16},
        {92, 8},
        {93, 0},
        {97, 72},
    }
    _pcsp__validate_one = [][2]uint32{
        {0, 0},
        {1, 8},
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
        {12, 176},
        {1738, 16},
        {1739, 8},
        {1740, 0},
        {1744, 176},
        {2018, 16},
        {2019, 8},
        {2020, 0},
        {2024, 176},
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
        {886, 48},
        {887, 40},
        {889, 32},
        {891, 24},
        {893, 16},
        {895, 8},
        {896, 0},
        {897, 88},
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
        {1677, 48},
        {1678, 40},
        {1680, 32},
        {1682, 24},
        {1684, 16},
        {1686, 8},
        {1687, 0},
        {1691, 56},
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
    }
    _pcsp__vsigned = [][2]uint32{
        {0, 0},
        {1, 8},
        {5, 16},
        {113, 8},
        {114, 0},
        {115, 16},
        {126, 8},
        {127, 0},
        {128, 16},
        {261, 8},
        {262, 0},
        {263, 16},
        {267, 8},
        {268, 0},
        {269, 16},
        {307, 8},
        {308, 0},
        {309, 16},
        {317, 8},
        {318, 0},
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
        {72, 0},
        {73, 8},
        {84, 0},
        {85, 8},
        {108, 0},
        {109, 8},
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
