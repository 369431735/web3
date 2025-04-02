// SPDX-License-Identifier: MIT

pragma solidity 0.6.12;

interface IUniswapV2Router02 {
    // 返回工厂合约地址
    function factory() external pure returns (address);
    // 返回WETH代币地址
    function WETH() external pure returns (address);

    // 添加流动性
    // tokenA和tokenB是要添加的两种代币地址
    // amount[A/B]Desired是期望添加的数量
    // amount[A/B]Min是最小接受的数量
    // to是接收LP代币的地址
    // deadline是交易截止时间
    function addLiquidity(
        address tokenA,
        address tokenB,
        uint amountADesired,
        uint amountBDesired,
        uint amountAMin,
        uint amountBMin,
        address to,
        uint deadline
    ) external returns (uint amountA, uint amountB, uint liquidity);

    // 添加ETH和代币的流动性
    // 参数含义同上，但使用ETH作为其中一种代币
    function addLiquidityETH(
        address token,
        uint amountTokenDesired,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline
    ) external payable returns (uint amountToken, uint amountETH, uint function addLiquidityETH(
        address token,
        uint amountTokenDesired,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline
    ) external payable returinterface IUniswapV2Router02 {
        // 获取工厂合约地址
        // @return 工厂合约地址
        function factory() external pure returns (address);

        // 获取WETH代币地址
        // @return WETH代币合约地址
        function WETH() external pure returns (address);

        // 添加流动性
        // @param tokenA 第一个代币的地址
        // @param tokenB 第二个代币的地址
        // @param amountADesired tokenA的期望数量
        // @param amountBDesired tokenB的期望数量
        // @param amountAMin tokenA的最小可接受数量
        // @param amountBMin tokenB的最小可接受数量
        // @param to 接收LP代币的地址
        // @param deadline 交易截止时间戳
        // @return amountA 实际添加的tokenA数量
        // @return amountB 实际添加的tokenB数量
        // @return liquidity 获得的LP代币数量
        function addLiquidity(
            address tokenA,
            address tokenB,
            uint amountADesired,
            uint amountBDesired,
            uint amountAMin,
            uint amountBMin,
            address to,
            uint deadline
        ) external returns (uint amountA, uint amountB, uint liquidity);

        // 添加ETH和代币的流动性
        // @param token 代币地址
        // @param amountTokenDesired 代币的期望数量
        // @param amountTokenMin 代币的最小可接受数量
        // @param amountETHMin ETH的最小可接受数量
        // @param to 接收LP代币的地址
        // @param deadline 交易截止时间戳
        // @return amountToken 实际添加的代币数量
        // @return amountETH 实际添加的ETH数量
        // @return liquidity 获得的LP代币数量
        function addLiquidityETH(
            address token,
            uint amountTokenDesired,
            uint amountTokenMin,
            uint amountETHMin,
            address to,
            uint deadline
        ) external payable returns (uint amountToken, uint amountETH, uint liquidity);

        // 移除流动性
        // @param tokenA 第一个代币的地址
        // @param tokenB 第二个代币的地址
        // @param liquidity 要移除的LP代币数量
        // @param amountAMin tokenA的最小可接受数量
        // @param amountBMin tokenB的最小可接受数量
        // @param to 接收代币的地址
        // @param deadline 交易截止时间戳
        // @return amountA 获得的tokenA数量
        // @return amountB 获得的tokenB数量
        function removeLiquidity(
            address tokenA,
            address tokenB,
            uint liquidity,
            uint amountAMin,
            uint amountBMin,
            address to,
            uint deadline
        ) external returns (uint amountA, uint amountB);

        // 移除ETH和代币的流动性
        // @param token 代币地址
        // @param liquidity 要移除的LP代币数量
        // @param amountTokenMin 代币的最小可接受数量
        // @param amountETHMin ETH的最小可接受数量
        // @param to 接收资产的地址
        // @param deadline 交易截止时间戳
        // @return amountToken 获得的代币数量
        // @return amountETH 获得的ETH数量
        function removeLiquidityETH(
            address token,
            uint liquidity,
            uint amountTokenMin,
            uint amountETHMin,
            address to,
            uint deadline
        ) external returns (uint amountToken, uint amountETH);

        // 使用许可证移除流动性
        // @param tokenA 第一个代币的地址
        // @param tokenB 第二个代币的地址
        // @param liquidity 要移除的LP代币数量
        // @param amountAMin tokenA的最小可接受数量
        // @param amountBMin tokenB的最小可接受数量
        // @param to 接收资产的地址
        // @param deadline 交易截止时间戳
        // @param approveMax 是否批准最大数量
        // @param v 签名的v值
        // @param r 签名的r值
        // @param s 签名的s值
        // @return amountA 获得的tokenA数量
        // @return amountB 获得的tokenB数量
        function removeLiquidityWithPermit(
            address tokenA,
            address tokenB,
            uint liquidity,
            uint amountAMin,
            uint amountBMin,
            address to,
            uint deadline,
            bool approveMax, uint8 v, bytes32 r, bytes32 s
        ) external returns (uint amountA, uint amountB);

        // 使用许可证移除ETH流动性
        // @param token 代币地址
        // @param liquidity 要移除的LP代币数量
        // @param amountTokenMin 代币的最小可接受数量
        // @param amountETHMin ETH的最小可接受数量
        // @param to 接收资产的地址
        // @param deadline 交易截止时间戳
        // @param approveMax 是否批准最大数量
        // @param v 签名的v值
        // @param r 签名的r值
        // @param s 签名的s值
        // @return amountToken 获得的代币数量
        // @return amountETH 获得的ETH数量
        function removeLiquidityETHWithPermit(
            address token,
            uint liquidity,
            uint amountTokenMin,
            uint amountETHMin,
            address to,
            uint deadline,
            bool approveMax, uint8 v, bytes32 r, bytes32 s
        ) external returns (uint amountToken, uint amountETH);

        // 按精确输入数量交换代币
        // @param amountIn 输入代币的确切数量
        // @param amountOutMin 输出代币的最小数量
        // @param path 交易路径地址数组
        // @param to 接收代币的地址
        // @param deadline 交易截止时间戳
        // @return amounts 交易路径中每个代币的数量
        function swapExactTokensForTokens(
            uint amountIn,
            uint amountOutMin,
            address[] calldata path,
            address to,
            uint deadline
        ) external returns (uint[] memory amounts);

        // 按精确输出数量交换代币
        // @param amountOut 期望获得的代币数量
        // @param amountInMax 最大输入代币数量
        // @param path 交易路径地址数组
        // @param to 接收代币的地址
        // @param deadline 交易截止时间戳
        // @return amounts 交易路径中每个代币的数量
        function swapTokensForExactTokens(
            uint amountOut,
            uint amountInMax,
            address[] calldata path,
            address to,
            uint deadline
        ) external returns (uint[] memory amounts);

        // 用精确的ETH数量换取代币
        // @param amountOutMin 输出代币的最小数量
        // @param path 交易路径地址数组
        // @param to 接收代币的地址
        // @param deadline 交易截止时间戳
        // @return amounts 交易路径中每个代币的数量
        function swapExactETHForTokens(
            uint amountOutMin,
            address[] calldata path,
            address to,
            uint deadline
        ) external payable returns (uint[] memory amounts);

        // 用代币换取精确的ETH数量
        // @param amountOut 期望获得的ETH数量
        // @param amountInMax 最大输入代币数量
        // @param path 交易路径地址数组
        // @param to 接收ETH的地址
        // @param deadline 交易截止时间戳
        // @return amounts 交易路径中每个代币的数量
        function swapTokensForExactETH(
            uint amountOut,
            uint amountInMax,
            address[] calldata path,
            address to,
            uint deadline
        ) external returns (uint[] memory amounts);

        // 用精确的代币数量换取ETH
        // @param amountIn 输入代币的确切数量
        // @param amountOutMin 输出ETH的最小数量
        // @param path 交易路径地址数组
        // @param to 接收ETH的地址
        // @param deadline 交易截止时间戳
        // @return amounts 交易路径中每个资产的数量
        function swapExactTokensForETH(
            uint amountIn,
            uint amountOutMin,
            address[] calldata path,
            address to,
            uint deadline
        ) external returns (uint[] memory amounts);

        // 用ETH换取精确数量的代币
        // @param amountOut 期望获得的代币数量
        // @param path 交易路径地址数组
        // @param to 接收代币的地址
        // @param deadline 交易截止时间戳
        // @return amounts 交易路径中每个资产的数量
        function swapETHForExactTokens(
            uint amountOut,
            address[] calldata path,
            address to,
            uint deadline
        ) external payable returns (uint[] memory amounts);

        // 计算代币兑换比率
        // @param amountA 输入代币数量
        // @param reserveA 输入代币储备量
        // @param reserveB 输出代币储备量
        // @return amountB 输出代币数量
        function quote(
            uint amountA,
            uint reserveA,
            uint reserveB
        ) external pure returns (uint amountB);

        // 计算代币兑换输出量
        // @param amountIn 输入代币数量
        // @param reserveIn 输入代币储备量
        // @param reserveOut 输出代币储备量
        // @return amountOut 输出代币数量
        function getAmountOut(
            uint amountIn,
            uint reserveIn,
            uint reserveOut
        ) external pure returns (uint amountOut);

        // 计算代币兑换输入量
        // @param amountOut 期望输出的代币数量
        // @param reserveIn 输入代币储备量
        // @param reserveOut 输出代币储备量
        // @return amountIn 需要的输入代币数量
        function getAmountIn(
            uint amountOut,
            uint reserveIn,
            uint reserveOut
        ) external pure returns (uint amountIn);

        // 计算多跳兑换的输出量
        // @param amountIn 输入代币数量
        // @param path 交易路径地址数组
        // @return amounts 交易路径中每个代币的数量
        function getAmountsOut(
            uint amountIn,
            address[] calldata path
        ) external view returns (uint[] memory amounts);

        // 计算多跳兑换的输入量
        // @param amountOut 期望输出的代币数量
        // @param path 交易路径地址数组
        // @return amounts 交易路径中每个代币的数量
        function getAmountsIn(
            uint amountOut,
            address[] calldata path
        ) external view returns (uint[] memory amounts);

        // 支持转账收费代币的移除ETH流动性
        // @param token 代币地址
        // @param liquidity 要移除的LP代币数量
        // @param amountTokenMin 代币的最小可接受数量
        // @param amountETHMin ETH的最小可接受数量
        // @param to 接收资产的地址
        // @param deadline 交易截止时间戳
        // @return amountETH 获得的ETH数量
        function removeLiquidityETHSupportingFeeOnTransferTokens(
            address token,
            uint liquidity,
            uint amountTokenMin,
            uint amountETHMin,
            address to,
            uint deadline
        ) external returns (uint amountETH);

        // 支持转账收费代币的带许可证移除ETH流动性
        // @param token 代币地址
        // @param liquidity 要移除的LP代币数量
        // @param amountTokenMin 代币的最小可接受数量
        // @param amountETHMin ETH的最小可接受数量
        // @param to 接收资产的地址
        // @param deadline 交易截止时间戳
        // @param approveMax 是否批准最大数量
        // @param v 签名的v值
        // @param r 签名的r值
        // @param s 签名的s值
        // @return amountETH 获得的ETH数量
        function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(
            address token,
            uint liquidity,
            uint amountTokenMin,
            uint amountETHMin,
            address to,
            uint deadline,
            bool approveMax, uint8 v, bytes32 r, bytes32 s
        ) external returns (uint amountETH);

        // 支持转账收费的代币兑换
        // @param amountIn 输入代币的确切数量
        // @param amountOutMin 输出代币的最小数量
        // @param path 交易路径地址数组
        // @param to 接收代币的地址
        // @param deadline 交易截止时间戳
        function swapExactTokensForTokensSupportingFeeOnTransferTokens(
            uint amountIn,
            uint amountOutMin,
            address[] calldata path,
            address to,
            uint deadline
        ) external;

        // 支持转账收费的ETH兑换代币
        // @param amountOutMin 输出代币的最小数量
        // @param path 交易路径地址数组
        // @param to 接收代币的地址
        // @param deadline 交易截止时间戳
        function swapExactETHForTokensSupportingFeeOnTransferTokens(
            uint amountOutMin,
            address[] calldata path,
            address to,
            uint deadline
        ) external payable;

        // 支持转账收费的代币兑换ETH
        // @param amountIn 输入代币的确切数量
        // @param amountOutMin 输出ETH的最小数量
        // @param path 交易路径地址数组
        // @param to 接收ETH的地址
        // @param deadline 交易截止时间戳
        function swapExactTokensForETHSupportingFeeOnTransferTokens(
            uint amountIn,
            uint amountOutMin,
            address[] calldata path,
            address to,
            uint deadline
        ) external;
    }ns (
        uint amountToken,  // 实际添加的代币数量
        uint amountETH,    // 实际添加的ETH数量
        uint liquidity     // 获得的LP代币数量
    ););

    // 移除流动性
    // liquidity是要移除的LP代币数量
    // amount[A/B]Min是最小接受的代币数量
    function removeLiquidity(
        address tokenA,
        address tokenB,
        uint liquidity,
        uint amountAMin,
        uint amountBMin,
        address to,
        uint deadline
    ) external returns (uint amountA, uint amountB);

    // 移除ETH和代币的流动性
    function removeLiquidityETH(
        address token,
        uint liquidity,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline
    ) external returns (uint amountToken, uint amountETH);

    // 使用许可证移除流动性
    // approveMax, v, r, s 是签名相关参数
    function removeLiquidityWithPermit(
        address tokenA,
        address tokenB,
        uint liquidity,
        uint amountAMin,
        uint amountBMin,
        address to,
        uint deadline,
        bool approveMax, uint8 v, bytes32 r, bytes32 s
    ) external returns (uint amountA, uint amountB);

    // 使用许可证移除ETH流动性
    function removeLiquidityETHWithPermit(
        address token,
        uint liquidity,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline,
        bool approveMax, uint8 v, bytes32 r, bytes32 s
    ) external returns (uint amountToken, uint amountETH);

    // 按精确输入数量交换代币
    // path是代币兑换路径
    function swapExactTokensForTokens(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external returns (uint[] memory amounts);

    // 按精确输出数量交换代币
    function swapTokensForExactTokens(
        uint amountOut,
        uint amountInMax,
        address[] calldata path,
        address to,
        uint deadline
    ) external returns (uint[] memory amounts);

    // 用精确的ETH数量换取代币
    function swapExactETHForTokens(uint amountOutMin, address[] calldata path, address to, uint deadline)
        external
        payable
        returns (uint[] memory amounts);

    // 用代币换取精确的ETH数量
    function swapTokensForExactETH(uint amountOut, uint amountInMax, address[] calldata path, address to, uint deadline)
        external
        returns (uint[] memory amounts);

    // 用精确的代币数量换取ETH
    function swapExactTokensForETH(uint amountIn, uint amountOutMin, address[] calldata path, address to, uint deadline)
        external
        returns (uint[] memory amounts);

    // 用ETH换取精确数量的代币
    function swapETHForExactTokens(uint amountOut, address[] calldata path, address to, uint deadline)
        external
        payable
        returns (uint[] memory amounts);

    // 计算代币兑换比率
    function quote(uint amountA, uint reserveA, uint reserveB) external pure returns (uint amountB);
    // 计算代币兑换输出量
    function getAmountOut(uint amountIn, uint reserveIn, uint reserveOut) external pure returns (uint amountOut);
    // 计算代币兑换输入量
    function getAmountIn(uint amountOut, uint reserveIn, uint reserveOut) external pure returns (uint amountIn);
    // 计算多跳兑换的输出量
    function getAmountsOut(uint amountIn, address[] calldata path) external view returns (uint[] memory amounts);
    // 计算多跳兑换的输入量
    function getAmountsIn(uint amountOut, address[] calldata path) external view returns (uint[] memory amounts);

    // 支持转账收费代币的移除ETH流动性
    function removeLiquidityETHSupportingFeeOnTransferTokens(
        address token,
        uint liquidity,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline
    ) external returns (uint amountETH);

    // 支持转账收费代币的带许可证移除ETH流动性
    function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(
        address token,
        uint liquidity,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline,
        bool approveMax, uint8 v, bytes32 r, bytes32 s
    ) external returns (uint amountETH);

    // 支持转账收费的代币兑换
    function swapExactTokensForTokensSupportingFeeOnTransferTokens(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external;

    // 支持转账收费的ETH兑换代币
    function swapExactETHForTokensSupportingFeeOnTransferTokens(
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external payable;

    // 支持转账收费的代币兑换ETH
    function swapExactTokensForETHSupportingFeeOnTransferTokens(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external;
}// SPDX-License-Identifier: MIT

pragma solidity 0.6.12;

interface IUniswapV2Router02 {
    function factory() external pure returns (address);
    function WETH() external pure returns (address);

    function addLiquidity(
        address tokenA,
        address tokenB,
        uint amountADesired,
        uint amountBDesired,
        uint amountAMin,
        uint amountBMin,
        address to,
        uint deadline
    ) external returns (uint amountA, uint amountB, uint liquidity);
    function addLiquidityETH(
        address token,
        uint amountTokenDesired,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline
    ) external payable returns (uint amountToken, uint amountETH, uint liquidity);
    function removeLiquidity(
        address tokenA,
        address tokenB,
        uint liquidity,
        uint amountAMin,
        uint amountBMin,
        address to,
        uint deadline
    ) external returns (uint amountA, uint amountB);
    function removeLiquidityETH(
        address token,
        uint liquidity,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline
    ) external returns (uint amountToken, uint amountETH);
    function removeLiquidityWithPermit(
        address tokenA,
        address tokenB,
        uint liquidity,
        uint amountAMin,
        uint amountBMin,
        address to,
        uint deadline,
        bool approveMax, uint8 v, bytes32 r, bytes32 s
    ) external returns (uint amountA, uint amountB);
    function removeLiquidityETHWithPermit(
        address token,
        uint liquidity,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline,
        bool approveMax, uint8 v, bytes32 r, bytes32 s
    ) external returns (uint amountToken, uint amountETH);
    function swapExactTokensForTokens(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external returns (uint[] memory amounts);
    function swapTokensForExactTokens(
        uint amountOut,
        uint amountInMax,
        address[] calldata path,
        address to,
        uint deadline
    ) external returns (uint[] memory amounts);
    function swapExactETHForTokens(uint amountOutMin, address[] calldata path, address to, uint deadline)
        external
        payable
        returns (uint[] memory amounts);
    function swapTokensForExactETH(uint amountOut, uint amountInMax, address[] calldata path, address to, uint deadline)
        external
        returns (uint[] memory amounts);
    function swapExactTokensForETH(uint amountIn, uint amountOutMin, address[] calldata path, address to, uint deadline)
        external
        returns (uint[] memory amounts);
    function swapETHForExactTokens(uint amountOut, address[] calldata path, address to, uint deadline)
        external
        payable
        returns (uint[] memory amounts);

    function quote(uint amountA, uint reserveA, uint reserveB) external pure returns (uint amountB);
    function getAmountOut(uint amountIn, uint reserveIn, uint reserveOut) external pure returns (uint amountOut);
    function getAmountIn(uint amountOut, uint reserveIn, uint reserveOut) external pure returns (uint amountIn);
    function getAmountsOut(uint amountIn, address[] calldata path) external view returns (uint[] memory amounts);
    function getAmountsIn(uint amountOut, address[] calldata path) external view returns (uint[] memory amounts);

    function removeLiquidityETHSupportingFeeOnTransferTokens(
        address token,
        uint liquidity,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline
    ) external returns (uint amountETH);
    function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(
        address token,
        uint liquidity,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline,
        bool approveMax, uint8 v, bytes32 r, bytes32 s
    ) external returns (uint amountETH);

    function swapExactTokensForTokensSupportingFeeOnTransferTokens(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external;
    function swapExactETHForTokensSupportingFeeOnTransferTokens(
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external payable;
    function swapExactTokensForETHSupportingFeeOnTransferTokens(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external;
}
