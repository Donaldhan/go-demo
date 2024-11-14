// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

/**
 * 复杂类型
 * 一维数组,传参，返回参数
 * 二位数组,传参，返回参数
 * 结构体：包含原始类型及以为数组，二位数组；,传参，返回参数
 * 结构体配置，及结构体数组配置,传参，返回参数
 */
// import "./ComplexTypeData.sol";
// contract ComplexType is ComplexTypeData {
contract ComplexType {
    event AddUser(uint256 indexed id, bytes name, uint256 ages, bool sex);
    event AddUserPlayDays(
        uint256 indexed id,
        uint256 birthTime, //出生时间
        uint16[] playMonths, //每年玩的月数
        uint256[][] playMonthDays //每月玩的月数
    );
    event UpdateUserPlayMonth(
        uint256 indexed id,
        uint16[] playMonths, //每年玩的月数
        uint256[][] playMonthDays //每月玩的月数
    );

    // uint16[] months; //月数
    // uint256[][] monthDays; //每月的天数
    // * 复杂类型
    // * 一维数组,传参，返回参数
    // * 二位数组,传参，返回参数
    // * 结构体：包含原始类型及以为数组，二位数组；,传参，返回参数
    // * 结构体配置，及结构体数组配置,传参，返回参数
    //生日
    struct People {
        uint256 id;
        bytes name; //名称
        uint256 ages; //年龄
        bool sex; //性别
        uint256 birthTime; //出生时间
        uint16[] playMonths; //每年玩的月数
        uint256[][] playMonthDays; //每月玩的日期（1-31），1,3,5
    }
    struct PeopleBaseConfig {
        bytes name; //名称
        uint256 ages; //年龄
        bool sex; //性别
    }
    struct PeopleBase {
        uint256 ages; //年龄
        bool sex; //性别
    }
    struct PlayDayConfig {
        uint256 birthTime; //出生时间
        uint16[] playMonths; //每年玩的月数
        uint256[][] playMonthDays; //每月玩的日期（1-31），1,3,5
    }
    mapping(uint256 => People) peopleMap; //用户游玩数据

    function addUser(uint256 id, PeopleBaseConfig calldata baseConfig) public {
        People memory people = peopleMap[id];
        people.name = baseConfig.name;
        people.ages = baseConfig.ages;
        people.sex = baseConfig.sex;
        peopleMap[id] = people;
        emit AddUser(id, people.name, people.ages, people.sex);
    }

    function addBatchUser(
        uint256[] calldata idList,
        PeopleBaseConfig[] calldata baseConfigList
    ) public {
        for (uint256 i = 0; i < idList.length; i++) {
            uint256 id = idList[i];
            PeopleBaseConfig calldata baseConfig = baseConfigList[i];
            addUser(id, baseConfig);
        }
    }

    function addUserBase(uint256 id, PeopleBase calldata base) public {
        People memory people = peopleMap[id];
        people.ages = base.ages;
        people.sex = base.sex;
        peopleMap[id] = people;
        emit AddUser(id, people.name, people.ages, people.sex);
    }

    //传struct
    function addUserPlayDays(
        uint256 id,
        PlayDayConfig calldata dayConfig
    ) public {
        People memory people = peopleMap[id];
        people.birthTime = dayConfig.birthTime;
        people.playMonths = dayConfig.playMonths;
        people.playMonthDays = dayConfig.playMonthDays;
        peopleMap[id] = people;
        emit AddUserPlayDays(
            id,
            people.birthTime,
            people.playMonths,
            people.playMonthDays
        );
    }

    //批量模式
    function addBatchUserPlayDays(
        uint256[] calldata idList,
        PlayDayConfig[] calldata dayConfigList
    ) public {
        for (uint256 i = 0; i < idList.length; i++) {
            uint256 id = idList[i];
            PlayDayConfig calldata dayConfig = dayConfigList[i];
            addUserPlayDays(id, dayConfig);
        }
    }

    //返回dynamic struct
    function getUserPlayDays(
        uint256 id
    ) public view returns (PlayDayConfig memory) {
        PlayDayConfig memory dayConfig;
        People memory people = peopleMap[id];
        dayConfig.birthTime = people.birthTime;
        dayConfig.playMonths = people.playMonths;
        dayConfig.playMonthDays = people.playMonthDays;
        return dayConfig;
    }

    //返回 struct
    function getBatchUserPlayDays(
        uint256[] calldata idList
    ) public view returns (PlayDayConfig[] memory) {
        PlayDayConfig[] memory dayConfigList = new PlayDayConfig[](
            idList.length
        );
        for (uint256 i = 0; i < idList.length; i++) {
            uint256 id = idList[i];
            People memory people = peopleMap[id];
            dayConfigList[i].playMonths = people.playMonths;
            dayConfigList[i].playMonthDays = people.playMonthDays;
        }
        return dayConfigList;
    }
    //返回struct数组
    function getBatchUser(
        uint256[] calldata idList
    ) public view returns (People[] memory) {
        People[] memory peopleList = new People[](idList.length);
        for (uint256 i = 0; i < idList.length; i++) {
            uint256 id = idList[i];
            People memory people = peopleMap[id];
            peopleList[i] = people;
        }
        return peopleList;
    }


    //返回static struct
    function getBatchUserBase(
        uint256[] calldata idList
    ) public view returns (PeopleBase[] memory) {
        PeopleBase[] memory useBaseList = new PeopleBase[](
            idList.length
        );
        for (uint256 i = 0; i < idList.length; i++) {
            uint256 id = idList[i];
            People memory people = peopleMap[id];
            useBaseList[i].ages = people.ages;
            useBaseList[i].sex = people.sex;
        }
        return useBaseList;
    }
    //数组传参
    function updateUserPlayMonth(
        uint256 id,
        uint16[] calldata playMonths,
        uint256[][] calldata playMonthDays
    ) public {
        People memory people = peopleMap[id];
        people.playMonths = playMonths;
        people.playMonthDays = playMonthDays;
        peopleMap[id] = people;
        emit UpdateUserPlayMonth(id, playMonths, playMonthDays);
    }

    //返回数组
    function getUserPlayMonthSingle(
        uint256 id
    ) public view returns (uint16[] memory, uint256[][] memory) {
        People memory people = peopleMap[id];
        return (people.playMonths, people.playMonthDays);
    }

    //返回对象
    function getUser(uint256 id) public view returns (People memory) {
        People memory people = peopleMap[id];
        return people;
    }
}
