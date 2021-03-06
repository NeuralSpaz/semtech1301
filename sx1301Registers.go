package sx1301

type loraRegister struct {
	page         int8
	address      uint8
	offset       uint8
	signed       bool // 2 complement
	length       uint8
	readonly     bool
	defaultValue int32
}

// map of sx1301 registers
var Registers = map[string]loraRegister{
	"LGW_PAGE_REG":                            {-1, 0, 0, false, 2, false, 0},     /* PAGE_REG */
	"LGW_SOFT_RESET":                          {-1, 0, 7, false, 1, false, 0},     /* SOFT_RESET */
	"LGW_VERSION":                             {-1, 1, 0, false, 8, true, 103},    /* VERSION */
	"LGW_RX_DATA_BUF_ADDR":                    {-1, 2, 0, false, 16, false, 0},    /* RX_DATA_BUF_ADDR */
	"LGW_RX_DATA_BUF_DATA":                    {-1, 4, 0, false, 8, false, 0},     /* RX_DATA_BUF_DATA */
	"LGW_TX_DATA_BUF_ADDR":                    {-1, 5, 0, false, 8, false, 0},     /* TX_DATA_BUF_ADDR */
	"LGW_TX_DATA_BUF_DATA":                    {-1, 6, 0, false, 8, false, 0},     /* TX_DATA_BUF_DATA */
	"LGW_CAPTURE_RAM_ADDR":                    {-1, 7, 0, false, 8, false, 0},     /* CAPTURE_RAM_ADDR */
	"LGW_CAPTURE_RAM_DATA":                    {-1, 8, 0, false, 8, true, 0},      /* CAPTURE_RAM_DATA */
	"LGW_MCU_PROM_ADDR":                       {-1, 9, 0, false, 8, false, 0},     /* MCU_PROM_ADDR */
	"LGW_MCU_PROM_DATA":                       {-1, 10, 0, false, 8, false, 0},    /* MCU_PROM_DATA */
	"LGW_RX_PACKET_DATA_FIFO_NUM_STORED":      {-1, 11, 0, false, 8, false, 0},    /* RX_PACKET_DATA_FIFO_NUM_STORED */
	"LGW_RX_PACKET_DATA_FIFO_ADDR_POINTER":    {-1, 12, 0, false, 16, true, 0},    /* RX_PACKET_DATA_FIFO_ADDR_POINTER */
	"LGW_RX_PACKET_DATA_FIFO_STATUS":          {-1, 14, 0, false, 8, true, 0},     /* RX_PACKET_DATA_FIFO_STATUS */
	"LGW_RX_PACKET_DATA_FIFO_PAYLOAD_SIZE":    {-1, 15, 0, false, 8, true, 0},     /* RX_PACKET_DATA_FIFO_PAYLOAD_SIZE */
	"LGW_MBWSSF_MODEM_ENABLE":                 {-1, 16, 0, false, 1, false, 0},    /* MBWSSF_MODEM_ENABLE */
	"LGW_CONCENTRATOR_MODEM_ENABLE":           {-1, 16, 1, false, 1, false, 0},    /* CONCENTRATOR_MODEM_ENABLE */
	"LGW_FSK_MODEM_ENABLE":                    {-1, 16, 2, false, 1, false, 0},    /* FSK_MODEM_ENABLE */
	"LGW_GLOBAL_EN":                           {-1, 16, 3, false, 1, false, 0},    /* GLOBAL_EN */
	"LGW_CLK32M_EN":                           {-1, 17, 0, false, 1, false, 1},    /* CLK32M_EN */
	"LGW_CLKHS_EN":                            {-1, 17, 1, false, 1, false, 1},    /* CLKHS_EN */
	"LGW_START_BIST0":                         {-1, 18, 0, false, 1, false, 0},    /* START_BIST0 */
	"LGW_START_BIST1":                         {-1, 18, 1, false, 1, false, 0},    /* START_BIST1 */
	"LGW_CLEAR_BIST0":                         {-1, 18, 2, false, 1, false, 0},    /* CLEAR_BIST0 */
	"LGW_CLEAR_BIST1":                         {-1, 18, 3, false, 1, false, 0},    /* CLEAR_BIST1 */
	"LGW_BIST0_FINISHED":                      {-1, 19, 0, false, 1, true, 0},     /* BIST0_FINISHED */
	"LGW_BIST1_FINISHED":                      {-1, 19, 1, false, 1, true, 0},     /* BIST1_FINISHED */
	"LGW_MCU_AGC_PROG_RAM_BIST_STATUS":        {-1, 20, 0, false, 1, true, 0},     /* MCU_AGC_PROG_RAM_BIST_STATUS */
	"LGW_MCU_ARB_PROG_RAM_BIST_STATUS":        {-1, 20, 1, false, 1, true, 0},     /* MCU_ARB_PROG_RAM_BIST_STATUS */
	"LGW_CAPTURE_RAM_BIST_STATUS":             {-1, 20, 2, false, 1, true, 0},     /* CAPTURE_RAM_BIST_STATUS */
	"LGW_CHAN_FIR_RAM0_BIST_STATUS":           {-1, 20, 3, false, 1, true, 0},     /* CHAN_FIR_RAM0_BIST_STATUS */
	"LGW_CHAN_FIR_RAM1_BIST_STATUS":           {-1, 20, 4, false, 1, true, 0},     /* CHAN_FIR_RAM1_BIST_STATUS */
	"LGW_CORR0_RAM_BIST_STATUS":               {-1, 21, 0, false, 1, true, 0},     /* CORR0_RAM_BIST_STATUS */
	"LGW_CORR1_RAM_BIST_STATUS":               {-1, 21, 1, false, 1, true, 0},     /* CORR1_RAM_BIST_STATUS */
	"LGW_CORR2_RAM_BIST_STATUS":               {-1, 21, 2, false, 1, true, 0},     /* CORR2_RAM_BIST_STATUS */
	"LGW_CORR3_RAM_BIST_STATUS":               {-1, 21, 3, false, 1, true, 0},     /* CORR3_RAM_BIST_STATUS */
	"LGW_CORR4_RAM_BIST_STATUS":               {-1, 21, 4, false, 1, true, 0},     /* CORR4_RAM_BIST_STATUS */
	"LGW_CORR5_RAM_BIST_STATUS":               {-1, 21, 5, false, 1, true, 0},     /* CORR5_RAM_BIST_STATUS */
	"LGW_CORR6_RAM_BIST_STATUS":               {-1, 21, 6, false, 1, true, 0},     /* CORR6_RAM_BIST_STATUS */
	"LGW_CORR7_RAM_BIST_STATUS":               {-1, 21, 7, false, 1, true, 0},     /* CORR7_RAM_BIST_STATUS */
	"LGW_MODEM0_RAM0_BIST_STATUS":             {-1, 22, 0, false, 1, true, 0},     /* MODEM0_RAM0_BIST_STATUS */
	"LGW_MODEM1_RAM0_BIST_STATUS":             {-1, 22, 1, false, 1, true, 0},     /* MODEM1_RAM0_BIST_STATUS */
	"LGW_MODEM2_RAM0_BIST_STATUS":             {-1, 22, 2, false, 1, true, 0},     /* MODEM2_RAM0_BIST_STATUS */
	"LGW_MODEM3_RAM0_BIST_STATUS":             {-1, 22, 3, false, 1, true, 0},     /* MODEM3_RAM0_BIST_STATUS */
	"LGW_MODEM4_RAM0_BIST_STATUS":             {-1, 22, 4, false, 1, true, 0},     /* MODEM4_RAM0_BIST_STATUS */
	"LGW_MODEM5_RAM0_BIST_STATUS":             {-1, 22, 5, false, 1, true, 0},     /* MODEM5_RAM0_BIST_STATUS */
	"LGW_MODEM6_RAM0_BIST_STATUS":             {-1, 22, 6, false, 1, true, 0},     /* MODEM6_RAM0_BIST_STATUS */
	"LGW_MODEM7_RAM0_BIST_STATUS":             {-1, 22, 7, false, 1, true, 0},     /* MODEM7_RAM0_BIST_STATUS */
	"LGW_MODEM0_RAM1_BIST_STATUS":             {-1, 23, 0, false, 1, true, 0},     /* MODEM0_RAM1_BIST_STATUS */
	"LGW_MODEM1_RAM1_BIST_STATUS":             {-1, 23, 1, false, 1, true, 0},     /* MODEM1_RAM1_BIST_STATUS */
	"LGW_MODEM2_RAM1_BIST_STATUS":             {-1, 23, 2, false, 1, true, 0},     /* MODEM2_RAM1_BIST_STATUS */
	"LGW_MODEM3_RAM1_BIST_STATUS":             {-1, 23, 3, false, 1, true, 0},     /* MODEM3_RAM1_BIST_STATUS */
	"LGW_MODEM4_RAM1_BIST_STATUS":             {-1, 23, 4, false, 1, true, 0},     /* MODEM4_RAM1_BIST_STATUS */
	"LGW_MODEM5_RAM1_BIST_STATUS":             {-1, 23, 5, false, 1, true, 0},     /* MODEM5_RAM1_BIST_STATUS */
	"LGW_MODEM6_RAM1_BIST_STATUS":             {-1, 23, 6, false, 1, true, 0},     /* MODEM6_RAM1_BIST_STATUS */
	"LGW_MODEM7_RAM1_BIST_STATUS":             {-1, 23, 7, false, 1, true, 0},     /* MODEM7_RAM1_BIST_STATUS */
	"LGW_MODEM0_RAM2_BIST_STATUS":             {-1, 24, 0, false, 1, true, 0},     /* MODEM0_RAM2_BIST_STATUS */
	"LGW_MODEM1_RAM2_BIST_STATUS":             {-1, 24, 1, false, 1, true, 0},     /* MODEM1_RAM2_BIST_STATUS */
	"LGW_MODEM2_RAM2_BIST_STATUS":             {-1, 24, 2, false, 1, true, 0},     /* MODEM2_RAM2_BIST_STATUS */
	"LGW_MODEM3_RAM2_BIST_STATUS":             {-1, 24, 3, false, 1, true, 0},     /* MODEM3_RAM2_BIST_STATUS */
	"LGW_MODEM4_RAM2_BIST_STATUS":             {-1, 24, 4, false, 1, true, 0},     /* MODEM4_RAM2_BIST_STATUS */
	"LGW_MODEM5_RAM2_BIST_STATUS":             {-1, 24, 5, false, 1, true, 0},     /* MODEM5_RAM2_BIST_STATUS */
	"LGW_MODEM6_RAM2_BIST_STATUS":             {-1, 24, 6, false, 1, true, 0},     /* MODEM6_RAM2_BIST_STATUS */
	"LGW_MODEM7_RAM2_BIST_STATUS":             {-1, 24, 7, false, 1, true, 0},     /* MODEM7_RAM2_BIST_STATUS */
	"LGW_MODEM_MBWSSF_RAM0_BIST_STATUS":       {-1, 25, 0, false, 1, true, 0},     /* MODEM_MBWSSF_RAM0_BIST_STATUS */
	"LGW_MODEM_MBWSSF_RAM1_BIST_STATUS":       {-1, 25, 1, false, 1, true, 0},     /* MODEM_MBWSSF_RAM1_BIST_STATUS */
	"LGW_MODEM_MBWSSF_RAM2_BIST_STATUS":       {-1, 25, 2, false, 1, true, 0},     /* MODEM_MBWSSF_RAM2_BIST_STATUS */
	"LGW_MCU_AGC_DATA_RAM_BIST0_STATUS":       {-1, 26, 0, false, 1, true, 0},     /* MCU_AGC_DATA_RAM_BIST0_STATUS */
	"LGW_MCU_AGC_DATA_RAM_BIST1_STATUS":       {-1, 26, 1, false, 1, true, 0},     /* MCU_AGC_DATA_RAM_BIST1_STATUS */
	"LGW_MCU_ARB_DATA_RAM_BIST0_STATUS":       {-1, 26, 2, false, 1, true, 0},     /* MCU_ARB_DATA_RAM_BIST0_STATUS */
	"LGW_MCU_ARB_DATA_RAM_BIST1_STATUS":       {-1, 26, 3, false, 1, true, 0},     /* MCU_ARB_DATA_RAM_BIST1_STATUS */
	"LGW_TX_TOP_RAM_BIST0_STATUS":             {-1, 26, 4, false, 1, true, 0},     /* TX_TOP_RAM_BIST0_STATUS */
	"LGW_TX_TOP_RAM_BIST1_STATUS":             {-1, 26, 5, false, 1, true, 0},     /* TX_TOP_RAM_BIST1_STATUS */
	"LGW_DATA_MNGT_RAM_BIST0_STATUS":          {-1, 26, 6, false, 1, true, 0},     /* DATA_MNGT_RAM_BIST0_STATUS */
	"LGW_DATA_MNGT_RAM_BIST1_STATUS":          {-1, 26, 7, false, 1, true, 0},     /* DATA_MNGT_RAM_BIST1_STATUS */
	"LGW_GPIO_SELECT_INPUT":                   {-1, 27, 0, false, 4, false, 0},    /* GPIO_SELECT_INPUT */
	"LGW_GPIO_SELECT_OUTPUT":                  {-1, 28, 0, false, 4, false, 0},    /* GPIO_SELECT_OUTPUT */
	"LGW_GPIO_MODE":                           {-1, 29, 0, false, 5, false, 0},    /* GPIO_MODE */
	"LGW_GPIO_PIN_REG_IN":                     {-1, 30, 0, false, 5, true, 0},     /* GPIO_PIN_REG_IN */
	"LGW_GPIO_PIN_REG_OUT":                    {-1, 31, 0, false, 5, false, 0},    /* GPIO_PIN_REG_OUT */
	"LGW_MCU_AGC_STATUS":                      {-1, 32, 0, false, 8, true, 0},     /* MCU_AGC_STATUS */
	"LGW_MCU_ARB_STATUS":                      {-1, 125, 0, false, 8, true, 0},    /* MCU_ARB_STATUS */
	"LGW_CHIP_ID":                             {-1, 126, 0, false, 8, true, 1},    /* CHIP_ID */
	"LGW_EMERGENCY_FORCE_HOST_CTRL":           {-1, 127, 0, false, 1, false, 1},   /* EMERGENCY_FORCE_HOST_CTRL */
	"LGW_RX_INVERT_IQ":                        {0, 33, 0, false, 1, false, 0},     /* RX_INVERT_IQ */
	"LGW_MODEM_INVERT_IQ":                     {0, 33, 1, false, 1, false, 1},     /* MODEM_INVERT_IQ */
	"LGW_MBWSSF_MODEM_INVERT_IQ":              {0, 33, 2, false, 1, false, 0},     /* MBWSSF_MODEM_INVERT_IQ */
	"LGW_RX_EDGE_SELECT":                      {0, 33, 3, false, 1, false, 0},     /* RX_EDGE_SELECT */
	"LGW_MISC_RADIO_EN":                       {0, 33, 4, false, 1, false, 0},     /* MISC_RADIO_EN */
	"LGW_FSK_MODEM_INVERT_IQ":                 {0, 33, 5, false, 1, false, 0},     /* FSK_MODEM_INVERT_IQ */
	"LGW_FILTER_GAIN":                         {0, 34, 0, false, 4, false, 7},     /* FILTER_GAIN */
	"LGW_RADIO_SELECT":                        {0, 35, 0, false, 8, false, 240},   /* RADIO_SELECT */
	"LGW_IF_FREQ_0":                           {0, 36, 0, true, 13, false, -384},  /* IF_FREQ_0 */
	"LGW_IF_FREQ_1":                           {0, 38, 0, true, 13, false, -128},  /* IF_FREQ_1 */
	"LGW_IF_FREQ_2":                           {0, 40, 0, true, 13, false, 128},   /* IF_FREQ_2 */
	"LGW_IF_FREQ_3":                           {0, 42, 0, true, 13, false, 384},   /* IF_FREQ_3 */
	"LGW_IF_FREQ_4":                           {0, 44, 0, true, 13, false, -384},  /* IF_FREQ_4 */
	"LGW_IF_FREQ_5":                           {0, 46, 0, true, 13, false, -128},  /* IF_FREQ_5 */
	"LGW_IF_FREQ_6":                           {0, 48, 0, true, 13, false, 128},   /* IF_FREQ_6 */
	"LGW_IF_FREQ_7":                           {0, 50, 0, true, 13, false, 384},   /* IF_FREQ_7 */
	"LGW_IF_FREQ_8":                           {0, 52, 0, true, 13, false, 0},     /* IF_FREQ_8 */
	"LGW_IF_FREQ_9":                           {0, 54, 0, true, 13, false, 0},     /* IF_FREQ_9 */
	"LGW_CHANN_OVERRIDE_AGC_GAIN":             {0, 64, 0, false, 1, false, 0},     /* CHANN_OVERRIDE_AGC_GAIN */
	"LGW_CHANN_AGC_GAIN":                      {0, 64, 1, false, 4, false, 7},     /* CHANN_AGC_GAIN */
	"LGW_CORR0_DETECT_EN":                     {0, 65, 0, false, 7, false, 0},     /* CORR0_DETECT_EN */
	"LGW_CORR1_DETECT_EN":                     {0, 66, 0, false, 7, false, 0},     /* CORR1_DETECT_EN */
	"LGW_CORR2_DETECT_EN":                     {0, 67, 0, false, 7, false, 0},     /* CORR2_DETECT_EN */
	"LGW_CORR3_DETECT_EN":                     {0, 68, 0, false, 7, false, 0},     /* CORR3_DETECT_EN */
	"LGW_CORR4_DETECT_EN":                     {0, 69, 0, false, 7, false, 0},     /* CORR4_DETECT_EN */
	"LGW_CORR5_DETECT_EN":                     {0, 70, 0, false, 7, false, 0},     /* CORR5_DETECT_EN */
	"LGW_CORR6_DETECT_EN":                     {0, 71, 0, false, 7, false, 0},     /* CORR6_DETECT_EN */
	"LGW_CORR7_DETECT_EN":                     {0, 72, 0, false, 7, false, 0},     /* CORR7_DETECT_EN */
	"LGW_CORR_SAME_PEAKS_OPTION_SF6":          {0, 73, 0, false, 1, false, 0},     /* CORR_SAME_PEAKS_OPTION_SF6 */
	"LGW_CORR_SAME_PEAKS_OPTION_SF7":          {0, 73, 1, false, 1, false, 1},     /* CORR_SAME_PEAKS_OPTION_SF7 */
	"LGW_CORR_SAME_PEAKS_OPTION_SF8":          {0, 73, 2, false, 1, false, 1},     /* CORR_SAME_PEAKS_OPTION_SF8 */
	"LGW_CORR_SAME_PEAKS_OPTION_SF9":          {0, 73, 3, false, 1, false, 1},     /* CORR_SAME_PEAKS_OPTION_SF9 */
	"LGW_CORR_SAME_PEAKS_OPTION_SF10":         {0, 73, 4, false, 1, false, 1},     /* CORR_SAME_PEAKS_OPTION_SF10 */
	"LGW_CORR_SAME_PEAKS_OPTION_SF11":         {0, 73, 5, false, 1, false, 1},     /* CORR_SAME_PEAKS_OPTION_SF11 */
	"LGW_CORR_SAME_PEAKS_OPTION_SF12":         {0, 73, 6, false, 1, false, 1},     /* CORR_SAME_PEAKS_OPTION_SF12 */
	"LGW_CORR_SIG_NOISE_RATIO_SF6":            {0, 74, 0, false, 4, false, 4},     /* CORR_SIG_NOISE_RATIO_SF6 */
	"LGW_CORR_SIG_NOISE_RATIO_SF7":            {0, 74, 4, false, 4, false, 4},     /* CORR_SIG_NOISE_RATIO_SF7 */
	"LGW_CORR_SIG_NOISE_RATIO_SF8":            {0, 75, 0, false, 4, false, 4},     /* CORR_SIG_NOISE_RATIO_SF8 */
	"LGW_CORR_SIG_NOISE_RATIO_SF9":            {0, 75, 4, false, 4, false, 4},     /* CORR_SIG_NOISE_RATIO_SF9 */
	"LGW_CORR_SIG_NOISE_RATIO_SF10":           {0, 76, 0, false, 4, false, 4},     /* CORR_SIG_NOISE_RATIO_SF10 */
	"LGW_CORR_SIG_NOISE_RATIO_SF11":           {0, 76, 4, false, 4, false, 4},     /* CORR_SIG_NOISE_RATIO_SF11 */
	"LGW_CORR_SIG_NOISE_RATIO_SF12":           {0, 77, 0, false, 4, false, 4},     /* CORR_SIG_NOISE_RATIO_SF12 */
	"LGW_CORR_NUM_SAME_PEAK":                  {0, 78, 0, false, 4, false, 4},     /* CORR_NUM_SAME_PEAK */
	"LGW_CORR_MAC_GAIN":                       {0, 78, 4, false, 3, false, 5},     /* CORR_MAC_GAIN */
	"LGW_ADJUST_MODEM_START_OFFSET_RDX4":      {0, 81, 0, false, 12, false, 0},    /* ADJUST_MODEM_START_OFFSET_RDX4 */
	"LGW_ADJUST_MODEM_START_OFFSET_SF12_RDX4": {0, 83, 0, false, 12, false, 4092}, /* ADJUST_MODEM_START_OFFSET_SF12_RDX4 */
	"LGW_DBG_CORR_SELECT_SF":                  {0, 85, 0, false, 8, false, 7},     /* DBG_CORR_SELECT_SF */
	"LGW_DBG_CORR_SELECT_CHANNEL":             {0, 86, 0, false, 8, false, 0},     /* DBG_CORR_SELECT_CHANNEL */
	"LGW_DBG_DETECT_CPT":                      {0, 87, 0, false, 8, true, 0},      /* DBG_DETECT_CPT */
	"LGW_DBG_SYMB_CPT":                        {0, 88, 0, false, 8, true, 0},      /* DBG_SYMB_CPT */
	"LGW_CHIRP_INVERT_RX":                     {0, 89, 0, false, 1, false, 1},     /* CHIRP_INVERT_RX */
	"LGW_DC_NOTCH_EN":                         {0, 89, 1, false, 1, false, 1},     /* DC_NOTCH_EN */
	"LGW_IMPLICIT_CRC_EN":                     {0, 90, 0, false, 1, false, 0},     /* IMPLICIT_CRC_EN */
	"LGW_IMPLICIT_CODING_RATE":                {0, 90, 1, false, 3, false, 0},     /* IMPLICIT_CODING_RATE */
	"LGW_IMPLICIT_PAYLOAD_LENGHT":             {0, 91, 0, false, 8, false, 0},     /* IMPLICIT_PAYLOAD_LENGHT */
	"LGW_FREQ_TO_TIME_INVERT":                 {0, 92, 0, false, 8, false, 29},    /* FREQ_TO_TIME_INVERT */
	"LGW_FREQ_TO_TIME_DRIFT":                  {0, 93, 0, false, 6, false, 9},     /* FREQ_TO_TIME_DRIFT */
	"LGW_PAYLOAD_FINE_TIMING_GAIN":            {0, 94, 0, false, 2, false, 2},     /* PAYLOAD_FINE_TIMING_GAIN */
	"LGW_PREAMBLE_FINE_TIMING_GAIN":           {0, 94, 2, false, 2, false, 1},     /* PREAMBLE_FINE_TIMING_GAIN */
	"LGW_TRACKING_INTEGRAL":                   {0, 94, 4, false, 2, false, 0},     /* TRACKING_INTEGRAL */
	"LGW_FRAME_SYNCH_PEAK1_POS":               {0, 95, 0, false, 4, false, 1},     /* FRAME_SYNCH_PEAK1_POS */ //ALSO SOMETHING TO DO WITH SYNC WORD
	"LGW_FRAME_SYNCH_PEAK2_POS":               {0, 95, 4, false, 4, false, 2},     /* FRAME_SYNCH_PEAK2_POS */ //ALSO SOMETHING TO DO WITH SYNC WORD
	"LGW_PREAMBLE_SYMB1_NB":                   {0, 96, 0, false, 16, false, 10},   /* PREAMBLE_SYMB1_NB */
	"LGW_FRAME_SYNCH_GAIN":                    {0, 98, 0, false, 1, false, 1},     /* FRAME_SYNCH_GAIN */
	"LGW_SYNCH_DETECT_TH":                     {0, 98, 1, false, 1, false, 1},     /* SYNCH_DETECT_TH */
	"LGW_LLR_SCALE":                           {0, 99, 0, false, 4, false, 8},     /* LLR_SCALE */
	"LGW_SNR_AVG_CST":                         {0, 99, 4, false, 2, false, 2},     /* SNR_AVG_CST */
	"LGW_PPM_OFFSET":                          {0, 100, 0, false, 7, false, 0},    /* PPM_OFFSET */
	"LGW_MAX_PAYLOAD_LEN":                     {0, 101, 0, false, 8, false, 255},  /* MAX_PAYLOAD_LEN */
	"LGW_ONLY_CRC_EN":                         {0, 102, 0, false, 1, false, 1},    /* ONLY_CRC_EN */
	"LGW_ZERO_PAD":                            {0, 103, 0, false, 8, false, 0},    /* ZERO_PAD */
	"LGW_DEC_GAIN_OFFSET":                     {0, 104, 0, false, 4, false, 8},    /* DEC_GAIN_OFFSET */
	"LGW_CHAN_GAIN_OFFSET":                    {0, 104, 4, false, 4, false, 7},    /* CHAN_GAIN_OFFSET */
	"LGW_FORCE_HOST_RADIO_CTRL":               {0, 105, 1, false, 1, false, 1},    /* FORCE_HOST_RADIO_CTRL */
	"LGW_FORCE_HOST_FE_CTRL":                  {0, 105, 2, false, 1, false, 1},    /* FORCE_HOST_FE_CTRL */
	"LGW_FORCE_DEC_FILTER_GAIN":               {0, 105, 3, false, 1, false, 1},    /* FORCE_DEC_FILTER_GAIN */
	"LGW_MCU_RST_0":                           {0, 106, 0, false, 1, false, 1},    /* MCU_RST_0 */
	"LGW_MCU_RST_1":                           {0, 106, 1, false, 1, false, 1},    /* MCU_RST_1 */
	"LGW_MCU_SELECT_MUX_0":                    {0, 106, 2, false, 1, false, 0},    /* MCU_SELECT_MUX_0 */
	"LGW_MCU_SELECT_MUX_1":                    {0, 106, 3, false, 1, false, 0},    /* MCU_SELECT_MUX_1 */
	"LGW_MCU_CORRUPTION_DETECTED_0":           {0, 106, 4, false, 1, true, 0},     /* MCU_CORRUPTION_DETECTED_0 */
	"LGW_MCU_CORRUPTION_DETECTED_1":           {0, 106, 5, false, 1, true, 0},     /* MCU_CORRUPTION_DETECTED_1 */
	"LGW_MCU_SELECT_EDGE_0":                   {0, 106, 6, false, 1, false, 0},    /* MCU_SELECT_EDGE_0 */
	"LGW_MCU_SELECT_EDGE_1":                   {0, 106, 7, false, 1, false, 0},    /* MCU_SELECT_EDGE_1 */
	"LGW_CHANN_SELECT_RSSI":                   {0, 107, 0, false, 8, false, 1},    /* CHANN_SELECT_RSSI */
	"LGW_RSSI_BB_DEFAULT_VALUE":               {0, 108, 0, false, 8, false, 32},   /* RSSI_BB_DEFAULT_VALUE */
	"LGW_RSSI_DEC_DEFAULT_VALUE":              {0, 109, 0, false, 8, false, 100},  /* RSSI_DEC_DEFAULT_VALUE */
	"LGW_RSSI_CHANN_DEFAULT_VALUE":            {0, 110, 0, false, 8, false, 100},  /* RSSI_CHANN_DEFAULT_VALUE */
	"LGW_RSSI_BB_FILTER_ALPHA":                {0, 111, 0, false, 5, false, 7},    /* RSSI_BB_FILTER_ALPHA */
	"LGW_RSSI_DEC_FILTER_ALPHA":               {0, 112, 0, false, 5, false, 5},    /* RSSI_DEC_FILTER_ALPHA */
	"LGW_RSSI_CHANN_FILTER_ALPHA":             {0, 113, 0, false, 5, false, 8},    /* RSSI_CHANN_FILTER_ALPHA */
	"LGW_IQ_MISMATCH_A_AMP_COEFF":             {0, 114, 0, false, 6, false, 0},    /* IQ_MISMATCH_A_AMP_COEFF */
	"LGW_IQ_MISMATCH_A_PHI_COEFF":             {0, 115, 0, false, 6, false, 0},    /* IQ_MISMATCH_A_PHI_COEFF */
	"LGW_IQ_MISMATCH_B_AMP_COEFF":             {0, 116, 0, false, 6, false, 0},    /* IQ_MISMATCH_B_AMP_COEFF */
	"LGW_IQ_MISMATCH_B_SEL_I":                 {0, 116, 6, false, 1, false, 0},    /* IQ_MISMATCH_B_SEL_I */
	"LGW_IQ_MISMATCH_B_PHI_COEFF":             {0, 117, 0, false, 6, false, 0},    /* IQ_MISMATCH_B_PHI_COEFF */
	"LGW_TX_TRIG_IMMEDIATE":                   {1, 33, 0, false, 1, false, 0},     /* TX_TRIG_IMMEDIATE */
	"LGW_TX_TRIG_DELAYED":                     {1, 33, 1, false, 1, false, 0},     /* TX_TRIG_DELAYED */
	"LGW_TX_TRIG_GPS":                         {1, 33, 2, false, 1, false, 0},     /* TX_TRIG_GPS */
	"LGW_TX_START_DELAY":                      {1, 34, 0, false, 16, false, 0},    /* TX_START_DELAY */
	"LGW_TX_FRAME_SYNCH_PEAK1_POS":            {1, 36, 0, false, 4, false, 1},     /* TX_FRAME_SYNCH_PEAK1_POS */ //I THINK WHAT SYNCWORD IS USED FOR TX
	"LGW_TX_FRAME_SYNCH_PEAK2_POS":            {1, 36, 4, false, 4, false, 2},     /* TX_FRAME_SYNCH_PEAK2_POS */ //I THINK WHAT SYNCWORD IS USED FOR TX
	"LGW_TX_RAMP_DURATION":                    {1, 37, 0, false, 3, false, 0},     /* TX_RAMP_DURATION */
	"LGW_TX_OFFSET_I":                         {1, 39, 0, true, 8, false, 0},      /* TX_OFFSET_I */
	"LGW_TX_OFFSET_Q":                         {1, 40, 0, true, 8, false, 0},      /* TX_OFFSET_Q */
	"LGW_TX_MODE":                             {1, 41, 0, false, 1, false, 0},     /* TX_MODE */
	"LGW_TX_ZERO_PAD":                         {1, 41, 1, false, 4, false, 0},     /* TX_ZERO_PAD */
	"LGW_TX_EDGE_SELECT":                      {1, 41, 5, false, 1, false, 0},     /* TX_EDGE_SELECT */
	"LGW_TX_EDGE_SELECT_TOP":                  {1, 41, 6, false, 1, false, 0},     /* TX_EDGE_SELECT_TOP */
	"LGW_TX_GAIN":                             {1, 42, 0, false, 2, false, 0},     /* TX_GAIN */
	"LGW_TX_CHIRP_LOW_PASS":                   {1, 42, 2, false, 3, false, 5},     /* TX_CHIRP_LOW_PASS */
	"LGW_TX_FCC_WIDEBAND":                     {1, 42, 5, false, 2, false, 0},     /* TX_FCC_WIDEBAND */
	"LGW_TX_SWAP_IQ":                          {1, 42, 7, false, 1, false, 1},     /* TX_SWAP_IQ */
	"LGW_MBWSSF_IMPLICIT_HEADER":              {1, 43, 0, false, 1, false, 0},     /* MBWSSF_IMPLICIT_HEADER */
	"LGW_MBWSSF_IMPLICIT_CRC_EN":              {1, 43, 1, false, 1, false, 0},     /* MBWSSF_IMPLICIT_CRC_EN */
	"LGW_MBWSSF_IMPLICIT_CODING_RATE":         {1, 43, 2, false, 3, false, 0},     /* MBWSSF_IMPLICIT_CODING_RATE */
	"LGW_MBWSSF_IMPLICIT_PAYLOAD_LENGHT":      {1, 44, 0, false, 8, false, 0},     /* MBWSSF_IMPLICIT_PAYLOAD_LENGHT */
	"LGW_MBWSSF_AGC_FREEZE_ON_DETECT":         {1, 45, 0, false, 1, false, 1},     /* MBWSSF_AGC_FREEZE_ON_DETECT */
	"LGW_MBWSSF_FRAME_SYNCH_PEAK1_POS":        {1, 46, 0, false, 4, false, 1},     /* MBWSSF_FRAME_SYNCH_PEAK1_POS */ //AKA SYNC WORD
	"LGW_MBWSSF_FRAME_SYNCH_PEAK2_POS":        {1, 46, 4, false, 4, false, 2},     /* MBWSSF_FRAME_SYNCH_PEAK2_POS */ //AKA SYNC WORD
	"LGW_MBWSSF_PREAMBLE_SYMB1_NB":            {1, 47, 0, false, 16, false, 10},   /* MBWSSF_PREAMBLE_SYMB1_NB */
	"LGW_MBWSSF_FRAME_SYNCH_GAIN":             {1, 49, 0, false, 1, false, 1},     /* MBWSSF_FRAME_SYNCH_GAIN */
	"LGW_MBWSSF_SYNCH_DETECT_TH":              {1, 49, 1, false, 1, false, 1},     /* MBWSSF_SYNCH_DETECT_TH */
	"LGW_MBWSSF_DETECT_MIN_SINGLE_PEAK":       {1, 50, 0, false, 8, false, 10},    /* MBWSSF_DETECT_MIN_SINGLE_PEAK */
	"LGW_MBWSSF_DETECT_TRIG_SAME_PEAK_NB":     {1, 51, 0, false, 3, false, 3},     /* MBWSSF_DETECT_TRIG_SAME_PEAK_NB */
	"LGW_MBWSSF_FREQ_TO_TIME_INVERT":          {1, 52, 0, false, 8, false, 29},    /* MBWSSF_FREQ_TO_TIME_INVERT */
	"LGW_MBWSSF_FREQ_TO_TIME_DRIFT":           {1, 53, 0, false, 6, false, 36},    /* MBWSSF_FREQ_TO_TIME_DRIFT */
	"LGW_MBWSSF_PPM_CORRECTION":               {1, 54, 0, false, 12, false, 0},    /* MBWSSF_PPM_CORRECTION */
	"LGW_MBWSSF_PAYLOAD_FINE_TIMING_GAIN":     {1, 56, 0, false, 2, false, 2},     /* MBWSSF_PAYLOAD_FINE_TIMING_GAIN */
	"LGW_MBWSSF_PREAMBLE_FINE_TIMING_GAIN":    {1, 56, 2, false, 2, false, 1},     /* MBWSSF_PREAMBLE_FINE_TIMING_GAIN */
	"LGW_MBWSSF_TRACKING_INTEGRAL":            {1, 56, 4, false, 2, false, 0},     /* MBWSSF_TRACKING_INTEGRAL */
	"LGW_MBWSSF_ZERO_PAD":                     {1, 57, 0, false, 8, false, 0},     /* MBWSSF_ZERO_PAD */
	"LGW_MBWSSF_MODEM_BW":                     {1, 58, 0, false, 2, false, 0},     /* MBWSSF_MODEM_BW */
	"LGW_MBWSSF_RADIO_SELECT":                 {1, 58, 2, false, 1, false, 0},     /* MBWSSF_RADIO_SELECT */
	"LGW_MBWSSF_RX_CHIRP_INVERT":              {1, 58, 3, false, 1, false, 1},     /* MBWSSF_RX_CHIRP_INVERT */
	"LGW_MBWSSF_LLR_SCALE":                    {1, 59, 0, false, 4, false, 8},     /* MBWSSF_LLR_SCALE */
	"LGW_MBWSSF_SNR_AVG_CST":                  {1, 59, 4, false, 2, false, 3},     /* MBWSSF_SNR_AVG_CST */
	"LGW_MBWSSF_PPM_OFFSET":                   {1, 59, 6, false, 1, false, 0},     /* MBWSSF_PPM_OFFSET */
	"LGW_MBWSSF_RATE_SF":                      {1, 60, 0, false, 4, false, 7},     /* MBWSSF_RATE_SF */
	"LGW_MBWSSF_ONLY_CRC_EN":                  {1, 60, 4, false, 1, false, 1},     /* MBWSSF_ONLY_CRC_EN */
	"LGW_MBWSSF_MAX_PAYLOAD_LEN":              {1, 61, 0, false, 8, false, 255},   /* MBWSSF_MAX_PAYLOAD_LEN */
	"LGW_TX_STATUS":                           {1, 62, 0, false, 8, true, 128},    /* TX_STATUS */
	"LGW_FSK_CH_BW_EXPO":                      {1, 63, 0, false, 3, false, 0},     /* FSK_CH_BW_EXPO */
	"LGW_FSK_RSSI_LENGTH":                     {1, 63, 3, false, 3, false, 0},     /* FSK_RSSI_LENGTH */
	"LGW_FSK_RX_INVERT":                       {1, 63, 6, false, 1, false, 0},     /* FSK_RX_INVERT */
	"LGW_FSK_PKT_MODE":                        {1, 63, 7, false, 1, false, 0},     /* FSK_PKT_MODE */
	"LGW_FSK_PSIZE":                           {1, 64, 0, false, 3, false, 0},     /* FSK_PSIZE */
	"LGW_FSK_CRC_EN":                          {1, 64, 3, false, 1, false, 0},     /* FSK_CRC_EN */
	"LGW_FSK_DCFREE_ENC":                      {1, 64, 4, false, 2, false, 0},     /* FSK_DCFREE_ENC */
	"LGW_FSK_CRC_IBM":                         {1, 64, 6, false, 1, false, 0},     /* FSK_CRC_IBM */
	"LGW_FSK_ERROR_OSR_TOL":                   {1, 65, 0, false, 5, false, 0},     /* FSK_ERROR_OSR_TOL */
	"LGW_FSK_RADIO_SELECT":                    {1, 65, 7, false, 1, false, 0},     /* FSK_RADIO_SELECT */
	"LGW_FSK_BR_RATIO":                        {1, 66, 0, false, 16, false, 0},    /* FSK_BR_RATIO */
	"LGW_FSK_REF_PATTERN_LSB":                 {1, 68, 0, false, 32, false, 0},    /* FSK_REF_PATTERN_LSB */
	"LGW_FSK_REF_PATTERN_MSB":                 {1, 72, 0, false, 32, false, 0},    /* FSK_REF_PATTERN_MSB */
	"LGW_FSK_PKT_LENGTH":                      {1, 76, 0, false, 8, false, 0},     /* FSK_PKT_LENGTH */
	"LGW_FSK_TX_GAUSSIAN_EN":                  {1, 77, 0, false, 1, false, 1},     /* FSK_TX_GAUSSIAN_EN */
	"LGW_FSK_TX_GAUSSIAN_SELECT_BT":           {1, 77, 1, false, 2, false, 0},     /* FSK_TX_GAUSSIAN_SELECT_BT */
	"LGW_FSK_TX_PATTERN_EN":                   {1, 77, 3, false, 1, false, 1},     /* FSK_TX_PATTERN_EN */
	"LGW_FSK_TX_PREAMBLE_SEQ":                 {1, 77, 4, false, 1, false, 0},     /* FSK_TX_PREAMBLE_SEQ */
	"LGW_FSK_TX_PSIZE":                        {1, 77, 5, false, 3, false, 0},     /* FSK_TX_PSIZE */
	"LGW_FSK_NODE_ADRS":                       {1, 80, 0, false, 8, false, 0},     /* FSK_NODE_ADRS */
	"LGW_FSK_BROADCAST":                       {1, 81, 0, false, 8, false, 0},     /* FSK_BROADCAST */
	"LGW_FSK_AUTO_AFC_ON":                     {1, 82, 0, false, 1, false, 1},     /* FSK_AUTO_AFC_ON */
	"LGW_FSK_PATTERN_TIMEOUT_CFG":             {1, 83, 0, false, 10, false, 0},    /* FSK_PATTERN_TIMEOUT_CFG */
	"LGW_SPI_RADIO_A__DATA":                   {2, 33, 0, false, 8, false, 0},     /* SPI_RADIO_A__DATA */
	"LGW_SPI_RADIO_A__DATA_READBACK":          {2, 34, 0, false, 8, true, 0},      /* SPI_RADIO_A__DATA_READBACK */
	"LGW_SPI_RADIO_A__ADDR":                   {2, 35, 0, false, 8, false, 0},     /* SPI_RADIO_A__ADDR */
	"LGW_SPI_RADIO_A__CS":                     {2, 37, 0, false, 1, false, 0},     /* SPI_RADIO_A__CS */
	"LGW_SPI_RADIO_B__DATA":                   {2, 38, 0, false, 8, false, 0},     /* SPI_RADIO_B__DATA */
	"LGW_SPI_RADIO_B__DATA_READBACK":          {2, 39, 0, false, 8, true, 0},      /* SPI_RADIO_B__DATA_READBACK */
	"LGW_SPI_RADIO_B__ADDR":                   {2, 40, 0, false, 8, false, 0},     /* SPI_RADIO_B__ADDR */
	"LGW_SPI_RADIO_B__CS":                     {2, 42, 0, false, 1, false, 0},     /* SPI_RADIO_B__CS */
	"LGW_RADIO_A_EN":                          {2, 43, 0, false, 1, false, 0},     /* RADIO_A_EN */
	"LGW_RADIO_B_EN":                          {2, 43, 1, false, 1, false, 0},     /* RADIO_B_EN */
	"LGW_RADIO_RST":                           {2, 43, 2, false, 1, false, 1},     /* RADIO_RST */
	"LGW_LNA_A_EN":                            {2, 43, 3, false, 1, false, 0},     /* LNA_A_EN */
	"LGW_PA_A_EN":                             {2, 43, 4, false, 1, false, 0},     /* PA_A_EN */
	"LGW_LNA_B_EN":                            {2, 43, 5, false, 1, false, 0},     /* LNA_B_EN */
	"LGW_PA_B_EN":                             {2, 43, 6, false, 1, false, 0},     /* PA_B_EN */
	"LGW_PA_GAIN":                             {2, 44, 0, false, 2, false, 0},     /* PA_GAIN */
	"LGW_LNA_A_CTRL_LUT":                      {2, 45, 0, false, 4, false, 2},     /* LNA_A_CTRL_LUT */
	"LGW_PA_A_CTRL_LUT":                       {2, 45, 4, false, 4, false, 4},     /* PA_A_CTRL_LUT */
	"LGW_LNA_B_CTRL_LUT":                      {2, 46, 0, false, 4, false, 2},     /* LNA_B_CTRL_LUT */
	"LGW_PA_B_CTRL_LUT":                       {2, 46, 4, false, 4, false, 4},     /* PA_B_CTRL_LUT */
	"LGW_CAPTURE_SOURCE":                      {2, 47, 0, false, 5, false, 0},     /* CAPTURE_SOURCE */
	"LGW_CAPTURE_START":                       {2, 47, 5, false, 1, false, 0},     /* CAPTURE_START */
	"LGW_CAPTURE_FORCE_TRIGGER":               {2, 47, 6, false, 1, false, 0},     /* CAPTURE_FORCE_TRIGGER */
	"LGW_CAPTURE_WRAP":                        {2, 47, 7, false, 1, false, 0},     /* CAPTURE_WRAP */
	"LGW_CAPTURE_PERIOD":                      {2, 48, 0, false, 16, false, 0},    /* CAPTURE_PERIOD */
	"LGW_MODEM_STATUS":                        {2, 51, 0, false, 8, true, 0},      /* MODEM_STATUS */
	"LGW_VALID_HEADER_COUNTER_0":              {2, 52, 0, false, 8, true, 0},      /* VALID_HEADER_COUNTER_0 */
	"LGW_VALID_PACKET_COUNTER_0":              {2, 54, 0, false, 8, true, 0},      /* VALID_PACKET_COUNTER_0 */
	"LGW_VALID_HEADER_COUNTER_MBWSSF":         {2, 56, 0, false, 8, true, 0},      /* VALID_HEADER_COUNTER_MBWSSF */
	"LGW_VALID_HEADER_COUNTER_FSK":            {2, 57, 0, false, 8, true, 0},      /* VALID_HEADER_COUNTER_FSK */
	"LGW_VALID_PACKET_COUNTER_MBWSSF":         {2, 58, 0, false, 8, true, 0},      /* VALID_PACKET_COUNTER_MBWSSF */
	"LGW_VALID_PACKET_COUNTER_FSK":            {2, 59, 0, false, 8, true, 0},      /* VALID_PACKET_COUNTER_FSK */
	"LGW_CHANN_RSSI":                          {2, 60, 0, false, 8, true, 0},      /* CHANN_RSSI */
	"LGW_BB_RSSI":                             {2, 61, 0, false, 8, true, 0},      /* BB_RSSI */
	"LGW_DEC_RSSI":                            {2, 62, 0, false, 8, true, 0},      /* DEC_RSSI */
	"LGW_DBG_MCU_DATA":                        {2, 63, 0, false, 8, true, 0},      /* DBG_MCU_DATA */
	"LGW_DBG_ARB_MCU_RAM_DATA":                {2, 64, 0, false, 8, true, 0},      /* DBG_ARB_MCU_RAM_DATA */
	"LGW_DBG_AGC_MCU_RAM_DATA":                {2, 65, 0, false, 8, true, 0},      /* DBG_AGC_MCU_RAM_DATA */
	"LGW_NEXT_PACKET_CNT":                     {2, 66, 0, false, 16, true, 0},     /* NEXT_PACKET_CNT */
	"LGW_ADDR_CAPTURE_COUNT":                  {2, 68, 0, false, 16, true, 0},     /* ADDR_CAPTURE_COUNT */
	"LGW_TIMESTAMP":                           {2, 70, 0, false, 32, true, 0},     /* TIMESTAMP */
	"LGW_DBG_CHANN0_GAIN":                     {2, 74, 0, false, 4, true, 0},      /* DBG_CHANN0_GAIN */
	"LGW_DBG_CHANN1_GAIN":                     {2, 74, 4, false, 4, true, 0},      /* DBG_CHANN1_GAIN */
	"LGW_DBG_CHANN2_GAIN":                     {2, 75, 0, false, 4, true, 0},      /* DBG_CHANN2_GAIN */
	"LGW_DBG_CHANN3_GAIN":                     {2, 75, 4, false, 4, true, 0},      /* DBG_CHANN3_GAIN */
	"LGW_DBG_CHANN4_GAIN":                     {2, 76, 0, false, 4, true, 0},      /* DBG_CHANN4_GAIN */
	"LGW_DBG_CHANN5_GAIN":                     {2, 76, 4, false, 4, true, 0},      /* DBG_CHANN5_GAIN */
	"LGW_DBG_CHANN6_GAIN":                     {2, 77, 0, false, 4, true, 0},      /* DBG_CHANN6_GAIN */
	"LGW_DBG_CHANN7_GAIN":                     {2, 77, 4, false, 4, true, 0},      /* DBG_CHANN7_GAIN */
	"LGW_DBG_DEC_FILT_GAIN":                   {2, 78, 0, false, 4, true, 0},      /* DBG_DEC_FILT_GAIN */
	"LGW_SPI_DATA_FIFO_PTR":                   {2, 79, 0, false, 3, true, 0},      /* SPI_DATA_FIFO_PTR */
	"LGW_PACKET_DATA_FIFO_PTR":                {2, 79, 3, false, 3, true, 0},      /* PACKET_DATA_FIFO_PTR */
	"LGW_DBG_ARB_MCU_RAM_ADDR":                {2, 80, 0, false, 8, false, 0},     /* DBG_ARB_MCU_RAM_ADDR */
	"LGW_DBG_AGC_MCU_RAM_ADDR":                {2, 81, 0, false, 8, false, 0},     /* DBG_AGC_MCU_RAM_ADDR */
	"LGW_SPI_MASTER_CHIP_SELECT_POLARITY":     {2, 82, 0, false, 1, false, 0},     /* SPI_MASTER_CHIP_SELECT_POLARITY */
	"LGW_SPI_MASTER_CPOL":                     {2, 82, 1, false, 1, false, 0},     /* SPI_MASTER_CPOL */
	"LGW_SPI_MASTER_CPHA":                     {2, 82, 2, false, 1, false, 0},     /* SPI_MASTER_CPHA */
	"LGW_SIG_GEN_ANALYSER_MUX_SEL":            {2, 83, 0, false, 1, false, 0},     /* SIG_GEN_ANALYSER_MUX_SEL */
	"LGW_SIG_GEN_EN":                          {2, 84, 0, false, 1, false, 0},     /* SIG_GEN_EN */
	"LGW_SIG_ANALYSER_EN":                     {2, 84, 1, false, 1, false, 0},     /* SIG_ANALYSER_EN */
	"LGW_SIG_ANALYSER_AVG_LEN":                {2, 84, 2, false, 2, false, 0},     /* SIG_ANALYSER_AVG_LEN */
	"LGW_SIG_ANALYSER_PRECISION":              {2, 84, 4, false, 3, false, 0},     /* SIG_ANALYSER_PRECISION */
	"LGW_SIG_ANALYSER_VALID_OUT":              {2, 84, 7, false, 1, true, 0},      /* SIG_ANALYSER_VALID_OUT */
	"LGW_SIG_GEN_FREQ":                        {2, 85, 0, false, 8, false, 0},     /* SIG_GEN_FREQ */
	"LGW_SIG_ANALYSER_FREQ":                   {2, 86, 0, false, 8, false, 0},     /* SIG_ANALYSER_FREQ */
	"LGW_SIG_ANALYSER_I_OUT":                  {2, 87, 0, false, 8, true, 0},      /* SIG_ANALYSER_I_OUT */
	"LGW_SIG_ANALYSER_Q_OUT":                  {2, 88, 0, false, 8, true, 0},      /* SIG_ANALYSER_Q_OUT */
	"LGW_GPS_EN":                              {2, 89, 0, false, 1, false, 0},     /* GPS_EN */
	"LGW_GPS_POL":                             {2, 89, 1, false, 1, false, 1},     /* GPS_POL */
	"LGW_SW_TEST_REG1":                        {2, 90, 0, true, 8, false, 0},      /* SW_TEST_REG1 */
	"LGW_SW_TEST_REG2":                        {2, 91, 2, true, 6, false, 0},      /* SW_TEST_REG2 */
	"LGW_SW_TEST_REG3":                        {2, 92, 0, true, 16, false, 0},     /* SW_TEST_REG3 */
	"LGW_DATA_MNGT_STATUS":                    {2, 94, 0, false, 4, true, 0},      /* DATA_MNGT_STATUS */
	"LGW_DATA_MNGT_CPT_FRAME_ALLOCATED":       {2, 95, 0, false, 5, true, 0},      /* DATA_MNGT_CPT_FRAME_ALLOCATED */
	"LGW_DATA_MNGT_CPT_FRAME_FINISHED":        {2, 96, 0, false, 5, true, 0},      /* DATA_MNGT_CPT_FRAME_FINISHED */
	"LGW_DATA_MNGT_CPT_FRAME_READEN":          {2, 97, 0, false, 5, true, 0},      /* DATA_MNGT_CPT_FRAME_READEN */
	"LGW_TX_TRIG_ALL":                         {1, 33, 0, false, 8, false, 0},
}
