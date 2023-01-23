# lo工具包
go 工具包，包含类似java相关的steam api， map， redis， foreach 和 filter，包括串行和并行模式

https://github.com/samber/lo

Supported helpers for slices:

Filter
Map
FilterMap
FlatMap
Reduce
ReduceRight
ForEach
Times
Uniq
UniqBy
GroupBy
Chunk
PartitionBy
Flatten
Interleave
Shuffle
Reverse
Fill
Repeat
RepeatBy
KeyBy
Associate / SliceToMap
Drop
DropRight
DropWhile
DropRightWhile
Reject
Count
CountBy
CountValues
CountValuesBy
Subset
Slice
Replace
ReplaceAll
Compact
IsSorted
IsSortedByKey
Supported helpers for maps:

Keys
Values
PickBy
PickByKeys
PickByValues
OmitBy
OmitByKeys
OmitByValues
Entries / ToPairs
FromEntries / FromPairs
Invert
Assign (merge of maps)
MapKeys
MapValues
MapEntries
MapToSlice
Supported math helpers:

Range / RangeFrom / RangeWithSteps
Clamp
Sum
SumBy
Supported helpers for strings:

RandomString
Substring
ChunkString
RuneLength
Supported helpers for tuples:

T2 -> T9
Unpack2 -> Unpack9
Zip2 -> Zip9
Unzip2 -> Unzip9
Supported helpers for channels:

ChannelDispatcher
SliceToChannel
<!-- Generator -->
Buffer
BufferWithTimeout
FanIn
FanOut
Supported intersection helpers:

Contains
ContainsBy
Every
EveryBy
Some
SomeBy
None
NoneBy
Intersect
Difference
Union
Without
WithoutEmpty
Supported search helpers:

IndexOf
LastIndexOf
Find
FindIndexOf
FindLastIndexOf
FindKey
FindKeyBy
FindUniques
FindUniquesBy
FindDuplicates
FindDuplicatesBy
Min
MinBy
Max
MaxBy
Last
Nth
Sample
Samples
Conditional helpers:

Ternary
TernaryF
If / ElseIf / Else
Switch / Case / Default
Type manipulation helpers:

ToPtr
FromPtr
FromPtrOr
ToSlicePtr
ToAnySlice
FromAnySlice
Empty
IsEmpty
IsNotEmpty
Coalesce
Function helpers:

Partial
Partial2 -> Partial5
Concurrency helpers:

Attempt
AttemptWhile
AttemptWithDelay
AttemptWhileWithDelay
Debounce
Synchronize
Async
Transaction
Error handling:

Validate
Must
Try
Try1 -> Try6
TryOr
TryOr1 -> TryOr6
TryCatch
TryWithErrorValue
TryCatchWithErrorValue
ErrorsAs
Constraints:

Clonable